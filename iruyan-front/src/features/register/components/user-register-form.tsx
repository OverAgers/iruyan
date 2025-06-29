"use client";

import MainButton from "@/components/ui/button/main-button";
import AuthInputText from "@/components/ui/input/authorization-input-text";
import UseRegisterForm from "@/features/register/hooks/use-register-hooks";
import { RegisterForm } from "@/schema/register-form-schema";
import { useState } from "react";
import { UserInfo } from "@/types/user-info";
import UsePostRegisterRequest from "../api/post-register";
import useUserStore from "@/stores/user-store";

type Props = {
  onSuccess: (data: UserInfo) => void;
};

export default function UserRegisterForm({ onSuccess }: Props) {
  const setUser = useUserStore((state) => state.setUser);
  const [errorMessage, setErrorMessage] = useState("");

  const {
    register: registerUser,
    isLoading,
    error: registerError,
  } = UsePostRegisterRequest();

  async function onSubmit(formData: RegisterForm): Promise<void> {
    try {
      const { passwordConfirm, ...dataToSubmit } = formData;
      const userData = await registerUser(dataToSubmit);

      const userInfo: UserInfo = {
        iruyanId: userData.user.iruyanId,
        name: userData.user.userName,
        email: userData.user.email,
        status: "idle", // 初期状態
        workTime: 0,
        restTime: 0,
        startTime: Date.now(), // number型
      };

      setUser(userInfo);
      onSuccess(userInfo);
    } catch (error: any) {
      console.error("登録に失敗しました", error);
      setErrorMessage(error.message || "登録に失敗しました");
      alert(errorMessage || "登録に失敗しました");
    }
  }

  const { errors, setValue, handleFormSubmit } = UseRegisterForm({
    onSubmit,
  });

  return (
    <form className="flex-col" onSubmit={handleFormSubmit}>
      <AuthInputText
        label="ユーザーID"
        placeholder="ユーザーID"
        onChange={(e) => setValue("iruyanId", e.target.value)}
        error={errors.iruyanId}
      />
      <AuthInputText
        label="ユーザー名（表示名）"
        placeholder="ユーザー名"
        onChange={(e) => setValue("name", e.target.value)}
        error={errors.name}
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
          title={isLoading ? "登録中..." : "新規登録"}
          type="submit"
          fullWidth={true}
          disabled={isLoading}
        />
      </div>
    </form>
  );
}
