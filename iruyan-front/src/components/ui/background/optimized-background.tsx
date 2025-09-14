"use client";

import { ReactNode } from 'react';
import Image from 'next/image';
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
  alt,
  children,
  priority = false,
  className,
  style,
  sx
}: OptimizedBackgroundProps) {
  return (
    <Box
      className={className}
      sx={{
        position: 'relative',
        width: '100%',
        height: '100%',
        overflow: 'hidden',
        ...sx
      }}
      style={style}
    >
      {/* Next.js Image as background */}
      <Image
        src={src}
        alt={alt}
        fill
        priority={priority}
        style={{
          objectFit: 'cover',
          zIndex: -1,
        }}
        sizes="100vw"
        placeholder="blur"
        blurDataURL="data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAYEBQYFBAYGBQYHBwYIChAKCgkJChQODwwQFxQYGBcUFhYaHSUfGhsjHBYWICwgIyYnKSopGR8tMC0oMCUoKSj/2wBDAQcHBwoIChMKChMoGhYaKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCj/wAARCAAoACgDASIAAhEBAxEB/8QAFwABAQEBAAAAAAAAAAAAAAAABAUGA//EACgQAAIBAwMEAQUBAAAAAAAAAAECAwAEEQUSITFBBhNRImFxgZGhsf/EABUBAQEAAAAAAAAAAAAAAAAAAAMF/8QAGhEAAgMBAQAAAAAAAAAAAAAAAAECERIhEf/aAAwDAQACEQMRAD8A0tFFFaQCiiigAooooA//2Q=="
      />

      {/* Content */}
      <Box sx={{ position: 'relative', zIndex: 1, width: '100%', height: '100%' }}>
        {children}
      </Box>
    </Box>
  );
}