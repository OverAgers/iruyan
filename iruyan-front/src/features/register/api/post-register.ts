import axios from "axios";
import { useCallback, useState } from "react";
import useSWRMutation from "swr/mutation";

import z from "zod";

export const postRegisterSchema = z.object({
  userID: z.string().min(1, "required"),
  userName: z.string().min(1, "required"),
  password: z.string().min(1, "required"),
  email: z.string().min(1, "required"),
});

export type PostRegisterRequest = z.infer<typeof postRegisterSchema>;

export default function UseRegister() {
  // const requestURL = `${process.env.NEXT_PUBLIC_API_URL}/login`;
  const requestURL = `http://localhost:8080/register`;

  const fetcher = useCallback(
    async (url: string, { arg }: { arg: PostRegisterRequest }) => {
      try {
        const res = await axios.post(url, arg);
        return res.data;
      } catch (error: any) {
        throw new Error(`ログインエラー: ${error.message}`);
      }
    },
    []
  );

  const {
    data,
    error,
    isMutating: isLoading,
    trigger,
  } = useSWRMutation(requestURL, fetcher);

  const register = (registerData: PostRegisterRequest) => {
    trigger(registerData);
  };

  return { data, error, isLoading, register };
}