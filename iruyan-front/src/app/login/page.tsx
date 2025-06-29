"use client";

import { useRouter } from "next/navigation";
import { useCallback, useEffect } from "react";
import AuthLayout from "@/components/layouts/auth-layout";
import UserLoginForm from "@/features/login/components/user-login-form";
import { useLoginErrorHandler } from "@/features/login/hooks/use-login-error-handler";
import ErrorModal from "@/components/ui/modal/ErrorModal";
import usePostLoginRequest from "@/features/login/api/post-login";
import { LoginForm } from "@/schema/login-form-schema";

export default function LoginPage() {
  const router = useRouter();
  const { data, error, login } = usePostLoginRequest();
  const { errorMessage, showModal, handleError, closeModal } = useLoginErrorHandler();

  const onSuccess = useCallback(async (formData: LoginForm) => {
    try {
      await login(formData);
    } catch (err: any) {
      // エラーレスポンスからステータスコードを取得
      const statusCode = err.response?.status?.toString() || "UNKNOWN_ERROR";
      const originalMessage = err.response?.data?.message || err.message || "Unknown error";
      handleError(statusCode, originalMessage);
    }
  }, [login, handleError]);

  useEffect(() => {
    if (data) {
      router.push("/lobby");
    }
  }, [data, router]);

  return (
    <AuthLayout>
      <UserLoginForm onSuccess={onSuccess} />
      
      <ErrorModal
        open={showModal}
        onClose={closeModal}
        errorMessage={errorMessage}
      />
    </AuthLayout>
  );
}
