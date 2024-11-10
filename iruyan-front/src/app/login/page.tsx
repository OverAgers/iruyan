'use client';

import AuthLayout from '@/components/layouts/auth-layout';
import UserLoginForm from '@/features/login/components/user-login-form';
import { useEffect } from 'react';
import UsePostLoginRequest from '@/features/login/api/post-login';
import { LoginForm } from '@/schema/login-form-schema';

export default function LoginPage() {
  const { data, login } = UsePostLoginRequest();

  const onSuccess = (data: LoginForm) => {
    login(data);
  };

  useEffect(() => {
    if (data) {
      window.location.href = '/lobby';
    }
  }, [data]);

  return (
    <AuthLayout>
      <UserLoginForm onSuccess={onSuccess} />
    </AuthLayout>
  );
}
