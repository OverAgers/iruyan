"use client";

import { useState } from 'react';
import Image from 'next/image';
import { Box, Skeleton } from '@mui/material';

interface OptimizedAvatarProps {
  src?: string;
  alt: string;
  width: number;
  height: number;
  priority?: boolean;
  fallbackSrc?: string;
}

export default function OptimizedAvatar({
  src,
  alt,
  width,
  height,
  priority = false,
  fallbackSrc = '/icons/default-avatar.png'
}: OptimizedAvatarProps) {
  const [isLoading, setIsLoading] = useState(true);
  const [hasError, setHasError] = useState(false);

  const actualSrc = src || fallbackSrc;

  return (
    <Box
      sx={{
        position: 'relative',
        width: width,
        height: height,
        borderRadius: '50%',
        overflow: 'hidden',
        backgroundColor: '#f0f0f0'
      }}
    >
      {isLoading && (
        <Skeleton
          variant="circular"
          width={width}
          height={height}
          sx={{
            position: 'absolute',
            top: 0,
            left: 0,
            zIndex: 1
          }}
        />
      )}
      <Image
        src={hasError ? fallbackSrc : actualSrc}
        alt={alt}
        width={width}
        height={height}
        priority={priority}
        style={{
          width: '100%',
          height: '100%',
          objectFit: 'cover',
          borderRadius: '50%',
        }}
        onLoad={() => setIsLoading(false)}
        onError={() => {
          setHasError(true);
          setIsLoading(false);
        }}
        placeholder="blur"
        blurDataURL="data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAYEBQYFBAYGBQYHBwYIChAKCgkJChQODwwQFxQYGBcUFhYaHSUfGhsjHBYWICwgIyYnKSopGR8tMC0oMCUoKSj/2wBDAQcHBwoIChMKChMoGhYaKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCj/wAARCAAoACgDASIAAhEBAxEB/8QAFwABAQEBAAAAAAAAAAAAAAAABAUGA//EACgQAAIBAwMEAQUBAAAAAAAAAAECAwAEEQUSITFBBhNRImFxgZGhsf/EABUBAQEAAAAAAAAAAAAAAAAAAAMF/8QAGhEAAgMBAQAAAAAAAAAAAAAAAAECERIhEf/aAAwDAQACEQMRAD8A0tFFFaQCiiigAooooA//2Q=="
      />
    </Box>
  );
}