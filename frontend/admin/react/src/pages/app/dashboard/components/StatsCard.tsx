import { Card } from 'antd';
import type { ReactNode } from 'react';

interface StatsCardProps {
  title: string;
  value: string;
  total: string;
  totalValue: string;
  icon: ReactNode;
}

/**
 * 统计卡片组件
 */
export const StatsCard = ({ title, value, total, totalValue, icon }: StatsCardProps) => {
  return (
    <Card style={{ height: '100%' }}>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start' }}>
        <div>
          <div
            style={{
              fontSize: 14,
              color: 'var(--ant-color-text-secondary)',
              marginBottom: 8,
            }}
          >
            {title}
          </div>
          <div style={{ fontSize: 24, fontWeight: 600, marginBottom: 16 }}>{value}</div>
          <div style={{ display: 'flex', justifyContent: 'space-between', fontSize: 12 }}>
            <span style={{ color: 'var(--ant-color-text-tertiary)' }}>
              {total}
            </span>
            <span style={{ color: 'var(--ant-color-text-secondary)' }}>
              {totalValue}
            </span>
          </div>
        </div>
        <div>{icon}</div>
      </div>
    </Card>
  );
};
