import { UserInfo } from "@/types/user-info";
import axios from "axios";
import { useCallback } from "react";
import useSWRMutation from "swr/mutation";

import z from "zod";

export const postRegisterSchema = z.object({
  iruyanId: z.string(),
  userName: z.string(),
  email: z.string(),
  password: z.string(),
});

export type PostRegisterRequest = z.infer<typeof postRegisterSchema>;

export default function UsePostRegisterRequest() {
  // const requestURL = `${process.env.NEXT_PUBLIC_API_URL}/login`;
  const requestURL = `http://localhost:8080/register`;

  const fetcher = useCallback(
    async (url: string, { arg }: { arg: PostRegisterRequest }) => {
      try {
        console.log("fetcher", arg);
        const res = await axios.post(url, arg);
        return res.data as UserInfo;
      } catch (error: any) {
        throw new Error(
          error.response?.data?.message || `登録エラー: ${error.message}`
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
  } = useSWRMutation<UserInfo, any, string, PostRegisterRequest>(requestURL, fetcher);

  const register = (registerData: PostRegisterRequest) => {
    return trigger(registerData);
  };

  return { data, error, isLoading, register };
}
