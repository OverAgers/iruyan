'use client';

import { useRouter } from 'next/navigation';
import AuthLayout from '@/components/layouts/auth-layout';
import UserRegisterForm from '@/features/register/components/user-register-form';

export default function RegisterPage() {
  const router = useRouter();

  const onSuccess = () => {
    router.push('/lobby');
  };

  return (
    <AuthLayout>
      <UserRegisterForm onSuccess={onSuccess} />
    </AuthLayout>
  );
}
