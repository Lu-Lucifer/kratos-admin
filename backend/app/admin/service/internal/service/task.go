package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
	"github.com/tx7do/go-utils/trans"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"github.com/tx7do/kratos-transport/broker"
	asynqServer "github.com/tx7do/kratos-transport/transport/asynq"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"
	"kratos-admin/app/admin/service/internal/middleware/auth"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	systemV1 "kratos-admin/api/gen/go/system/service/v1"

	"kratos-admin/pkg/task"
)

type TaskService struct {
	adminV1.TaskServiceHTTPServer

	log *log.Helper

	Server *asynqServer.Server

	userRepo *data.UserRepo
	taskRepo *data.TaskRepo
}

func NewTaskService(
	logger log.Logger,
	taskRepo *data.TaskRepo,
	userRepo *data.UserRepo,
) *TaskService {
	l := log.NewHelper(log.With(logger, "module", "task/service/admin-service"))
	return &TaskService{
		log:      l,
		taskRepo: taskRepo,
		userRepo: userRepo,
	}
}

func (s *TaskService) ListTask(ctx context.Context, req *pagination.PagingRequest) (*systemV1.ListTaskResponse, error) {
	return s.taskRepo.List(ctx, req)
}

func (s *TaskService) GetTask(ctx context.Context, req *systemV1.GetTaskRequest) (*systemV1.Task, error) {
	return s.taskRepo.Get(ctx, req.GetId())
}

func (s *TaskService) GetTaskByTypeName(ctx context.Context, req *systemV1.GetTaskByTypeNameRequest) (*systemV1.Task, error) {
	return s.taskRepo.GetByTypeName(ctx, req.GetTypeName())
}

func (s *TaskService) CreateTask(ctx context.Context, req *systemV1.CreateTaskRequest) (*emptypb.Empty, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, adminV1.ErrorAccessForbidden("用户认证失败")
	}

	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = trans.Ptr(authInfo.UserId)

	var t *systemV1.Task
	t, err = s.taskRepo.Create(ctx, req)
	if err != nil {

		return nil, err
	}

	if err = s.startTask(t); err != nil {
		s.log.Error(err)
	}

	return &emptypb.Empty{}, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, req *systemV1.UpdateTaskRequest) (*emptypb.Empty, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, adminV1.ErrorAccessForbidden("用户认证失败")
	}

	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = trans.Ptr(authInfo.UserId)

	var t *systemV1.Task
	t, err = s.taskRepo.Update(ctx, req)
	if err != nil {

		return nil, err
	}

	if err = s.startTask(t); err != nil {
		s.log.Error(err)
	}

	return &emptypb.Empty{}, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, req *systemV1.DeleteTaskRequest) (*emptypb.Empty, error) {
	authInfo, err := auth.FromContext(ctx)
	if err != nil {
		s.log.Errorf("用户认证失败[%s]", err.Error())
		return nil, adminV1.ErrorAccessForbidden("用户认证失败")
	}

	req.OperatorId = trans.Ptr(authInfo.UserId)

	var t *systemV1.Task
	t, err = s.taskRepo.Get(ctx, req.GetId())
	if err != nil {
		s.log.Error(err)
	}

	_, err = s.taskRepo.Delete(ctx, req)
	if err != nil {
		return nil, err
	}

	if t != nil {
		_ = s.stopTask(t)
	}

	return &emptypb.Empty{}, nil
}

