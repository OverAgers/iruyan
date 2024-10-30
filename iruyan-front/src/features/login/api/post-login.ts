import axios from "axios";
import { useCallback } from "react";
import useSWRMutation from "swr/mutation";

import z from "zod";

export const postLoginSchema = z.object({
  username: z.string().min(1, "required"),
  password: z.string().min(1, "required"),
});

export type PostLoginRequest = z.infer<typeof postLoginSchema>;

export default function UseLogin() {
  // const requestURL = `${process.env.NEXT_PUBLIC_API_URL}/login`;
  const requestURL = `http://localhost:8080/login`;

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

  // trigger に loginData を直接渡す
  const login = (loginData: PostLoginRequest) => {
    trigger(loginData);
  };

  return { data, error, isLoading, login };
}
