"use client";

import MainButton from "@/components/ui/button/main-button";
import AuthInputText from "@/components/ui/input/authorization-input-text";
import UseRegisterForm from "@/features/register/hooks/use-register-hooks";
import { RegisterForm } from "@/schema/register-form-schema";
import { UserInfo } from "@/types/user-info";
import UsePostRegisterRequest from "../api/post-register";
import useUserStore from "@/stores/user-store";

type Props = {
  onSuccess: (data: UserInfo) => void;
  onError?: (errorCode: string, errorMessage: string) => void;
};

export default function UserRegisterForm({ onSuccess, onError }: Props) {
  const setUser = useUserStore((state) => state.setUser);

  const {
    register: registerUser,
    isLoading,
  } = UsePostRegisterRequest();

  async function onSubmit(formData: RegisterForm): Promise<void> {
    try {
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
      const { passwordConfirm, ...dataToSubmit } = formData;
      const userData = await registerUser(dataToSubmit);

      const userInfo: UserInfo = {
        iruyanId: userData.user.iruyanId,
        name: userData.user.userName,
        email: userData.user.email,
        status: "idle",
        workTime: 0,
        restTime: 0,
        startTime: Date.now(),
      };

      setUser(userInfo);
      onSuccess(userInfo);
    } catch (error: unknown) {
      console.error("登録に失敗しました", error);

      let errorCode = "UNKNOWN_ERROR";
      let errorMessage = "登録に失敗しました";

      // Check for AxiosError first (since AxiosError extends Error)
      if (typeof error === 'object' && error !== null && 'response' in error && 'isAxiosError' in error) {
        // Handle axios error response
        const axiosError = error as {
          response?: {
            data?: {
              message?: string;
              error?: string;
              code?: string;
            };
            status?: number;
          };
          isAxiosError: boolean;
        };

        // Extract error message from response
        if (axiosError.response?.data?.message) {
          errorMessage = axiosError.response.data.message;
        } else if (axiosError.response?.data?.error) {
          errorMessage = axiosError.response.data.error;
        } else if (typeof axiosError.response?.data === 'string') {
          errorMessage = axiosError.response.data;
        }

        // Extract error code
        if (axiosError.response?.status) {
          errorCode = axiosError.response.status.toString();
        }
        if (axiosError.response?.data?.code) {
          errorCode = axiosError.response.data.code;
        }

        // Handle specific error messages
        if (errorMessage.includes("already taken")) {
          errorCode = "409";
        }
      } else if (error instanceof Error) {
        errorMessage = error.message;
      }

      // Call parent error handler if provided
      if (onError) {
        onError(errorCode, errorMessage);
      }
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
