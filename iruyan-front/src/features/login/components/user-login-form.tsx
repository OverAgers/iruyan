"use client";

import React, { useState } from "react";
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
  const [loginError, setLoginError] = useState<string | null>(null); // ← エラー表示用 state
  const [formValues, setFormValues] = useState({ iruyanId: '', password: '' }); // フォーム値の追跡

  const { errors, setValue, onSubmit: handleFormSubmit } = UseLoginForm({ onSubmit });
  const { isMutating, login } = UseLoginRequest();

  // フィールド値変更ハンドラ
  const handleFieldChange = (fieldName: keyof LoginForm, value: string) => {
    setValue(fieldName, value);
    setFormValues(prev => ({ ...prev, [fieldName]: value }));
    // ログインエラーをクリア（ユーザーが入力を開始したら）
    if (loginError) {
      setLoginError(null);
    }
  };

  async function onSubmit(formData: LoginForm) {
    try {
      setLoginError(null);
      const userData = await login(formData as PostLoginRequest);

      const userInfo = {
        iruyanId: userData.user.iruyanId,
        name: userData.user.userName,
        email: userData.user.email,
        status: "working" as StatusType,
        workTime: 0,
        restTime: 0,
        startTime: Date.now(),
      };

      setUser(userInfo);
      onSuccess?.(formData);
    } catch (error: unknown) {
      console.error("ログインに失敗しました", error);

      const errorMessage =
        (error as Error)?.message || 
        (error as { response?: { data?: { message?: string } } })?.response?.data?.message || 
        "";

      if (errorMessage.includes("user not found")) {
        setLoginError("ユーザーIDが存在しません");
      } else if (errorMessage.includes("invalid password")) {
        setLoginError("パスワードが正しくありません");
      } else {
        setLoginError("IDまたはパスワードが異なります");
      }
    }
  }

  return (
    <form className="flex-col" onSubmit={handleFormSubmit}>
      <AuthInputText
        label="ユーザーID"
        placeholder="ユーザーID"
        onChange={(e) => handleFieldChange("iruyanId", e.target.value)}
        error={errors.iruyanId}
        type="text"
        isValid={!errors.iruyanId && Boolean(formValues.iruyanId)}
        showSuccess={true}
      />
      <AuthInputText
        label="パスワード"
        placeholder="パスワード"
        onChange={(e) => handleFieldChange("password", e.target.value)}
        error={errors.password}
        type="password"
        isValid={!errors.password && Boolean(formValues.password)}
        showSuccess={true}
      />

      {/* ログインエラー表示 */}
      {loginError && (
        <p className="text-red-600 text-sm mt-2 mb-1">{loginError}</p>
      )}

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
