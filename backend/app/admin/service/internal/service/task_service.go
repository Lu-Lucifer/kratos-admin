package service

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
	paginationV1 "github.com/tx7do/go-crud/api/gen/go/pagination/v1"
	"github.com/tx7do/go-utils/trans"
	"github.com/tx7do/kratos-bootstrap/bootstrap"
	"google.golang.org/protobuf/types/known/emptypb"

	"go-wind-admin/app/admin/service/internal/data"

	adminV1 "go-wind-admin/api/gen/go/admin/service/v1"
	taskV1 "go-wind-admin/api/gen/go/task/service/v1"

	"go-wind-admin/pkg/middleware/auth"
	"go-wind-admin/pkg/oss"
	"go-wind-admin/pkg/task"
)

// TaskScheduler 任务调度接口
type TaskScheduler interface {
	TaskTypeExists(taskType string) bool
	GetRegisteredTaskTypes() []string

	NewTask(typeName string, msg any, opts ...asynq.Option) error
	NewWaitResultTask(typeName string, msg any, opts ...asynq.Option) error
	NewPeriodicTask(cronSpec, typeName string, msg any, opts ...asynq.Option) (string, error)

	RemovePeriodicTask(id string) error
	RemoveAllPeriodicTask()
}

// backupBucket 备份文件存放的 OSS 桶名。
const backupBucket = "backups"

// TaskService 任务服务
type TaskService struct {
	adminV1.TaskServiceHTTPServer

	log *log.Helper

	taskScheduler TaskScheduler

	userRepo   data.UserRepo
	taskRepo   *data.TaskRepo
	backupRepo *data.BackupRepo
	mc         *oss.MinIOClient
}

func NewTaskService(
	ctx *bootstrap.Context,
	taskRepo *data.TaskRepo,
	userRepo data.UserRepo,
	backupRepo *data.BackupRepo,
	mc *oss.MinIOClient,
) *TaskService {
	svc := &TaskService{
		log:        ctx.NewLoggerHelper("task/service/admin-service"),
		taskRepo:   taskRepo,
		userRepo:   userRepo,
		backupRepo: backupRepo,
		mc:         mc,
	}

	return svc
}

func (s *TaskService) RegisterTaskScheduler(taskScheduler TaskScheduler) {
	s.taskScheduler = taskScheduler
}

// hasScheduler 检查调度器是否可用（未配置 asynq 时为 nil）。
// 在所有调用 taskScheduler 的地方前置检查，避免 nil 解引用 panic。
func (s *TaskService) hasScheduler() bool {
	return s.taskScheduler != nil
}

func (s *TaskService) List(ctx context.Context, req *paginationV1.PagingRequest) (*taskV1.ListTaskResponse, error) {
	return s.taskRepo.List(ctx, req)
}

func (s *TaskService) Get(ctx context.Context, req *taskV1.GetTaskRequest) (*taskV1.Task, error) {
	return s.taskRepo.Get(ctx, req)
}

func (s *TaskService) ListTaskTypeName(_ context.Context, _ *emptypb.Empty) (*taskV1.ListTaskTypeNameResponse, error) {
	// nil 调度器防护：asynq 未配置时返回空列表而非 panic
	if !s.hasScheduler() {
		return &taskV1.ListTaskTypeNameResponse{}, nil
	}
	typeNames := s.taskScheduler.GetRegisteredTaskTypes()
	return &taskV1.ListTaskTypeNameResponse{
		TypeNames: typeNames,
	}, nil
}

