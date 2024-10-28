"use client"

import AuthLayout from "@/components/layouts/auth-layout";
import RegisterForm from "@/features/register/components/register-form";

export default function RegisterPage() {
  return (
    <AuthLayout>
      <RegisterForm />
    </AuthLayout>
  );
}