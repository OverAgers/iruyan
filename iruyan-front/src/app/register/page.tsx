"use client";

import { useRouter } from "next/navigation";
import { useCallback } from "react";
import AuthLayout from "@/components/layouts/auth-layout";
import UserRegisterForm from "@/features/register/components/user-register-form";
import { useRegisterErrorHandler } from "@/features/register/hooks/use-register-error-handler";
import ErrorModal from "@/components/ui/modal/ErrorModal";

export default function RegisterPage() {
  const router = useRouter();
  const { errorMessage, showModal, handleError, closeModal } = useRegisterErrorHandler();

  const onSuccess = useCallback(() => {
    router.push("/lobby");
  }, [router]);

  return (
    <AuthLayout>
      <UserRegisterForm onSuccess={onSuccess} onError={handleError} />
      
      <ErrorModal
        open={showModal}
        onClose={closeModal}
        errorMessage={errorMessage}
      />
    </AuthLayout>
  );
}
