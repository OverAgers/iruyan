"use client";

import { useEffect } from "react";
import AuthLayout from "@/components/layouts/auth-layout";
import UsePostRegisterRequest from "@/features/register/api/post-register";
import UserRegisterForm from "@/features/register/components/user-register-form";
import { RegisterForm } from "@/schema/register-form-schema";

export default function RegisterPage() {
  const { data, register } = UsePostRegisterRequest();

  const onSuccess = (data: RegisterForm) => {
    register(data);
  };

  useEffect(() => {
    if (data) {
      window.location.href = "/lobby";
    }
  }, [data]);

  return (
    <AuthLayout>
      <UserRegisterForm onSuccess={onSuccess} />
    </AuthLayout>
  );
}
