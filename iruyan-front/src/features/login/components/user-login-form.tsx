"use client";

import useUserStore from "@/stores/user-store";
import MainButton from "@/components/ui/button/main-button";
import AuthInputText from "@/components/ui/input/authorization-input-text";
import UseLoginForm from "@/features/login/hooks/use-login-hooks";
import UseLoginRequest, { PostLoginRequest } from "@/features/login/api/post-login";
import { LoginForm } from "@/schema/login-form-schema";
import type { StatusType } from "@/types/user-info";

type Props = {
  onSuccess?: (data: LoginForm) => void;
};

export default function UserLoginForm({ onSuccess }: Props) {
  const setUser = useUserStore((state) => state.setUser);
  const { errors, setValue, onSubmit: handleFormSubmit } = UseLoginForm({
    onSubmit,
  });
  const { isMutating, login } = UseLoginRequest();

  async function onSubmit(formData: LoginForm) {
    try {
      const userData = await login(formData as PostLoginRequest);

      const userInfo = {
        iruyanId: userData.user.iruyanId,
        name: userData.user.userName, // ← APIのkeyに合わせる
        email: userData.user.email,
        status: "working" as StatusType,
        workTime: 0,
        restTime: 0,
        startTime: Date.now(),
      };

      setUser(userInfo);

      onSuccess?.(formData);
    } catch (error: any) {
      console.error("ログインに失敗しました", error);

      // エラーの構造を見やすく出力
      if (error?.response) {
        console.log("🔴 error.response.data:", JSON.stringify(error.response.data, null, 2));
      } else if (error?.message) {
        console.log("🔴 error.message:", error.message);
      } else {
        console.log("🔴 error (raw):", JSON.stringify(error, null, 2));
      }

      const message =
        error?.response?.data?.message ??
        error?.message ??
        "ログインに失敗しました。";
      alert(message);
    }
  }

  return (
    <form className="flex-col" onSubmit={handleFormSubmit}>
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
          fullWidth
          disabled={isMutating}
        />
      </div>
    </form>
  );
}
