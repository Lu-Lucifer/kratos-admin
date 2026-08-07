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

// 图表色随主题切换：每次渲染从 <html> 上的主题 CSS 变量读取。
const buildOption = (): any => ({
  legend: {
    bottom: 0,
    data: ['访问', '趋势'],
  },
  radar: {
    indicator: [
      {
        name: '网页',
      },
      {
        name: '移动端',
      },
      {
        name: 'Ipad',
      },
      {
        name: '客户端',
      },
      {
        name: '第三方',
      },
      {
        name: '其它',
      },
    ],
    radius: '60%',
    splitNumber: 8,
  },
  series: [
    {
      areaStyle: {
        opacity: 1,
        shadowBlur: 0,
        shadowColor: 'transparent',
        shadowOffsetX: 0,
        shadowOffsetY: 10,
      },
      data: [
        {
          itemStyle: {
            color: getAccentColor(0),
          },
          name: '访问',
          value: [90, 50, 86, 40, 50, 20],
        },
        {
          itemStyle: {
            color: getAccentColor(1),
          },
          name: '趋势',
          value: [70, 75, 70, 76, 20, 85],
        },
      ],
      itemStyle: {
        // borderColor: '#fff',
        borderRadius: 10,
        borderWidth: 2,
      },
      symbolSize: 0,
      type: 'radar',
    },
  ],
  tooltip: {},
});

const render = () => renderEcharts(buildOption());

onMounted(render);
watch(isDark, render);
</script>

<template>
  <EchartsUI ref="chartRef" />
</template>
