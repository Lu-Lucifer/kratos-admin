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
  legend: {
    bottom: '2%',
    left: 'center',
  },
  series: [
    {
      animationDelay() {
        return Math.random() * 100;
      },
      animationEasing: 'exponentialInOut',
      animationType: 'scale',
      avoidLabelOverlap: false,
      color: getSeriesColors(),
      data: [
        { name: '搜索引擎', value: 1048 },
        { name: '直接访问', value: 735 },
        { name: '邮件营销', value: 580 },
        { name: '联盟广告', value: 484 },
      ],
      emphasis: {
        label: {
          fontSize: '12',
          fontWeight: 'bold',
          show: true,
        },
      },
      itemStyle: {
        // borderColor: '#fff',
        borderRadius: 10,
        borderWidth: 2,
      },
      label: {
        position: 'center',
        show: false,
      },
      labelLine: {
        show: false,
      },
      name: '访问来源',
      radius: ['40%', '65%'],
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
