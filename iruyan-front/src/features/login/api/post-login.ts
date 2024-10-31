import axios from "axios";
import { useCallback } from "react";
import useSWRMutation from "swr/mutation";
import z from "zod";
import { UserInfo } from "@/types/user-info";

export const postLoginSchema = z.object({
  iruyanId: z.string().min(1, "入力してください"),
  password: z.string().min(1, "入力してください"),
});

export type PostLoginRequest = z.infer<typeof postLoginSchema>;

export default function UseLoginRequest() {
  const requestURL = `http://localhost:8080/Login`;

  const fetcher = useCallback(
    async (url: string, { arg }: { arg: PostLoginRequest }) => {
      try {
        const res = await axios.post(url, arg);
        return res.data as UserInfo;
      } catch (error: any) {
        throw new Error(
          `ログインエラー: ${error.response?.data?.message || error.message}`
        );
      }
    },
    []
  );

  const {
    data,
    error,
    isMutating: isLoading,
    trigger,
  } = useSWRMutation<UserInfo, any, string, PostLoginRequest>(
    requestURL,
    fetcher
  );

  const login = async (loginData: PostLoginRequest): Promise<UserInfo> => {
    const userData = await trigger(loginData);
    return userData;
  };

  return { data, error, isLoading, login };
}
