import { Card, theme } from 'antd';
import ReactECharts from 'echarts-for-react';
import { useI18n } from '@/core/i18n';

/**
 * 访问来源饼图组件
 */
export const SourcePieChart = () => {
  const { token } = theme.useToken();
  const { t } = useI18n('dashboard');

  const option = {
    tooltip: {
      trigger: 'item',
      formatter: '{b}: {c} ({d}%)',
    },
    series: [
      {
        name: t('charts.visitSource'),
        type: 'pie',
        radius: ['20%', '70%'],
        center: ['50%', '50%'],
        roseType: 'area',
        itemStyle: {
          borderRadius: 8,
          borderColor: token.colorBgContainer,
          borderWidth: 2,
        },
        label: {
          color: token.colorText,
        },
        labelLine: {
          lineStyle: {
            color: token.colorTextSecondary,
          },
        },
        data: [
          { value: 40, name: t('categories.outsourcing'), itemStyle: { color: '#2dd4bf' } },
          { value: 30, name: t('categories.remote'), itemStyle: { color: '#5eead4' } },
          { value: 20, name: t('categories.customization'), itemStyle: { color: '#a78bfa' } },
          { value: 15, name: t('categories.technicalSupport'), itemStyle: { color: '#3b82f6' } },
        ],
      },
    ],
  };

  return (
    <Card title={t('charts.visitSource')} style={{ height: '100%' }}>
      <ReactECharts option={option} style={{ height: 280 }} />
    </Card>
  );
};
