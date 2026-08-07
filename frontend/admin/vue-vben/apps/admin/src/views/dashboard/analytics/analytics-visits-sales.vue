<script lang="ts" setup>
import { onMounted, ref, watch } from 'vue';

import { usePreferences } from '@vben/preferences';
import {
  EchartsUI,
  type EchartsUIType,
  useEcharts,
} from '@vben/plugins/echarts';

import { getSeriesColors } from './chart-theme';

const chartRef = ref<EchartsUIType>();
const { renderEcharts } = useEcharts(chartRef);
const { isDark } = usePreferences();

// 图表色随主题切换：每次渲染从 <html> 上的主题 CSS 变量读取。
const buildOption = (): any => ({
  series: [
    {
      animationDelay() {
        return Math.random() * 400;
      },
      animationEasing: 'exponentialInOut',
      animationType: 'scale',
      center: ['50%', '50%'],
      color: getSeriesColors(),
      data: [
        { name: '外包', value: 500 },
        { name: '定制', value: 310 },
        { name: '技术支持', value: 274 },
        { name: '远程', value: 400 },
      ].sort((a, b) => {
        return a.value - b.value;
      }),
      name: '商业占比',
      radius: '80%',
      roseType: 'radius',
      type: 'pie',
    },
  ],

  tooltip: {
    trigger: 'item',
  },
});

const render = () => renderEcharts(buildOption());

onMounted(render);
watch(isDark, render);
</script>

<template>
  <EchartsUI ref="chartRef" />
</template>
