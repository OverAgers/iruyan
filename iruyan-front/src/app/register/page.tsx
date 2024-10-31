"use client";

import { useRouter } from "next/navigation";
import AuthLayout from "@/components/layouts/auth-layout";
import UserRegisterForm from "@/features/register/components/user-register-form";
import { UserInfo } from "@/types/user-info";

export default function RegisterPage() {
  const router = useRouter();

  const onSuccess = (data: UserInfo) => {
    router.push("/lobby");
  };

  return (
    <AuthLayout>
      <UserRegisterForm onSuccess={onSuccess} />
    </AuthLayout>
  );
}
