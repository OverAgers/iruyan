'use client';

import useUserStore from '@/stores/user-store';
import MainButton from '@/components/ui/button/main-button';
import AuthInputText from '@/components/ui/input/authorization-input-text';
import useLoginForm from '@/features/login/hooks/use-login-hooks';
import usePostLoginRequest from '@/features/login/api/post-login';
import { LoginForm } from '@/schema/login-form-schema';
import { useEffect } from 'react';

type Props = {
  onSuccess: () => void;
};

export default function UserLoginForm({ onSuccess }: Props) {
  const setUser = useUserStore((state) => state.setUser);
  const { isMutating, error, trigger } = usePostLoginRequest();
  const { errors, setValue, getValues, handleFormSubmit } = useLoginForm({
    onSubmit,
  });

  async function onSubmit() {
    try {
      const formData = getValues();
      const userData = await trigger(formData);
      if (userData) {
        setUser(userData.user);
      }
      if (onSuccess) {
        onSuccess();
      }
    } catch (error) {
      console.error('ログインに失敗しました', error);
      if (error instanceof Error) {
        alert(error.message || 'ログインに失敗しました');
      } else {
        alert('ログインに失敗しました');
      }
    }
  }

  return (
    <form className="flex-col" onSubmit={handleFormSubmit}>
      <AuthInputText
        label="ユーザーID"
        placeholder="ユーザーID"
        onChange={(e) => setValue('iruyanId', e.target.value)}
        error={errors.iruyanId}
      />
      <AuthInputText
        label="パスワード"
        placeholder="パスワード"
        onChange={(e) => setValue('password', e.target.value)}
        error={errors.password}
        type="password"
      />
      <div>
        <MainButton component="button" title="入店する" type="submit" fullWidth={true} disabled={isMutating} />
      </div>
    </form>
  );
}