// ControlTask 控制调度任务
func (s *TaskService) ControlTask(ctx context.Context, req *systemV1.ControlTaskRequest) (*emptypb.Empty, error) {
	t, err := s.taskRepo.GetByTypeName(ctx, req.GetTypeName())
	if err != nil {
		s.log.Errorf("获取任务失败[%s]", err.Error())
		return nil, err
	}

	switch req.GetControlType() {
	case systemV1.TaskControlType_ControlType_Restart:
		if err = s.stopTask(t); err != nil {
			return nil, err
		}

		if err = s.startTask(t); err != nil {
			return nil, err
		}

	case systemV1.TaskControlType_ControlType_Stop:
		err = s.stopTask(t)
		return nil, err

	case systemV1.TaskControlType_ControlType_Start:
		err = s.startTask(t)
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// StopAllTask 停止所有的调度任务
func (s *TaskService) StopAllTask(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	s.stopAllTask()
	return &emptypb.Empty{}, nil
}

// RestartAllTask 重启所有的调度任务
func (s *TaskService) RestartAllTask(ctx context.Context, _ *emptypb.Empty) (*systemV1.RestartAllTaskResponse, error) {
	// 停止所有的任务
	s.stopAllTask()

	// 重新启动所有的任务
	count, err := s.StartAllTask(ctx)

	return &systemV1.RestartAllTaskResponse{
		Count: count,
	}, err
}

// StartAllTask 启动所有的任务
func (s *TaskService) StartAllTask(ctx context.Context) (int32, error) {
	//_, _ = s.Server.NewPeriodicTask("*/1 * * * ?", task.BackupTaskType, task.BackupTaskData{Name: "test"})

	resp, err := s.ListTask(ctx, &pagination.PagingRequest{
		NoPaging: trans.Ptr(true),
		Query:    trans.Ptr(""),
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
	s.log.Infof("开始清除所有的定时任务...")

	// 清除所有的定时任务
	s.Server.RemoveAllPeriodicTask()

	s.log.Infof("完成清除所有的定时任务")
}

// stopTask 停止一个任务
func (s *TaskService) stopTask(t *systemV1.Task) error {
	if t == nil {
		return errors.New("task is nil")
	}

	if t.GetEnable() == false {
		return errors.New("task is not enable")
	}

	switch t.GetType() {
	case systemV1.TaskType_TaskType_Periodic:
		return s.Server.RemovePeriodicTask(t.GetTypeName())

	case systemV1.TaskType_TaskType_Delay:

	case systemV1.TaskType_TaskType_WaitResult:
	}

	return nil
}

// convertTaskOption 转换任务选项
func (s *TaskService) convertTaskOption(t *systemV1.Task) (opts []asynq.Option, payload broker.Any) {
	if t == nil {
		return
	}

	if len(t.GetTaskPayload()) > 0 {
		_ = json.Unmarshal([]byte(t.GetTaskPayload()), &payload)
	}

	if t.GetRetryCount() > 0 {
		opts = append(opts, asynq.MaxRetry(int(t.GetRetryCount())))
	}
	if t.GetTimeout() != nil {
		opts = append(opts, asynq.Timeout(t.Timeout.AsDuration()))
	}
	if t.GetDeadline() != nil {
		opts = append(opts, asynq.Deadline(t.GetDeadline().AsTime()))
	}
	if t.GetProcessIn() != nil {
		opts = append(opts, asynq.ProcessIn(t.GetProcessIn().AsDuration()))
	}
	if t.GetProcessAt() != nil {
		opts = append(opts, asynq.ProcessAt(t.GetProcessAt().AsTime()))
	}

	return
}

// startTask 启动一个任务
func (s *TaskService) startTask(t *systemV1.Task) error {
	if t == nil {
		return errors.New("task is nil")
	}

	if t.GetEnable() == false {
		return errors.New("task is not enable")
	}

	var opts []asynq.Option
	var payload broker.Any
	var err error

	switch t.GetType() {
	case systemV1.TaskType_TaskType_Periodic:
		opts, payload = s.convertTaskOption(t)
		if _, err = s.Server.NewPeriodicTask(t.GetCronSpec(), t.GetTypeName(), payload, opts...); err != nil {
			s.log.Errorf("[%s] 创建定时任务失败[%s]", t.GetTypeName(), err.Error())
			return err
		}

	case systemV1.TaskType_TaskType_Delay:
		opts, payload = s.convertTaskOption(t)
		if err = s.Server.NewTask(t.GetTypeName(), payload, opts...); err != nil {
			s.log.Errorf("[%s] 创建延迟任务失败[%s]", t.GetTypeName(), err.Error())
			return err
		}

	case systemV1.TaskType_TaskType_WaitResult:
		opts, payload = s.convertTaskOption(t)
		if err = s.Server.NewWaitResultTask(t.GetTypeName(), payload, opts...); err != nil {
			s.log.Errorf("[%s] 创建等待结果任务失败[%s]", t.GetTypeName(), err.Error())
			return err
		}
	}

	return nil
}

// AsyncBackup 异步备份
func (s *TaskService) AsyncBackup(taskType string, taskData *task.BackupTaskData) error {
	s.log.Infof("AsyncBackup [%s] [%+v] [%s]", taskType, taskData, taskData.Name)
	return nil
}
