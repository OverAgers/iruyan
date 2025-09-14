'use client';

import React from 'react';
import { useGlobalAuthErrorHandler } from '@/hooks/use-auth-error-handler';

interface GlobalErrorProviderProps {
  children: React.ReactNode;
}

/**
 * グローバルエラー処理を提供するプロバイダー
 * 認証エラーなどを自動的にハンドリング
 */
export function GlobalErrorProvider({ children }: GlobalErrorProviderProps) {
  // グローバル認証エラーハンドリングを有効化
  useGlobalAuthErrorHandler();

  return <>{children}</>;
}