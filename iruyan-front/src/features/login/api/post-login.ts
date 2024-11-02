import axios from "axios";
import { useCallback } from "react";
import useSWRMutation from "swr/mutation";
import z from "zod";
import { UserInfo } from "@/types/user-info";

export const postLoginSchema = z.object({
  iruyanID: z.string().min(1, "入力してください"),
  password: z.string().min(1, "入力してください"),
});

export type PostLoginRequest = z.infer<typeof postLoginSchema>;

export default function UseLoginRequest() {
  const requestURL = `http://localhost:8080/login`;

  const fetcher = useCallback(
    async (url: string, { arg }: { arg: PostLoginRequest }) => {
      try {
        const request = postLoginSchema.parse(arg);
        const res = await axios.post(url, request, {
          headers: { "Content-Type": "application/x-www-form-urlencoded" }
        });
        const data = res.data.user;
        const userInfo: UserInfo = {
          iruyanID: data.iruyanID,
          name: data.name,
          email: data.email,
          status: "idle",
          workTime: 0,
          restTime: 0,
          startTime: 0
        }
        return userInfo;
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
