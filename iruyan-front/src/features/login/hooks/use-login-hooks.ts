import { useForm } from "react-hook-form";

import { zodResolver } from "@hookform/resolvers/zod";

import { LoginForm, loginFormSchema } from "@/schema/login-form-schema";

type Props = {
  onSubmit: (data: LoginForm) => void;
};

export default function UseLoginForm({ onSubmit }: Props) {
  const {
    register,
    handleSubmit,
    setValue,
    watch,
    formState: { errors },
  } = useForm<LoginForm>({
    resolver: zodResolver(loginFormSchema),
    mode: "onBlur",
  });
  return {
    register,
    handleSubmit,
    setValue,
    watch,
    errors,
    onSubmit: handleSubmit(onSubmit),
  };
}
