'use client';

import AuthLayout from '@/components/layouts/auth-layout';
import UserLoginForm from '@/features/login/components/user-login-form';
import { useRouter } from 'next/navigation';

export default function LoginPage() {
  const router = useRouter();
  const onSuccess = () => {
    router.push('/lobby');
  };

  return (
    <AuthLayout>
      <UserLoginForm onSuccess={onSuccess} />
    </AuthLayout>
  );
}
