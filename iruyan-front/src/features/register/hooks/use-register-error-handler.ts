import { useState, useCallback } from "react";

export type RegisterErrorCode = "409" | "400" | "500" | "UNKNOWN_ERROR";

export function useRegisterErrorHandler() {
  const [errorMessage, setErrorMessage] = useState<string>("");
  const [showModal, setShowModal] = useState<boolean>(false);

  const handleError = useCallback((errorCode: string, originalErrorMessage: string) => {
    let newErrorMessage = "";
    
    switch (errorCode) {
      case "409":
        // ユーザーID重複エラー
        newErrorMessage = "このユーザーIDは既に使用されています。\n別のユーザーIDを選択してください。";
        break;
      case "400":
        // バリデーションエラー
        if (originalErrorMessage.includes("already taken")) {
          newErrorMessage = "このユーザーIDは既に使用されています。\n別のユーザーIDを選択してください。";
        } else {
          newErrorMessage = "入力内容に問題があります。\n確認してください。";
        }
        break;
      case "500":
        // サーバーエラー
        newErrorMessage = "サーバーエラーが発生しました。\nしばらく時間をおいて再度お試しください。";
        break;
      default:
        // その他のエラー
        if (originalErrorMessage.includes("already taken")) {
          newErrorMessage = "このユーザーIDは既に使用されています。\n別のユーザーIDを選択してください。";
        } else {
          newErrorMessage = "登録に失敗しました。\n再度お試しください。";
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