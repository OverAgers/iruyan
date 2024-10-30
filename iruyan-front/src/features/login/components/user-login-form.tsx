"use client";

import MainButton from "@/components/ui/button/main-button";
import AuthInputText from "@/components/ui/input/authorization-input-text";
import UseLoginForm from "@/features/login/hooks/use-login-hooks";
import { LoginForm } from "@/schema/login-form-schema";

type Props = {
  onSuccess: (data: LoginForm) => void;
};

export default function UserLoginForm({ onSuccess }: Props) {
  const { errors, setValue, onSubmit } = UseLoginForm({
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
        label="パスワード"
        placeholder="パスワード"
        onChange={(e) => setValue("password", e.target.value)}
        error={errors.password}
        type="password"
      />
      <div>
        <MainButton
          component="button"
          title="入店する"
          type="submit"
          fullWidth={true}
        />
      </div>
    </form>
  );
}
