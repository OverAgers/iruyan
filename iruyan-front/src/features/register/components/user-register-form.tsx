'use client';

import MainButton from '@/components/ui/button/main-button';
import AuthInputText from '@/components/ui/input/authorization-input-text';
import useRegisterForm from '@/features/register/hooks/use-register-hooks';
import usePostRegisterRequest from '../api/post-register';
import useUserStore from '@/stores/user-store';

type Props = {
  onSuccess: () => void;
};

export default function UserRegisterForm({ onSuccess }: Props) {
  const setUser = useUserStore((state) => state.setUser);
  const { isMutating, error, trigger } = usePostRegisterRequest();
  const { errors, setValue, handleFormSubmit, getValues } = useRegisterForm({
    onSubmit,
  });

  async function onSubmit(): Promise<void> {
    try {
      const { passwordConfirm, ...formData } = getValues();
      const userData = await trigger(formData);
      if (userData) {
        setUser(userData.user);
        onSuccess();
      }
    } catch (error: any) {
      const message = error.message || '登録に失敗しました';
      console.error('登録に失敗しました', message);
      alert(message);
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
        label="ユーザー名（表示名）"
        placeholder="ユーザー名"
        onChange={(e) => setValue('userName', e.target.value)}
        error={errors.userName}
      />
      <AuthInputText
        label="メールアドレス"
        placeholder="メールアドレス"
        onChange={(e) => setValue('email', e.target.value)}
        error={errors.email}
      />
      <AuthInputText
        label="パスワード"
        placeholder="パスワード"
        onChange={(e) => setValue('password', e.target.value)}
        error={errors.password}
        type="password"
      />
      <AuthInputText
        label="パスワード確認"
        placeholder="パスワード確認"
        onChange={(e) => setValue('passwordConfirm', e.target.value)}
        error={errors.passwordConfirm}
        type="password"
      />
      <div>
        <MainButton
          component="button"
          title={isMutating ? '登録中...' : '新規登録'}
          type="submit"
          fullWidth={true}
          disabled={isMutating}
        />
      </div>
    </form>
  );
}
