import axios from "axios";
import { useCallback } from "react";
import useSWRMutation from "swr/mutation";
import { z } from "zod";
import { UserInfo } from "@/types/user-info";

export const postLogoutSchema = z.object({
  iruyanID: z.string(),
});

export type PostLogoutRequest = z.infer<typeof postLogoutSchema>;

export default function usePostLogoutRequest() {
  const BASE_URL =
    process.env.NEXT_PUBLIC_API_BASE_URL || "http://localhost:8080";
  const requestURL = `${BASE_URL}/logout`;

  const fetcher = useCallback(
    async (url: string, { arg }: { arg: PostLogoutRequest }) => {
      try {
        const validatedData = postLogoutSchema.parse(arg);
        const response = await axios.post(url, validatedData);
        return response.data;
      } catch (error: any) {
        if (axios.isAxiosError(error)) {
          throw new Error(
            `入室エラー: ${error.response?.data?.message || error.message}`
          );
        } else if (error instanceof z.ZodError) {
          throw new Error(
            `入力エラー: ${error.errors.map((e) => e.message).join(", ")}`
          );
        } else {
          throw new Error(`未知のエラーが発生しました: ${error.message}`);
        }
      }
    },
    [requestURL]
  );

  const {
    data,
    error,
    isMutating: isLoading,
    trigger,
  } = useSWRMutation<UserInfo, Error, string, PostLogoutRequest>(
    requestURL,
    fetcher
  );

  const logout = async (
    requestData: PostLogoutRequest
  ): Promise<UserInfo> => {
    return await trigger(requestData);
  };

  return { data, error, isLoading, logout };
}
