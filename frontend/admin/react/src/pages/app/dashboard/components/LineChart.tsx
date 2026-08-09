import ReactECharts from 'echarts-for-react';
import { theme } from 'antd';
import { useMemo } from 'react';
import type { LoginTrendResponse } from '@/api/generated/admin/service/v1';

interface LineChartProps {
  data?: LoginTrendResponse;
}

/**
 * 登录趋势折线图组件。
 * 数据由父组件从后端 GetLoginTrend 拉取后通过 prop 下发；
 * 后端已按日补零、升序返回 points。
 */
export const LineChart = ({ data }: LineChartProps) => {
  const { token } = theme.useToken();

  const option = useMemo(() => {
    const points = data?.points ?? [];
    return {
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
        data: points.map((p) => p.date),
        axisLine: {
          lineStyle: {
            color: token.colorBorderSecondary,
          },
        },
        axisTick: {
          lineStyle: {
            color: token.colorBorderSecondary,
          },
        },
        axisLabel: {
          color: token.colorTextSecondary,
        },
      },
      yAxis: {
        type: 'value',
        axisLine: {
          lineStyle: {
            color: token.colorBorderSecondary,
          },
        },
        axisTick: {
          lineStyle: {
            color: token.colorBorderSecondary,
          },
        },
        axisLabel: {
          color: token.colorTextSecondary,
        },
        splitLine: {
          lineStyle: {
            color: token.colorSplit,
          },
        },
      },
      series: [
        {
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
          data: points.map((p) => p.count),
        },
      ],
    };
  }, [data, token]);

  return <ReactECharts option={option} style={{ height: 300 }} />;
};
