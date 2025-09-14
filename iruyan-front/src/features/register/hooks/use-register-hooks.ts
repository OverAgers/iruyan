import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useCallback } from "react";

import {
  RegisterForm,
  registerFormSchema,
} from "@/schema/register-form-schema";
import { useRealtimeValidation } from "@/hooks/use-realtime-validation";

type Props = {
  onSubmit: (data: RegisterForm) => void;
};

export default function UseRegisterForm({ onSubmit }: Props) {
  const {
    register,
    handleSubmit,
    setValue: setFormValue,
    watch,
    getValues,
    formState: { errors: formErrors },
  } = useForm<RegisterForm>({
    resolver: zodResolver(registerFormSchema),
    mode: "onBlur",
  });

  // リアルタイムバリデーションフック
  const {
    errors: realtimeErrors,
    handleFieldChange,
    validateAll,
    clearAllErrors,
    hasErrors,
  } = useRealtimeValidation(registerFormSchema, 300);

  // フィールド値設定（react-hook-formとリアルタイムバリデーション両方に対応）
  const setValue = useCallback((fieldName: keyof RegisterForm, value: string) => {
    setFormValue(fieldName, value);
    const allValues = getValues();
    handleFieldChange(fieldName, value, allValues);
  }, [setFormValue, getValues, handleFieldChange]);

  // フォーム送信ハンドラ
  const onFormSubmit = useCallback((data: RegisterForm) => {
    // 最終バリデーション実行
    if (validateAll(data)) {
      clearAllErrors();
      onSubmit(data);
    }
  }, [onSubmit, validateAll, clearAllErrors]);

  // エラー表示（リアルタイムエラーを優先、fallbackでform errors）
  const getFieldError = useCallback((fieldName: keyof RegisterForm) => {
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
      name: getFieldError('name'),
      email: getFieldError('email'),
      password: getFieldError('password'),
      passwordConfirm: getFieldError('passwordConfirm'),
    },
    realtimeErrors,
    hasRealtimeErrors: hasErrors,
    handleFormSubmit: handleSubmit(onFormSubmit),
    clearErrors: clearAllErrors,
  };
}
