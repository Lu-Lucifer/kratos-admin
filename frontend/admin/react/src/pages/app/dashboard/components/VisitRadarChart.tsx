import { Card, theme } from 'antd';
import ReactECharts from 'echarts-for-react';
import { useI18n } from '@/core/i18n';

/**
 * 访问数量雷达图组件
 */
export const VisitRadarChart = () => {
  const { token } = theme.useToken();
  const { t } = useI18n('dashboard');

  const option = {
    tooltip: {
      trigger: 'item',
    },
    legend: {
      data: [t('data.visit'), t('data.trend')],
      bottom: 0,
      textStyle: {
        color: token.colorTextSecondary,
      },
    },
    radar: {
      indicator: [
        { name: t('devices.webpage'), max: 100 },
        { name: t('devices.other'), max: 100 },
        { name: t('devices.thirdParty'), max: 100 },
        { name: t('devices.client'), max: 100 },
        { name: t('devices.ipad'), max: 100 },
        { name: t('devices.mobile'), max: 100 },
      ],
      axisName: {
        color: token.colorTextSecondary,
      },
      splitLine: {
        lineStyle: {
          color: token.colorSplit,
        },
      },
      splitArea: {
        show: false,
      },
      axisLine: {
        lineStyle: {
          color: token.colorBorderSecondary,
        },
      },
    },
    series: [
      {
        name: t('charts.visitCount'),
        type: 'radar',
        data: [
          {
            value: [42, 30, 20, 35, 50, 60],
            name: t('data.visit'),
            itemStyle: {
              color: '#a78bfa',
            },
            areaStyle: {
              opacity: 0.3,
            },
          },
          {
            value: [50, 40, 35, 45, 60, 70],
            name: t('data.trend'),
            itemStyle: {
              color: '#3b82f6',
            },
            areaStyle: {
              opacity: 0.3,
            },
          },
        ],
      },
    ],
  };

  return (
    <Card title={t('charts.visitCount')} style={{ height: '100%' }}>
      <ReactECharts option={option} style={{ height: 280 }} />
    </Card>
  );
};
