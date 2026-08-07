/**
 * ECharts 系列色辅助：从 <html> 上的主题 CSS 变量读取色值并包裹为 hsl() 串。
 * 这些变量（--primary / --success / --warning / --destructive）由 vben 偏好系统
 * 在主题切换时更新，因此返回值随亮/暗模式自动变化，避免硬编码 hex 在暗黑下不协调。
 */
const THEME_VARS = [
  '--primary',
  '--success',
  '--warning',
  '--destructive',
] as const;

const readHsl = (varName: string): string => {
  const root = document.documentElement;
  const raw = getComputedStyle(root).getPropertyValue(varName).trim();
  // 变量值为 "H S% L%" 形式（无 hsl() 包裹），需手动包裹为合法 CSS 颜色。
  return raw ? `hsl(${raw})` : 'transparent';
};

/** 返回主题感知的系列色数组，用于 ECharts 多系列区分。 */
export const getSeriesColors = (): string[] =>
  THEME_VARS.map(readHsl);

/** 返回主题感知的单一强调色（用于单系列 / axisPointer 等）。 */
export const getAccentColor = (index: number): string => {
  const colors = getSeriesColors();
  return colors[index % colors.length] ?? 'transparent';
};
