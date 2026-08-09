import { defineConfig, presetWind3 } from 'unocss';

/**
 * UnoCSS 配置
 *
 * dark variant 重定向到 [data-theme="dark"] 属性选择器，与项目
 * ThemeProvider 在 <html> 上设置的 data-theme attribute 一致，
 * 使 dark: 前缀跟随项目的 class 切换策略而非默认的 media 策略。
 *
 * 注意：一旦存在本配置文件，UnoCSS 不再使用 vite 插件注入的默认
 * {presets:[presetWind3()]}，必须在此显式声明 presets，否则所有
 * 工具类失效。
 */
export default defineConfig({
  presets: [
    presetWind3({
      dark: {
        dark: '[data-theme="dark"]',
        light: '[data-theme="light"]',
      },
    }),
  ],
});
