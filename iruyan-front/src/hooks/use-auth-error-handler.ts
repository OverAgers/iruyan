import { useEffect } from 'react';
import { useRouter } from 'next/navigation';
import useUserStore from '@/stores/user-store';

/**
 * 認証エラーを処理するカスタムフック
 * UnauthorizedErrorが発生した時にログイン画面にリダイレクト
 */
export function useAuthErrorHandler() {
  const router = useRouter();
  const { clearUser } = useUserStore();

  const handleAuthError = (error: unknown) => {
    if (error instanceof Error && error.name === 'UnauthorizedError') {
      // ユーザー状態をクリア
      clearUser();

      // ログイン画面にリダイレクト
      router.replace('/login');

      return true; // エラーが処理されたことを示す
    }

    return false; // このフックでは処理されなかった
  };

  return { handleAuthError };
}

/**
 * コンポーネントレベルで認証エラーを自動的に処理するフック
 */
export function useGlobalAuthErrorHandler() {
  const { handleAuthError } = useAuthErrorHandler();

  useEffect(() => {
    const handleUnhandledRejection = (event: PromiseRejectionEvent) => {
      if (handleAuthError(event.reason)) {
        event.preventDefault();
      }
    };

    window.addEventListener('unhandledrejection', handleUnhandledRejection);

    return () => {
      window.removeEventListener('unhandledrejection', handleUnhandledRejection);
    };
  }, [handleAuthError]);
}