func (s *TaskService) Create(ctx context.Context, req *taskV1.CreateTaskRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// H8：校验 typeName 已在调度器注册，防止创建无 handler 的幽灵任务（每个 cron tick 报错）
	if !s.hasScheduler() || !s.taskScheduler.TaskTypeExists(req.Data.GetTypeName()) {
		return nil, adminV1.ErrorBadRequest("task type [%s] is not registered", req.Data.GetTypeName())
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	req.Data.CreatedBy = trans.Ptr(operator.UserId)

	var t *taskV1.Task
	if t, err = s.taskRepo.Create(ctx, req); err != nil {
		return nil, err
	}

	if err = s.startTask(t); err != nil {
		// 调度失败不掩盖：DB 记录已建，但任务实际不会运行，需明确告知
		s.log.Errorf("create task [%s] succeeded but scheduling failed: %s", t.GetTypeName(), err.Error())
		return nil, adminV1.ErrorInternalServerError("task scheduling failed: %s", err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *TaskService) Update(ctx context.Context, req *taskV1.UpdateTaskRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// H8：校验 typeName 已在调度器注册
	if !s.hasScheduler() || !s.taskScheduler.TaskTypeExists(req.Data.GetTypeName()) {
		return nil, adminV1.ErrorBadRequest("task type [%s] is not registered", req.Data.GetTypeName())
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	// 获取更新前的任务，用于判断调度器中是否有正在运行的注册项
	oldTask, _ := s.taskRepo.Get(ctx, &taskV1.GetTaskRequest{QueryBy: &taskV1.GetTaskRequest_Id{Id: req.GetId()}})

	req.Data.Id = trans.Ptr(req.GetId())

	req.Data.UpdatedBy = trans.Ptr(operator.UserId)
	if req.UpdateMask != nil {
		req.UpdateMask.Paths = append(req.UpdateMask.Paths, "updated_by")
	}

	var t *taskV1.Task
	if t, err = s.taskRepo.Update(ctx, req); err != nil {

		return nil, err
	}

	// 先移除调度器中旧的注册项（若存在），避免停用后仍运行、或启用时重复注册。
	// 直接调用 RemovePeriodicTask 绕过 stopTask 内部的 enable 保护，
	// 因为 oldTask 可能已是禁用状态，但其注册项可能仍残留在调度器中。
	if s.hasScheduler() && oldTask != nil && oldTask.GetType() == taskV1.Task_PERIODIC && oldTask.GetTypeName() != "" {
		if removeErr := s.taskScheduler.RemovePeriodicTask(oldTask.GetTypeName()); removeErr != nil {
			s.log.Warnf("移除旧定时任务注册项失败[%s]: %v", oldTask.GetTypeName(), removeErr)
		}
	}

	// 根据更新后的 enable 状态决定是否重新启动
	if err = s.startTask(t); err != nil {
		s.log.Error(err)
		return nil, adminV1.ErrorInternalServerError("task scheduling failed: %s", err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *TaskService) Delete(ctx context.Context, req *taskV1.DeleteTaskRequest) (*emptypb.Empty, error) {
	var err error
	var t *taskV1.Task
	if t, err = s.taskRepo.Get(ctx, &taskV1.GetTaskRequest{QueryBy: &taskV1.GetTaskRequest_Id{Id: req.GetId()}}); err != nil {
		s.log.Error(err)
	}

	if err = s.taskRepo.Delete(ctx, req); err != nil {
		return nil, err
	}

	if t != nil {
		_ = s.stopTask(t)
	}

	return &emptypb.Empty{}, nil
}

// ControlTask 控制调度任务
func (s *TaskService) ControlTask(ctx context.Context, req *taskV1.ControlTaskRequest) (*emptypb.Empty, error) {
	t, err := s.taskRepo.Get(ctx, &taskV1.GetTaskRequest{QueryBy: &taskV1.GetTaskRequest_TypeName{TypeName: req.GetTypeName()}})
	if err != nil {
		s.log.Errorf("获取任务失败[%s]", err.Error())
		return nil, err
	}

	switch req.GetControlType() {
	case taskV1.ControlTaskRequest_Restart:
		if err = s.stopTask(t); err != nil {
			return nil, err
		}

		if err = s.startTask(t); err != nil {
			return nil, err
		}

	case taskV1.ControlTaskRequest_Stop:
		err = s.stopTask(t)
		return nil, err

	case taskV1.ControlTaskRequest_Start:
		err = s.startTask(t)
		return nil, err

	default:
		// H8：未知控制类型必须明确拒绝，避免静默返回成功
		return nil, adminV1.ErrorBadRequest("invalid control type")
	}

	return &emptypb.Empty{}, nil
}

// StopAllTask 停止所有的调度任务
func (s *TaskService) StopAllTask(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	s.stopAllTask()
	return &emptypb.Empty{}, nil
}

// StartAllTask 启动所有的调度任务
func (s *TaskService) StartAllTask(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	_, err := s.startAllTask(ctx)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// RestartAllTask 重启所有的调度任务
func (s *TaskService) RestartAllTask(ctx context.Context, _ *emptypb.Empty) (*taskV1.RestartAllTaskResponse, error) {
	// 停止所有的任务
	s.stopAllTask()

	// 重新启动所有的任务
	count, err := s.startAllTask(ctx)

	return &taskV1.RestartAllTaskResponse{
		Count: count,
	}, err
}

// startAllTask 启动所有的任务
func (s *TaskService) startAllTask(ctx context.Context) (int32, error) {
	resp, err := s.List(ctx, &paginationV1.PagingRequest{
		NoPaging: trans.Ptr(true),
	})
	if err != nil {
		s.log.Errorf("获取任务列表失败[%s]", err.Error())
		return 0, err
	}

	s.log.Infof("开始开启定时任务，总计[%d]个", resp.GetTotal())

	// 重新启动任务
	var count int32
	for _, t := range resp.GetItems() {
		if s.startTask(t) != nil {
			continue
		} else {
			count++
		}
	}

	s.log.Infof("总共成功开启定时任务[%d]个", count)

	return count, nil
}

// stopAllTask 停止所有的任务
func (s *TaskService) stopAllTask() {
	// nil 调度器防护
	if !s.hasScheduler() {
		s.log.Warnf("task scheduler is not configured, skip stopAllTask")
		return
	}

	s.log.Infof("开始清除所有的定时任务...")

	// 清除所有的定时任务
	s.taskScheduler.RemoveAllPeriodicTask()

	s.log.Infof("完成清除所有的定时任务")
}

// stopTask 停止一个任务
func (s *TaskService) stopTask(t *taskV1.Task) error {
	if t == nil {
		return errors.New("task is nil")
	}

	if t.GetEnable() == false {
		return errors.New("task is not enable")
	}

	// nil 调度器防护
	if !s.hasScheduler() {
		return errors.New("task scheduler is not configured")
	}

	switch t.GetType() {
	case taskV1.Task_PERIODIC:
		return s.taskScheduler.RemovePeriodicTask(t.GetTypeName())

	case taskV1.Task_DELAY:

	case taskV1.Task_WAIT_RESULT:
	}

	return nil
}

// convertTaskOption 转换任务选项
func (s *TaskService) convertTaskOption(t *taskV1.Task) (opts []asynq.Option, payload any) {
	if t == nil {
		return
	}

	if len(t.GetTaskPayload()) > 0 {
		_ = json.Unmarshal([]byte(t.GetTaskPayload()), &payload)
	}

	if t.TaskOptions != nil {
		if t.GetTaskOptions().GetMaxRetry() > 0 {
			opts = append(opts, asynq.MaxRetry(int(t.GetTaskOptions().GetMaxRetry())))
		}
		if t.GetTaskOptions().Timeout != nil {
			opts = append(opts, asynq.Timeout(t.GetTaskOptions().GetTimeout().AsDuration()))
		}
		if t.GetTaskOptions().Deadline != nil {
			opts = append(opts, asynq.Deadline(t.GetTaskOptions().GetDeadline().AsTime()))
		}
		if t.GetTaskOptions().ProcessIn != nil {
			opts = append(opts, asynq.ProcessIn(t.GetTaskOptions().GetProcessIn().AsDuration()))
		}
		if t.GetTaskOptions().ProcessAt != nil {
			opts = append(opts, asynq.ProcessAt(t.GetTaskOptions().GetProcessAt().AsTime()))
		}
		if t.GetTaskOptions().UniqueTtl != nil {
			opts = append(opts, asynq.Unique(t.GetTaskOptions().GetUniqueTtl().AsDuration()))
		}
		if t.GetTaskOptions().Retention != nil {
			opts = append(opts, asynq.Retention(t.GetTaskOptions().GetRetention().AsDuration()))
		}
		opts = append(opts, asynq.Group(t.GetTaskOptions().GetGroup()))
		opts = append(opts, asynq.TaskID(t.GetTaskOptions().GetTaskId()))
	}

	return
}

// startTask 启动一个任务
func (s *TaskService) startTask(t *taskV1.Task) error {
	if t == nil {
		return errors.New("task is nil")
	}

	if t.GetEnable() == false {
		return errors.New("task is not enable")
	}

	// nil 调度器防护
	if !s.hasScheduler() {
		return errors.New("task scheduler is not configured")
	}

	var opts []asynq.Option
	var payload any
	var err error

	switch t.GetType() {
	case taskV1.Task_PERIODIC:
		opts, payload = s.convertTaskOption(t)
		if _, err = s.taskScheduler.NewPeriodicTask(t.GetCronSpec(), t.GetTypeName(), payload, opts...); err != nil {
			s.log.Errorf("[%s] 创建定时任务失败[%s]", t.GetTypeName(), err.Error())
			return err
		}

	case taskV1.Task_DELAY:
		opts, payload = s.convertTaskOption(t)
		if err = s.taskScheduler.NewTask(t.GetTypeName(), payload, opts...); err != nil {
			s.log.Errorf("[%s] 创建延迟任务失败[%s]", t.GetTypeName(), err.Error())
			return err
		}

	case taskV1.Task_WAIT_RESULT:
		opts, payload = s.convertTaskOption(t)
		if err = s.taskScheduler.NewWaitResultTask(t.GetTypeName(), payload, opts...); err != nil {
			s.log.Errorf("[%s] 创建等待结果任务失败[%s]", t.GetTypeName(), err.Error())
			return err
		}
	}

	return nil
}

// AsyncBackup 异步备份任务的实际执行逻辑。
//
// H8：实现真正的备份——导出核心业务表为 JSON，gzip 压缩后上传到 OSS 的 backups 桶。
// 纯 Go 实现，跨数据库驱动（MySQL/PostgreSQL/SQLite）。
// 当前覆盖核心身份/权限/组织表；审计日志等大体量表暂不纳入（可按需扩展 BackupRepo.ExportCoreTables）。
func (s *TaskService) AsyncBackup(taskType string, taskData *task.BackupTaskData) error {
	s.log.Infof("AsyncBackup [%s] [%+v] [%s]", taskType, taskData, taskData.Name)

	ctx := context.Background()
	backupName := ""
	if taskData != nil {
		backupName = taskData.Name
	}
	if backupName == "" {
		backupName = fmt.Sprintf("backup-%s", time.Now().UTC().Format("20060102-150405"))
	}

	if s.backupRepo == nil || !s.backupRepo.IsConfigured() || s.mc == nil {
		return fmt.Errorf("backup dependencies not configured (backupRepo or minio is nil)")
	}

	// 1. 导出核心表（ent 访问收敛在 data 层的 BackupRepo 内）
	tables := s.backupRepo.ExportCoreTables(ctx)

	// 2. 序列化为 JSON
	jsonBytes, err := json.MarshalIndent(map[string]any{
		"_meta": map[string]any{
			"exportedAt": time.Now().UTC().Format(time.RFC3339),
			"name":       backupName,
			"tables":     tableNamesOf(tables),
		},
		"data": tables,
	}, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal backup json failed: %w", err)
	}

	// 3. gzip 压缩
	compressed, err := gzipBytes(jsonBytes)
	if err != nil {
		return fmt.Errorf("gzip backup failed: %w", err)
	}

	// 4. 上传到 OSS
	objectName := fmt.Sprintf("%s/%s-%s.json.gz",
		time.Now().UTC().Format("2006/01/02"),
		backupName,
		time.Now().UTC().Format("150405"),
	)

	s.log.Infof("backup: uploading %s (%d bytes raw, %d bytes compressed) to bucket %q",
		objectName, len(jsonBytes), len(compressed), backupBucket)

	if _, _, _, err = s.mc.UploadFile(ctx, backupBucket, objectName, "application/gzip", compressed); err != nil {
		s.log.Errorf("backup: upload to oss failed: %s", err.Error())
		return fmt.Errorf("upload backup to oss failed: %w", err)
	}

	s.log.Infof("backup: completed successfully, object=%s", objectName)
	return nil
}

// tableNamesOf 返回 map 的键列表（用于备份元信息）。
func tableNamesOf(tables map[string]any) []string {
	names := make([]string, 0, len(tables))
	for k := range tables {
		names = append(names, k)
	}
	return names
}

// gzipBytes 使用标准库 gzip 压缩字节切片。
func gzipBytes(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	if _, err := gz.Write(data); err != nil {
		_ = gz.Close()
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
