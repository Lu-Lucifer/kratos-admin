import React, { useCallback, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Form, Input, Button, Checkbox, App } from 'antd';
import { UserOutlined, LockOutlined, SafetyOutlined } from '@ant-design/icons';
import { useAuthStore } from '@/stores';
import { useNavigate, useSearchParams } from 'react-router-dom';
import { fetchGenerateCaptcha } from '@/api';
import '../auth-form.style.less';

const Login: React.FC = () => {
  const { t } = useTranslation('auth');
  const { login, loginLoading } = useAuthStore();
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const { message } = App.useApp();

  // 验证码状态
  const [captchaId, setCaptchaId] = useState<string>('');
  const [captchaImage, setCaptchaImage] = useState<string>('');
  const [captchaLoading, setCaptchaLoading] = useState(false);

  // 获取验证码
  const refreshCaptcha = useCallback(async () => {
    setCaptchaLoading(true);
    try {
      const resp = await fetchGenerateCaptcha();
      setCaptchaId(resp.captchaId ?? '');
      setCaptchaImage(resp.imageBase64 ?? '');
    } catch {
      // 验证码获取失败不阻断页面，登录时会再次校验
    } finally {
      setCaptchaLoading(false);
    }
  }, []);

  useEffect(() => {
    // React 18 StrictMode 下 effect 会执行两次，导致两个并发的
    // fetchGenerateCaptcha 请求；后到的响应会覆盖先到的 captchaId，
    // 而后端通常一次性消费 captcha，登录时用的 captchaId 可能对应已被
    // 先到请求作废的 captcha。这里用 cancelled 标志位丢弃首次（被 double-invoke
    // 的第一次）请求的结果，确保 state 始终是最后一次请求的值。
    let cancelled = false;
    const run = async () => {
      setCaptchaLoading(true);
      try {
        const resp = await fetchGenerateCaptcha();
        if (cancelled) return;
        setCaptchaId(resp.captchaId ?? '');
        setCaptchaImage(resp.imageBase64 ?? '');
      } catch {
        // 验证码获取失败不阻断页面，登录时会再次校验
      } finally {
        if (!cancelled) setCaptchaLoading(false);
      }
    };
    run();
    return () => {
      cancelled = true;
    };
  }, []);

  const handleSubmit = async (values: {
    username: string;
    password: string;
    tenant_code?: string;
    remember?: boolean;
    captcha?: string;
  }) => {
    try {
      await login(
        {
          username: values.username,
          password: values.password,
          tenant_code: values.tenant_code,
          grant_type: 'password',
        },
        undefined,
        { id: captchaId, value: values.captcha ?? '' },
      );

      message.success(t('loginSuccess'));

      // 跳转到重定向页面或首页
      // 校验 redirect 必须为同源相对路径，防止开放重定向（如 ?redirect=https://evil.com 或 //evil.com）
      const rawRedirect = searchParams.get('redirect') || '/';
      const safeRedirect =
        typeof rawRedirect === 'string' &&
        rawRedirect.startsWith('/') &&
        !rawRedirect.startsWith('//')
          ? rawRedirect
          : '/';
      setTimeout(() => {
        navigate(safeRedirect);
      }, 300);
    } catch (error: any) {
      // 登录失败后刷新验证码
      refreshCaptcha();
    }
  };

  return (
    <div className="auth-form-container">
      {/* 标题 */}
      <div className="auth-form-header">
        <h2 className="auth-form-title">{t('welcomeBack')}</h2>
        <p className="auth-form-description">
          {t('loginDescription')}
        </p>
      </div>

      {/* 登录表单 */}
      <Form
        name="login"
        onFinish={handleSubmit}
        size="large"
        initialValues={{ remember: true }}
      >
        <Form.Item
          name="tenant_code"
          className="auth-form-item"
        >
          <Input
            prefix={<UserOutlined />}
            placeholder={t('tenantCodePlaceholder')}
            autoComplete="off"
          />
        </Form.Item>

        <Form.Item
          name="username"
          className="auth-form-item"
          rules={[
            {
              required: true,
              message: t('usernameRequired'),
            },
          ]}
        >
          <Input
            prefix={<UserOutlined />}
            placeholder={t('usernamePlaceholder')}
            autoComplete="username"
          />
        </Form.Item>

        <Form.Item
          name="password"
          className="auth-form-item"
          rules={[
            {
              required: true,
              message: t('passwordRequired'),
            },
          ]}
        >
          <Input.Password
            prefix={<LockOutlined />}
            placeholder={t('passwordPlaceholder')}
            autoComplete="current-password"
          />
        </Form.Item>

        {/* 验证码 */}
        <Form.Item
          name="captcha"
          className="auth-form-item"
          rules={[
            {
              required: true,
              message: t('captchaRequired'),
            },
          ]}
        >
          <div className="flex items-center gap-2">
            <Input
              prefix={<SafetyOutlined />}
              placeholder={t('captchaPlaceholder')}
              autoComplete="off"
              className="flex-1"
            />
            {/* 点击图片刷新验证码 */}
            <div
              className="flex items-center justify-center overflow-hidden border border-solid bg-[#f5f5f5] dark:border-zinc-700 dark:bg-zinc-800"
              title={t('captchaRefresh')}
              onClick={() => !captchaLoading && refreshCaptcha()}
              style={{
                height: 40,
                width: 120,
                flexShrink: 0,
                cursor: 'pointer',
                borderRadius: 6,
              }}
            >
              {captchaImage ? (
                <img
                  src={captchaImage}
                  alt="captcha"
                  style={{ height: '100%', width: '100%', objectFit: 'cover' }}
                />
              ) : (
                <span className="text-[#999] text-xs dark:text-zinc-300">
                  {captchaLoading ? '...' : t('captchaRefresh')}
                </span>
              )}
            </div>
          </div>
        </Form.Item>

        <Form.Item className="auth-remember-checkbox">
          <div className="flex items-center justify-between">
            <Form.Item name="remember" valuePropName="checked" noStyle>
              <Checkbox>{t('rememberAccount')}</Checkbox>
            </Form.Item>
          </div>
        </Form.Item>

        <Form.Item>
          <Button
            type="primary"
            htmlType="submit"
            loading={loginLoading}
            block
            className="auth-submit-button"
          >
            {loginLoading ? t('loggingIn') : t('loginButton')}
          </Button>
        </Form.Item>
      </Form>

      {/* 底部链接 */}
      <div className="auth-footer-link">
        <span className="auth-footer-text">
          {t('noAccount')}{' '}
        </span>
        <a
          href="/auth/register"
          className="text-sky-400 hover:text-sky-300 dark:text-sky-400 dark:hover:text-sky-300"
        >
          {t('createAccount')}
        </a>
      </div>
    </div>
  );
};

export default Login;
