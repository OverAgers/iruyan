import { useState, useCallback } from "react";

export type LoginErrorCode = "401" | "400" | "500" | "UNKNOWN_ERROR";

export function useLoginErrorHandler() {
  const [errorMessage, setErrorMessage] = useState<string>("");
  const [showModal, setShowModal] = useState<boolean>(false);

  const handleError = useCallback((errorCode: string, originalErrorMessage: string) => {
    let newErrorMessage = "";
    
    switch (errorCode) {
      case "401":
        // 認証エラー
        newErrorMessage = "ユーザーIDまたはパスワードが正しくありません。\n確認してください。";
        break;
      case "400":
        // バリデーションエラー
        newErrorMessage = "入力内容に問題があります。\n確認してください。";
        break;
      case "500":
        // サーバーエラー
        newErrorMessage = "サーバーエラーが発生しました。\nしばらく時間をおいて再度お試しください。";
        break;
      default:
        // その他のエラー
        if (originalErrorMessage.includes("Network Error")) {
          newErrorMessage = "ネットワークエラーが発生しました。\nインターネット接続を確認してください。";
        } else {
          newErrorMessage = "ログインに失敗しました。\n再度お試しください。";
        }
        break;
    }
    
    // 状態更新を一括で行う
    setErrorMessage(newErrorMessage);
    setShowModal(true);
  }, []);

  const clearError = useCallback(() => {
    setErrorMessage("");
    setShowModal(false);
  }, []);

  const closeModal = useCallback(() => {
    setShowModal(false);
  }, []);

  return {
    errorMessage,
    showModal,
    handleError,
    clearError,
    closeModal,
  };
} 