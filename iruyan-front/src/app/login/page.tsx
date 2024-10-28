"use client"

import AuthLayout from "@/components/layouts/auth-layout";
import LoginForm from "@/features/login/components/login-form";

export default function LoginPage() {
  return (
    <AuthLayout>
      <LoginForm />
    </AuthLayout>
  );
}