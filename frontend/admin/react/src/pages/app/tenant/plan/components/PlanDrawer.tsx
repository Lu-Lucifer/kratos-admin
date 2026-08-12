import { useRef, useState, useEffect } from 'react';
import type { ProFormInstance } from '@ant-design/pro-components';
import {
  DrawerForm,
  ProFormText,
  ProFormSelect,
  ProFormDigit,
  ProFormTextArea,
} from '@ant-design/pro-components';
import { App } from 'antd';
import { useQueryClient } from '@tanstack/react-query';
import { useTranslation } from 'react-i18next';

import { useCreatePlan, useUpdatePlan } from '@/api/hooks/plan';
import { getPlanVersionOptions, getExpiryPolicyOptions } from '../constants';
import { SELECT_FILTER_PROPS } from '../constants';

interface PlanDrawerProps {
  open: boolean;
  mode: 'create' | 'edit';
  data?: any;
  onClose: () => void;
  onSuccess: () => void;
}

/**
 * 套餐目录编辑/创建抽屉组件
 */
const PlanDrawer: React.FC<PlanDrawerProps> = ({
  open,
  mode,
  data,
  onClose,
  onSuccess,
}) => {
  const { t } = useTranslation('plan');
  const formRef = useRef<ProFormInstance>(null);
  const queryClient = useQueryClient();
  const { message } = App.useApp();

  const [confirmLoading, setConfirmLoading] = useState(false);

  // 编辑模式下设置表单值（destroyOnClose 时需延迟赋值）
  useEffect(() => {
    if (open && mode === 'edit' && data) {
      setTimeout(() => {
        formRef.current?.setFieldsValue({
          name: data.name || '',
          version: data.version,
          expiryPolicy: data.expiryPolicy,
          dataRetentionDays: data.dataRetentionDays,
          description: data.description || '',
          remark: data.remark || '',
        });
      }, 0);
    }
  }, [open, mode, data]);

  // 创建 mutation
  const createMutation = useCreatePlan({
    onSuccess: () => {
      message.success(t('createSuccess'));
      queryClient.invalidateQueries({ queryKey: ['listPlans'] });
      onSuccess();
      onClose();
    },
    onError: (error: Error) => {
      message.error(error.message || t('createFailed'));
    },
  });

  // 更新 mutation
  const updateMutation = useUpdatePlan({
    onSuccess: () => {
      message.success(t('updateSuccess'));
      queryClient.invalidateQueries({ queryKey: ['listPlans'] });
      onSuccess();
      onClose();
    },
    onError: (error: Error) => {
      message.error(error.message || t('updateFailed'));
    },
  });

  // 提交表单
  const handleSubmit = async (values: Record<string, any>) => {
    try {
      setConfirmLoading(true);
      if (mode === 'edit' && data?.id) {
        await updateMutation.mutateAsync({ id: data.id, values });
      } else {
        await createMutation.mutateAsync({ data: values });
      }
      return true;
    } catch {
      return false;
    } finally {
      setConfirmLoading(false);
    }
  };

  return (
    <DrawerForm
      formRef={formRef}
      title={mode === 'create' ? t('create') : t('edit')}
      open={open}
      onOpenChange={(visible) => {
        if (!visible) {
          formRef.current?.resetFields();
          onClose();
        }
      }}
      onFinish={handleSubmit}
      submitter={{
        searchConfig: {
          submitText: t('common:button.submit'),
          resetText: t('common:button.cancel'),
        },
        submitButtonProps: {
          loading: confirmLoading || createMutation.isPending || updateMutation.isPending,
        },
        resetButtonProps: { onClick: onClose },
      }}
      drawerProps={{ destroyOnClose: true, onClose, size: 600 }}
    >
      <ProFormText
        name="name"
        label={t('name')}
        placeholder={t('namePlaceholder')}
        rules={[{ required: true, message: t('requiredName') }]}
        fieldProps={{ allowClear: true }}
      />

      <ProFormSelect
        name="version"
        label={t('version')}
        placeholder={t('versionPlaceholder')}
        options={getPlanVersionOptions(t)}
        rules={[{ required: true, message: t('requiredVersion') }]}
        fieldProps={SELECT_FILTER_PROPS}
      />

      <ProFormSelect
        name="expiryPolicy"
        label={t('expiryPolicy')}
        placeholder={t('expiryPolicyPlaceholder')}
        options={getExpiryPolicyOptions(t)}
        rules={[{ required: true, message: t('requiredExpiryPolicy') }]}
        fieldProps={SELECT_FILTER_PROPS}
      />

      <ProFormDigit
        name="dataRetentionDays"
        label={t('dataRetentionDays')}
        placeholder={t('dataRetentionDaysPlaceholder')}
        fieldProps={{ precision: 0, min: 0 }}
      />

      <ProFormTextArea
        name="description"
        label={t('description')}
        placeholder={t('descriptionPlaceholder')}
        fieldProps={{ allowClear: true, autoSize: { minRows: 2, maxRows: 4 } }}
      />

      <ProFormTextArea
        name="remark"
        label={t('remark')}
        placeholder={t('remarkPlaceholder')}
        fieldProps={{ allowClear: true, autoSize: { minRows: 2, maxRows: 4 } }}
      />
    </DrawerForm>
  );
};

export default PlanDrawer;
