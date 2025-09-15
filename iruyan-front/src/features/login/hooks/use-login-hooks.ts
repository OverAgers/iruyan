import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useCallback } from "react";

import { LoginForm, loginFormSchema } from "@/schema/login-form-schema";
import { useRealtimeValidation } from "@/hooks/use-realtime-validation";

type Props = {
  onSubmit: (data: LoginForm) => void;
};

export default function UseLoginForm({ onSubmit }: Props) {
  const {
    register,
    handleSubmit,
    setValue: setFormValue,
    watch,
    getValues,
    formState: { errors: formErrors },
  } = useForm<LoginForm>({
    resolver: zodResolver(loginFormSchema),
    mode: "onBlur",
  });

  // リアルタイムバリデーションフック
  const {
    errors: realtimeErrors,
    handleFieldChange,
    validateAll,
    clearAllErrors,
    hasErrors,
  } = useRealtimeValidation(loginFormSchema, 300);

  // フィールド値設定（react-hook-formとリアルタイムバリデーション両方に対応）
  const setValue = useCallback((fieldName: keyof LoginForm, value: string) => {
    setFormValue(fieldName, value);
    const allValues = getValues();
    handleFieldChange(fieldName, value, allValues);
  }, [setFormValue, getValues, handleFieldChange]);

  // フォーム送信ハンドラ
  const onFormSubmit = useCallback((data: LoginForm) => {
    // 最終バリデーション実行
    if (validateAll(data)) {
      clearAllErrors();
      onSubmit(data);
    }
  }, [onSubmit, validateAll, clearAllErrors]);

  // エラー表示（リアルタイムエラーを優先、fallbackでform errors）
  const getFieldError = useCallback((fieldName: keyof LoginForm) => {
    return realtimeErrors[fieldName] || formErrors[fieldName]?.message;
  }, [realtimeErrors, formErrors]);

  return {
    register,
    handleSubmit,
    setValue,
    watch,
    getValues,
    errors: {
      iruyanId: getFieldError('iruyanId'),
      password: getFieldError('password'),
    },
    realtimeErrors,
    hasRealtimeErrors: hasErrors,
    onSubmit: handleSubmit(onFormSubmit),
    clearErrors: clearAllErrors,
  };
}
