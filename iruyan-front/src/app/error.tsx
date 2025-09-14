'use client';

import React from 'react';
import { Box, Typography, Button } from '@mui/joy';

interface ErrorPageProps {
  error: Error & { digest?: string };
  reset: () => void;
}

export default function ErrorPage({ error, reset }: ErrorPageProps) {
  React.useEffect(() => {
    // Log the error to an error reporting service
    console.error('Global error boundary caught an error:', error);
  }, [error]);

  return (
    <Box
      sx={{
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        justifyContent: 'center',
        minHeight: '100vh',
        padding: 4,
        textAlign: 'center',
      }}
    >
      <Typography level="h1" sx={{ mb: 2, color: 'danger.main' }}>
        エラーが発生しました
      </Typography>

      <Typography level="body-lg" sx={{ mb: 3, maxWidth: '600px' }}>
        申し訳ございません。予期しないエラーが発生しました。
        ページを再読み込みするか、しばらく時間をおいてから再度お試しください。
      </Typography>

      {process.env.NODE_ENV === 'development' && (
        <Box
          sx={{
            mb: 3,
            padding: 2,
            backgroundColor: 'neutral.100',
            borderRadius: 'md',
            maxWidth: '600px',
            textAlign: 'left',
          }}
        >
          <Typography level="body-sm" sx={{ fontFamily: 'monospace', wordBreak: 'break-all' }}>
            {error.message}
          </Typography>
        </Box>
      )}

      <Box sx={{ display: 'flex', gap: 2 }}>
        <Button
          color="primary"
          variant="solid"
          onClick={reset}
        >
          再試行
        </Button>

        <Button
          color="neutral"
          variant="outlined"
          onClick={() => window.location.href = '/'}
        >
          ホームに戻る
        </Button>
      </Box>
    </Box>
  );
}