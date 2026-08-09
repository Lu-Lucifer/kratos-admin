import React, { useCallback, useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Form, Input, Button, Checkbox, App } from 'antd';
import { UserOutlined, LockOutlined, SafetyOutlined } from '@ant-design/icons';
import { useAuthStore } from '@/stores';
import { useNavigate, useSearchParams } from 'react-router-dom';
import { fetchGenerateCaptcha } from '@/api';

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
      // 弹出错误提示（与 CRUD 页面统一模式：优先用后端 message，兜底走 i18n）
      message.error(error?.message || t('loginFailed'));
    }
  };

  /**
   * 验证码图片组件，作为输入框的 addonAfter 承载。
   * 点击刷新验证码，与输入框视觉融为一体。
   */
  const captchaAddon = (
    <div
      className="flex items-center justify-center overflow-hidden w-[110px] h-9 rounded cursor-pointer border border-solid bg-[#f5f5f5] dark:border-zinc-700 dark:bg-zinc-800"
      title={t('captchaRefresh')}
      onClick={() => !captchaLoading && refreshCaptcha()}
    >
      {captchaImage ? (
        <img
          src={captchaImage}
          alt="captcha"
          className="h-full w-full object-cover"
        />
      ) : (
        <span className="text-[#999] text-xs dark:text-zinc-300">
          {captchaLoading ? '...' : t('captchaRefresh')}
        </span>
      )}
    </div>
  );

  return (
    <div className="w-full max-w-[420px]">
      {/* 标题 */}
      <div className="mb-11">
        <h2 className="text-[34px] font-extrabold tracking-[-0.5px] mb-2.5 text-[color:var(--ant-color-text)]">
          {t('welcomeBack')}
        </h2>
        <p className="text-[15px] leading-relaxed text-slate-400">
          {t('loginDescription')}
        </p>
      </div>

      {/* 登录表单 —— 卡片包裹，暗黑半透明 + 极细边框 + 圆角 + 柔和阴影 */}
      <div className="rounded-xl border border-white/10 bg-white/5 p-8 shadow-[0_8px_32px_rgba(0,0,0,0.3)] backdrop-blur-md light:border-black/10 light:bg-black/5 light:shadow-[0_8px_32px_rgba(0,0,0,0.08)]">
        <Form
          name="login"
          onFinish={handleSubmit}
          size="large"
          initialValues={{ remember: true }}
          className="login-form"
        >
          <Form.Item name="tenant_code" className="login-form-item">
            <Input
              prefix={<UserOutlined />}
              placeholder={t('tenantCodePlaceholder')}
              autoComplete="off"
            />
          </Form.Item>

          <Form.Item
            name="username"
            className="login-form-item"
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
            className="login-form-item"
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

          {/* 验证码 —— 融入输入框 addonAfter */}
          <Form.Item
            name="captcha"
            className="login-form-item"
            rules={[
              {
                required: true,
                message: t('captchaRequired'),
              },
            ]}
          >
            <Input
              prefix={<SafetyOutlined />}
              placeholder={t('captchaPlaceholder')}
              autoComplete="off"
              addonAfter={captchaAddon}
            />
          </Form.Item>

          <Form.Item className="login-remember-item">
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
              className="login-submit-btn h-11 rounded-lg"
            >
              {loginLoading ? t('loggingIn') : t('loginButton')}
            </Button>
          </Form.Item>
        </Form>
      </div>

      {/* 底部链接 */}
      <div className="mt-6 text-center text-[13px]">
        <span className="text-[color:var(--ant-color-text-secondary)]">
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
