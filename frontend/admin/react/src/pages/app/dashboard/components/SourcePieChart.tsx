import { Card, theme } from 'antd';
import ReactECharts from 'echarts-for-react';
import { useMemo } from 'react';
import type { StatusDistributionResponse } from '@/api/generated/admin/service/v1';
import { useI18n } from '@/core/i18n';

interface SourcePieChartProps {
  data?: StatusDistributionResponse;
}

const PALETTE = [
  'var(--ant-color-primary)',
  'var(--ant-color-success)',
  'var(--ant-color-warning)',
  'var(--ant-color-info)',
];

/**
 * 登录成功/失败占比饼图。
 * 后端返回 status 枚举名（SUCCESS/FAILED/...），legend 直接展示枚举名。
 */
export const SourcePieChart = ({ data }: SourcePieChartProps) => {
  const { token } = theme.useToken();
  const { t } = useI18n('dashboard');

  const option = useMemo(() => {
    const items = data?.items ?? [];
    return {
      tooltip: {
        trigger: 'item',
        formatter: '{b}: {c} ({d}%)',
      },
      series: [
        {
          name: t('charts.loginStatusDistribution'),
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
    <Card title={t('charts.loginStatusDistribution')} style={{ height: '100%' }}>
      <ReactECharts option={option} style={{ height: 280 }} />
    </Card>
  );
};
