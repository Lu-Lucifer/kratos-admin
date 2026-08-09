import { Card, theme } from 'antd';
import ReactECharts from 'echarts-for-react';
import { useMemo } from 'react';
import type { ActionDistributionResponse } from '@/api/generated/admin/service/v1';
import { useI18n } from '@/core/i18n';

interface SourceDonutChartProps {
  data?: ActionDistributionResponse;
}

/**
 * 操作类型分布环形图。
 * 后端返回 action 枯举名（CREATE/UPDATE/...），legend 直接展示枯举名。
 */
export const SourceDonutChart = ({ data }: SourceDonutChartProps) => {
  const { token } = theme.useToken();
  const { t } = useI18n('dashboard');

  // 调色板取 antd 运行时 token 的实际颜色值。ECharts 在 canvas 上渲染，
  // 无法解析 CSS 变量（var(--ant-color-*)），必须传入已解析的十六进制色值。
  const palette = [
    token.colorPrimary,
    token.colorSuccess,
    token.colorWarning,
    token.colorInfo,
  ];

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
            itemStyle: { color: palette[i % palette.length] },
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
