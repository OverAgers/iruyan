import axios from "axios";
import { useCallback } from "react";
import useSWRMutation from "swr/mutation";

import z from "zod";

export const postLoginSchema = z.object({
  iruyanId: z.string().min(1, "入力してください"),
  password: z.string().min(1, "入力してください"),
});

export type PostLoginRequest = z.infer<typeof postLoginSchema>;

export default function UseLoginRequest() {
  // const requestURL = `${process.env.NEXT_PUBLIC_API_URL}/login`;
  const requestURL = `http://localhost:8080/Login`;

  const fetcher = useCallback(
    async (url: string, { arg }: { arg: PostLoginRequest }) => {
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

  const login = (LoginData: PostLoginRequest) => {
    trigger(LoginData);
  };

  return { data, error, isLoading, login };
}
