import ReactECharts from 'echarts-for-react';
import { theme } from 'antd';
import { useI18n } from '@/core/i18n';

/**
 * 访问趋势折线图组件
 */
export const LineChart = () => {
  const { token } = theme.useToken();
  const { t } = useI18n('dashboard');

  const option = {
    tooltip: {
      trigger: 'axis',
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true,
    },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: [
        '6:00',
        '7:00',
        '8:00',
        '9:00',
        '10:00',
        '11:00',
        '12:00',
        '13:00',
        '14:00',
        '15:00',
        '16:00',
        '17:00',
        '18:00',
        '19:00',
        '20:00',
        '21:00',
        '22:00',
        '23:00',
      ],
      axisLine: {
        lineStyle: {
          color: token.colorBorderSecondary,
        },
      },
    },
    yAxis: {
      type: 'value',
      axisLine: {
        lineStyle: {
          color: token.colorBorderSecondary,
        },
      },
      splitLine: {
        lineStyle: {
          color: token.colorSplit,
        },
      },
    },
    series: [
      {
        name: t('data.visit'),
        type: 'line',
        smooth: true,
        areaStyle: {
          opacity: 0.3,
          color: token.colorPrimary,
        },
        lineStyle: {
          width: 2,
          color: token.colorPrimary,
        },
        itemStyle: {
          color: token.colorPrimary,
        },
        data: [
          0, 5000, 15000, 25000, 40000, 55000, 65000, 45000, 20000, 35000, 50000, 70000, 45000,
          25000, 15000, 8000, 4000, 2000,
        ],
      },
      {
        name: t('data.trend'),
        type: 'line',
        smooth: true,
        areaStyle: {
          opacity: 0.3,
          color: token.colorSuccess,
        },
        lineStyle: {
          width: 2,
          color: token.colorSuccess,
        },
        itemStyle: {
          color: token.colorSuccess,
        },
        data: [
          0, 1000, 3000, 8000, 15000, 20000, 22000, 15000, 8000, 12000, 18000, 23000, 15000, 8000,
          4000, 2000, 1000, 500,
        ],
      },
    ],
  };

  return <ReactECharts option={option} style={{ height: 300 }} />;
};
