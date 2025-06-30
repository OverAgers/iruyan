"use client";

import { useRouter } from "next/navigation";
import { useCallback, useEffect } from "react";
import AuthLayout from "@/components/layouts/auth-layout";
import UserLoginForm from "@/features/login/components/user-login-form";
import { useLoginErrorHandler } from "@/features/login/hooks/use-login-error-handler";
import ErrorModal from "@/components/ui/modal/ErrorModal";
import usePostLoginRequest from "@/features/login/api/post-login";
import { LoginForm } from "@/schema/login-form-schema";
import useUserStore from "@/stores/user-store";

export default function LoginPage() {
  const router = useRouter();
  const { data, login } = usePostLoginRequest();
  const { errorMessage, showModal, handleError, closeModal } = useLoginErrorHandler();
  const setUser = useUserStore((state) => state.setUser);

  const onSuccess = useCallback(async (formData: LoginForm) => {
    try {
      const result = await login(formData);
      
      // ユーザー情報をストアに保存
      setUser({
        iruyanId: result.user.iruyanId,
        name: result.user.userName,
        email: result.user.email,
        status: "idle",
        workTime: 0,
        restTime: 0,
        startTime: 0,
      });
    } catch (err: unknown) {
      // エラーレスポンスからステータスコードを取得
      const axiosError = err as { response?: { status?: number; data?: { message?: string } }; message?: string };
      const statusCode = axiosError.response?.status?.toString() || "UNKNOWN_ERROR";
      const originalMessage = axiosError.response?.data?.message || axiosError.message || "Unknown error";
      handleError(statusCode, originalMessage);
    }
  }, [login, handleError, setUser]);

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
