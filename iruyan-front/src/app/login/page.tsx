"use client";

import AuthLayout from "@/components/layouts/auth-layout";
import UserLoginForm from "@/features/login/components/user-login-form";
import { useEffect } from "react";
import usePostLoginRequest from "@/features/login/api/post-login";
import { LoginForm } from "@/schema/login-form-schema";

export default function LoginPage() {
  const { data, login } = usePostLoginRequest();

  const onSuccess = (formData: LoginForm) => {
    login(formData); // login関数にformデータを渡すだけ
  };

  useEffect(() => {
    if (data) {
      window.location.href = "/lobby";
    }
  }, [data]);

  return (
    <AuthLayout>
      <UserLoginForm onSuccess={onSuccess} />
    </AuthLayout>
  );
}
