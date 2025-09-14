"use client";

import { useState, useEffect, Suspense } from 'react';
import dynamic from 'next/dynamic';
import { Box, Skeleton } from '@mui/material';

// BarChart を動的インポート
const BarChart = dynamic(
  () => import('@mui/x-charts/BarChart').then(mod => ({ default: mod.BarChart })),
  {
    ssr: false,
    loading: () => (
      <Box sx={{ width: 400, height: 400 }}>
        <Skeleton variant="rectangular" width="100%" height="100%" />
      </Box>
    )
  }
);

interface ChartData {
  day: string;
  hours: number;
}

interface DynamicBarChartProps {
  data: ChartData[];
  width?: number;
  height?: number;
  color?: string;
}

export default function DynamicBarChart({
  data,
  width = 400,
  height = 400,
  color = "#D3AE6F"
}: DynamicBarChartProps) {
  const [isClient, setIsClient] = useState(false);

  useEffect(() => {
    setIsClient(true);
  }, []);

  if (!isClient) {
    return (
      <Box sx={{ width, height }}>
        <Skeleton variant="rectangular" width="100%" height="100%" />
      </Box>
    );
  }

  return (
    <Suspense
      fallback={
        <Box sx={{ width, height }}>
          <Skeleton variant="rectangular" width="100%" height="100%" />
        </Box>
      }
    >
      <BarChart
        series={[
          {
            data: data.map((item) => item.hours),
            color,
          },
        ]}
        height={height}
        xAxis={[
          {
            data: data.map((item) => item.day),
            scaleType: "band",
          },
        ]}
        margin={{ top: 61, bottom: 30, left: 40, right: 10 }}
      />
    </Suspense>
  );
}