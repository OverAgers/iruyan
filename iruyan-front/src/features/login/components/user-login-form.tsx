"use client";

import useUserStore from "@/stores/user-store";
import MainButton from "@/components/ui/button/main-button";
import AuthInputText from "@/components/ui/input/authorization-input-text";
import UseLoginForm from "@/features/login/hooks/use-login-hooks";
import UseLoginRequest, {
  PostLoginRequest,
} from "@/features/login/api/post-login";
import { LoginForm } from "@/schema/login-form-schema";

type Props = {
  onSuccess?: (data: LoginForm) => void;
};

export default function UserLoginForm({ onSuccess }: Props) {
  const setUser = useUserStore((state) => state.setUser);
  const {
    errors,
    setValue,
    onSubmit: handleFormSubmit,
  } = UseLoginForm({
    onSubmit,
  });
  const { isLoading, login } = UseLoginRequest();

  // フォームの onSubmit 関数を定義
  async function onSubmit(formData: LoginForm) {
    try {
      // ログインリクエストを実行
      const userData = await login(formData as PostLoginRequest);

      // currentUser を更新
      setUser(userData);

      // 成功時のコールバックを呼び出す（必要であれば）
      if (onSuccess) {
        onSuccess(formData);
      }
    } catch (error) {
      console.error("ログインに失敗しました", error);
      if (error instanceof Error) {
        alert(error.message || "ログインに失敗しました");
      } else {
        alert("ログインに失敗しました");
      }
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
          fullWidth={true}
          disabled={isLoading}
        />
      </div>
    </form>
  );
}
