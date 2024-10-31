import { useForm } from "react-hook-form";

import { zodResolver } from "@hookform/resolvers/zod";

import {
  RegisterForm,
  registerFormSchema,
} from "@/schema/register-form-schema";

type Props = {
  onSubmit: (data: RegisterForm) => void;
};

export default function UseRegisterForm({ onSubmit }: Props) {
  const {
    register,
    handleSubmit,
    setValue,
    watch,
    formState: { errors },
  } = useForm<RegisterForm>({
    resolver: zodResolver(registerFormSchema),
    mode: "onBlur",
  });
  return {
    register,
    handleSubmit,
    setValue,
    watch,
    errors,
    handleFormSubmit: handleSubmit(onSubmit),
  };
}
