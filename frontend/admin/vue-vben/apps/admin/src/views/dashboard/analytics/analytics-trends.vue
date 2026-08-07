<script lang="ts" setup>
import { onMounted, ref, watch } from 'vue';

import { usePreferences } from '@vben/preferences';
import {
  EchartsUI,
  type EchartsUIType,
  useEcharts,
} from '@vben/plugins/echarts';

import { getAccentColor } from './chart-theme';

const chartRef = ref<EchartsUIType>();
const { renderEcharts } = useEcharts(chartRef);
const { isDark } = usePreferences();

// 图表色随主题切换：每次渲染从 <html> 上的主题 CSS 变量读取，
// 避免硬编码 hex 在暗黑下与背景不协调。
const buildOption = (): any => ({
  grid: {
    bottom: 0,
    containLabel: true,
    left: '1%',
    right: '1%',
    top: '2 %',
  },
  series: [
    {
      areaStyle: {},
      data: [
        111, 2000, 6000, 16_000, 33_333, 55_555, 64_000, 33_333, 18_000,
        36_000, 70_000, 42_444, 23_222, 13_000, 8000, 4000, 1200, 333, 222,
        111,
      ],
      itemStyle: {
        color: getAccentColor(0),
      },
      smooth: true,
      type: 'line',
    },
    {
      areaStyle: {},
      data: [
        33, 66, 88, 333, 3333, 6200, 20_000, 3000, 1200, 13_000, 22_000,
        11_000, 2221, 1201, 390, 198, 60, 30, 22, 11,
      ],
      itemStyle: {
        color: getAccentColor(1),
      },
      smooth: true,
      type: 'line',
    },
  ],
  tooltip: {
    axisPointer: {
      lineStyle: {
        color: getAccentColor(1),
        width: 1,
      },
    },
    trigger: 'axis',
  },
  xAxis: {
    axisTick: {
      show: false,
    },
    boundaryGap: false,
    data: Array.from({ length: 18 }).map((_item, index) => `${index + 6}:00`),
    splitLine: {
      lineStyle: {
        type: 'solid',
        width: 1,
      },
      show: true,
    },
    type: 'category',
  },
  yAxis: [
    {
      axisTick: {
        show: false,
      },
      max: 80_000,
      splitArea: {
        show: true,
      },
      splitNumber: 4,
      type: 'value',
    },
  ],
});

const render = () => renderEcharts(buildOption());

onMounted(render);
watch(isDark, render);
</script>

<template>
  <EchartsUI ref="chartRef" />
</template>
