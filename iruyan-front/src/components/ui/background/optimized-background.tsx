"use client";

import { ReactNode } from 'react';
import { Box } from '@mui/material';

interface OptimizedBackgroundProps {
  src: string;
  alt: string;
  children?: ReactNode;
  priority?: boolean;
  className?: string;
  style?: React.CSSProperties;
  sx?: Record<string, unknown>; // Material UI sx prop
}

export default function OptimizedBackground({
  src,
  children,
  className,
  style,
  sx,
}: OptimizedBackgroundProps) {
  return (
    <Box
      className={className}
      sx={{
        position: 'relative',
        width: '100%',
        height: '100vh',
        backgroundImage: `url(${src})`,
        backgroundSize: 'cover',
        backgroundPosition: 'center',
        backgroundRepeat: 'no-repeat',
      }}
      style={style}
    >
      {/* Content */}
      <Box sx={{ position: 'relative', zIndex: 1, width: '100%', height: '100%', ...sx }}>
        {children}
      </Box>
    </Box>
  );
}