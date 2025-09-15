import { useCallback, useState } from 'react';
import { ZodSchema, ZodError } from 'zod';

/**
 * デバウンス用のユーティリティ関数
 */
function debounce<T extends (...args: unknown[]) => void>(func: T, delay: number): T {
  let timeoutId: NodeJS.Timeout;
  return ((...args: Parameters<T>) => {
    clearTimeout(timeoutId);
    timeoutId = setTimeout(() => func(...args), delay);
  }) as T;
}

/**
 * リアルタイムバリデーション用のフック
 */
export function useRealtimeValidation<T extends Record<string, unknown>>(
  schema: ZodSchema<T>,
  debounceMs = 300
) {
  const [errors, setErrors] = useState<Partial<Record<keyof T, string>>>({});
  const [touchedFields, setTouchedFields] = useState<Set<keyof T>>(new Set());
  const [isValid, setIsValid] = useState(false);

  // バリデーション実行関数
  const validateField = useCallback((fieldName: keyof T, value: unknown, allValues: Partial<T>) => {
    try {
      // 特定のフィールドのみをバリデーション
      const testData = { ...allValues, [fieldName]: value } as T;
      schema.parse(testData);

      // エラーがない場合、そのフィールドのエラーを削除
      setErrors(prev => {
        const newErrors = { ...prev };
        delete newErrors[fieldName];
        return newErrors;
      });
    } catch (error) {
      if (error instanceof ZodError) {
        // 該当フィールドのエラーのみを抽出
        const fieldError = error.errors.find(err =>
          err.path.length > 0 && err.path[0] === fieldName
        );

        if (fieldError) {
          setErrors(prev => ({
            ...prev,
            [fieldName]: fieldError.message
          }));
        }
      }
    }
  }, [schema]);

  // デバウンス付きバリデーション
  const debouncedValidate = useCallback(
    (fieldName: keyof T, value: unknown, allValues: Partial<T>) => {
      const debounced = debounce(() => {
        validateField(fieldName, value, allValues);
      }, debounceMs);
      debounced();
    },
    [validateField, debounceMs]
  );

  // フィールド値の変更ハンドラ
  const handleFieldChange = useCallback((
    fieldName: keyof T,
    value: unknown,
    allValues: Partial<T>
  ) => {
    // フィールドをタッチ済みとしてマーク
    setTouchedFields(prev => new Set([...prev, fieldName]));

    // タッチ済みフィールドのみリアルタイムバリデーション実行
    if (touchedFields.has(fieldName) || value !== '') {
      debouncedValidate(fieldName, value, allValues);
    }
  }, [debouncedValidate, touchedFields]);

  // 全フィールドのバリデーション
  const validateAll = useCallback((values: T) => {
    try {
      schema.parse(values);
      setErrors({});
      setIsValid(true);
      return true;
    } catch (error) {
      if (error instanceof ZodError) {
        const formattedErrors: Partial<Record<keyof T, string>> = {};
        error.errors.forEach(err => {
          if (err.path.length > 0) {
            const fieldName = err.path[0] as keyof T;
            formattedErrors[fieldName] = err.message;
          }
        });
        setErrors(formattedErrors);
      }
      setIsValid(false);
      return false;
    }
  }, [schema]);

  // エラーをクリア
  const clearError = useCallback((fieldName: keyof T) => {
    setErrors(prev => {
      const newErrors = { ...prev };
      delete newErrors[fieldName];
      return newErrors;
    });
  }, []);

  // 全エラーをクリア
  const clearAllErrors = useCallback(() => {
    setErrors({});
    setTouchedFields(new Set());
    setIsValid(false);
  }, []);

  return {
    errors,
    touchedFields,
    isValid,
    handleFieldChange,
    validateAll,
    clearError,
    clearAllErrors,
    // 特定フィールドにエラーがあるかチェック
    hasError: (fieldName: keyof T) => Boolean(errors[fieldName]),
    // 全体的にエラーがあるかチェック
    hasErrors: Object.keys(errors).length > 0,
  };
}