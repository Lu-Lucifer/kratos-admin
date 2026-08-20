import React, { useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Form, Input, Button, App } from 'antd';
import { SafetyOutlined } from '@ant-design/icons';
import { useAuthStore } from '@/stores';

const MfaChallenge: React.FC = () => {
  const { t } = useTranslation('auth');
  const { completeMfaChallenge, loginLoading, mfaOperationId } = useAuthStore();
  const { message } = App.useApp();
  const [submitting, setSubmitting] = useState(false);

  // 退出 MFA 挑战：清状态并回登录页
  const cancelMfa = () => {
    useAuthStore.setState({ mfaOperationId: null });
    window.location.href = '/auth/login';
  };

  const handleSubmit = async (values: { totpCode: string }) => {
    if (!mfaOperationId) {
      message.error(t('mfaNotRequired'));
      window.location.href = '/auth/login';
      return;
    }
    setSubmitting(true);
    try {
      // completeMfaChallenge 内部会存 token + 跳首页，无需此处再跳。
      // 传 onSuccess 走完整登录成功路径。
      await completeMfaChallenge(values.totpCode);
    } catch (error: any) {
      message.error(error?.message || t('mfaVerifyFailed'));
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <div className="flex min-h-screen items-center justify-center bg-gradient-to-br from-gray-900 via-gray-800 to-gray-900 p-4">
      <div className="w-full max-w-md rounded-xl border border-white/10 bg-white/5 p-8 shadow-[0_8px_32px_rgba(0,0,0,0.3)] backdrop-blur-md light:border-gray-100 light:bg-white light:shadow-lg">
        <h2 className="mb-2 text-center text-xl font-semibold text-white light:text-gray-900">
          {t('mfaChallengeTitle')}
        </h2>
        <p className="mb-6 text-center text-sm text-gray-300 light:text-gray-500">
          {t('mfaChallengeDesc')}
        </p>
        <Form name="mfa-challenge" onFinish={handleSubmit} size="large">
          <Form.Item
            name="totpCode"
            rules={[
              { required: true, message: t('mfaCodeRequired') },
              {
                pattern: /^\d{6}$/,
                message: t('mfaCodeFormatError'),
              },
            ]}
          >
            <Input
              prefix={<SafetyOutlined />}
              placeholder={t('mfaCodePlaceholder')}
              autoComplete="one-time-code"
              inputMode="numeric"
              maxLength={6}
            />
          </Form.Item>
          <Form.Item>
            <Button
              type="primary"
              htmlType="submit"
              block
              loading={submitting || loginLoading}
            >
              {t('mfaVerify')}
            </Button>
          </Form.Item>
          <Form.Item>
            <Button block onClick={cancelMfa}>
              {t('mfaCancel')}
            </Button>
          </Form.Item>
        </Form>
      </div>
    </div>
  );
};

export default MfaChallenge;
