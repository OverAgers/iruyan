import { UserInfo } from "@/types/user-info";
import axios from "axios";
import { useCallback } from "react";
import useSWRMutation from "swr/mutation";

import z from "zod";

export const postRegisterSchema = z.object({
  iruyanId: z.string(),
  name: z.string(),
  email: z.string(),
  password: z.string(),
});

export type PostRegisterRequest = z.infer<typeof postRegisterSchema>;

export default function UsePostRegisterRequest() {
  // const requestURL = `${process.env.NEXT_PUBLIC_API_URL}/login`;
  const requestURL = `http://localhost:8080/register`;

  const fetcher = useCallback(
    async (url: string, { arg }: { arg: PostRegisterRequest }) => {
      const request = postRegisterSchema.parse(arg);
      try {
        console.log("request", request);
        const res = await axios.post(url, request, {
          headers: { "Content-Type": "application/x-www-form-urlencoded" },
        });
        console.log("res", res.data);
        const data = res.data.user;
        const userInfo: UserInfo = {
          iruyanId: data.iruyanId,
          name: data.name,
          email: data.email,
          status: "idle",
          workTime: 0,
          restTime: 0,
          startTime: 0,
        };
        return userInfo;
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
