"use client";

import MainButton from "@/components/ui/button/main-button";
import AuthInputText from "@/components/ui/input/authorization-input-text";
import UseRegisterForm from "@/features/register/hooks/use-register-hooks";
import { RegisterForm } from "@/schema/register-form-schema";

type Props = {
  onSuccess: (data: RegisterForm) => void;
};

export default function UserRegisterForm({ onSuccess }: Props) {
  const { errors, setValue, onSubmit } = UseRegisterForm({
    onSubmit: onSuccess,
  });

  return (
    <form className="flex-col" onSubmit={onSubmit}>
      <AuthInputText
        label="ユーザーID"
        placeholder="ユーザーID"
        onChange={(e) => setValue("iruyanId", e.target.value)}
        error={errors.iruyanId}
      />
      <AuthInputText
        label="ユーザー名（表示名）"
        placeholder="ユーザー名"
        onChange={(e) => setValue("userName", e.target.value)}
        error={errors.userName}
      />
      <AuthInputText
        label="メールアドレス"
        placeholder="メールアドレス"
        onChange={(e) => setValue("email", e.target.value)}
        error={errors.email}
      />
      <AuthInputText
        label="パスワード"
        placeholder="パスワード"
        onChange={(e) => setValue("password", e.target.value)}
        error={errors.password}
        type="password"
      />
      <AuthInputText
        label="パスワード確認"
        placeholder="パスワード確認"
        onChange={(e) => setValue("passwordConfirm", e.target.value)}
        error={errors.passwordConfirm}
        type="password"
      />
      <div>
        <MainButton
          component="button"
          title="新規登録"
          type="submit"
          fullWidth={true}
        />
      </div>
    </form>
  );
}
