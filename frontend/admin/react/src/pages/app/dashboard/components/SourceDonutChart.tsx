import { Card, theme } from 'antd';
import ReactECharts from 'echarts-for-react';
import { useMemo } from 'react';
import type { ActionDistributionResponse } from '@/api/generated/admin/service/v1';
import { useI18n } from '@/core/i18n';

interface SourceDonutChartProps {
  data?: ActionDistributionResponse;
}

const PALETTE = [
  'var(--ant-color-primary)',
  'var(--ant-color-success)',
  'var(--ant-color-warning)',
  'var(--ant-color-info)',
];

/**
 * 操作类型分布环形图。
 * 后端返回 action 枚举名（CREATE/UPDATE/...），legend 直接展示枚举名。
 */
export const SourceDonutChart = ({ data }: SourceDonutChartProps) => {
  const { token } = theme.useToken();
  const { t } = useI18n('dashboard');

  const option = useMemo(() => {
    const items = data?.items ?? [];
    return {
      tooltip: {
        trigger: 'item',
        formatter: '{b}: {c} ({d}%)',
      },
      legend: {
        orient: 'horizontal',
        bottom: 0,
        data: items.map((it) => it.label),
        textStyle: {
          color: token.colorTextSecondary,
        },
      },
      series: [
        {
          name: t('charts.operationActionDistribution'),
          type: 'pie',
          radius: ['40%', '70%'],
          avoidLabelOverlap: false,
          itemStyle: {
            borderRadius: 10,
            borderColor: token.colorBgContainer,
            borderWidth: 2,
          },
          label: {
            show: false,
            position: 'center',
          },
          emphasis: {
            label: {
              show: true,
              fontSize: 20,
              fontWeight: 'bold',
            },
          },
          labelLine: {
            show: false,
          },
          data: items.map((it, i) => ({
            value: it.count,
            name: it.label,
            itemStyle: { color: PALETTE[i % PALETTE.length] },
          })),
        },
      ],
    };
  }, [data, token, t]);

  return (
    <Card title={t('charts.operationActionDistribution')} style={{ height: '100%' }}>
      <ReactECharts option={option} style={{ height: 280 }} />
    </Card>
  );
};
