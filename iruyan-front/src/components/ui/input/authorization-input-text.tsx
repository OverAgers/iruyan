import { FormControl, FormHelperText, FormLabel, Input } from "@mui/joy";
import CheckIcon from '@mui/icons-material/Check';
import { FieldError } from "react-hook-form";
import { useState } from "react";

type Props = {
  label: string;
  placeholder: string;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  defaultValue?: string;
  endDecorator?: React.ReactNode;
  error?: FieldError | string | undefined;
  type?: "text" | "password";
  isValid?: boolean; // 成功状態の表示用
  showSuccess?: boolean; // 成功状態を表示するか
};

export default function AuthInputText(props: Props) {
  const [isFocused, setIsFocused] = useState(false);

  // エラーメッセージの統一化
  const errorMessage = typeof props.error === 'string'
    ? props.error
    : props.error?.message;

  const hasError = Boolean(errorMessage);
  const showSuccessIndicator = props.showSuccess && props.isValid && !hasError && !isFocused;

  // 動的なスタイル設定
  const getInputStyles = () => {
    const baseStyles = {
      border: "2px solid",
      borderRadius: "8px",
      boxShadow: "none",
      transition: "all 0.2s ease-in-out",
    };

    if (hasError) {
      return {
        ...baseStyles,
        borderColor: "#d32f2f",
        "&:hover": { borderColor: "#d32f2f" },
        "&:focus-within": { borderColor: "#d32f2f" }
      };
    } else if (showSuccessIndicator) {
      return {
        ...baseStyles,
        borderColor: "#2e7d32",
        "&:hover": { borderColor: "#2e7d32" },
        "&:focus-within": { borderColor: "#1976d2" }
      };
    } else if (isFocused) {
      return {
        ...baseStyles,
        borderColor: "#1976d2",
        "&:hover": { borderColor: "#1976d2" }
      };
    } else {
      return {
        ...baseStyles,
        borderColor: "transparent",
        backgroundColor: "#f5f5f5",
        "&:hover": { borderColor: "#e0e0e0" },
        "&:focus-within": { borderColor: "#1976d2", backgroundColor: "#fff" }
      };
    }
  };

  return (
    <FormControl error={hasError} sx={{ mb: 2 }}>
      <FormLabel
        sx={{
          mb: 0.5,
          fontSize: "lg",
          fontWeight: "bold",
          color: hasError ? "#d32f2f" : "#3C2800",
          transition: "color 0.2s ease-in-out"
        }}
      >
        {props.label}
      </FormLabel>
      <Input
        placeholder={props.placeholder}
        onChange={props.onChange}
        type={props.type}
        defaultValue={props.defaultValue}
        endDecorator={
          showSuccessIndicator ? (
            <CheckIcon sx={{ color: "#2e7d32", fontSize: "sm" }} />
          ) : (
            props.endDecorator
          )
        }
        fullWidth={true}
        onFocus={() => setIsFocused(true)}
        onBlur={() => setIsFocused(false)}
        sx={getInputStyles()}
      />
      {hasError ? (
        <FormHelperText sx={{ color: "#d32f2f", fontSize: "sm" }}>
          {errorMessage}
        </FormHelperText>
      ) : showSuccessIndicator ? (
        <FormHelperText sx={{ color: "#2e7d32", fontSize: "sm" }}>
          ✓ 入力内容が正しいです
        </FormHelperText>
      ) : null}
    </FormControl>
  );
}
