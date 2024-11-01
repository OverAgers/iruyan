import axios from "axios";
import { useCallback } from "react";
import useSWRMutation from "swr/mutation";
import { z } from "zod";
import { UserInfo } from "@/types/user-info";

// スキーマ定義
export const postRoomSchema = z.object({
  name: z.string(),
});

export type PostRoomRequest = z.infer<typeof postRoomSchema>;

export default function usePostRoomRequest() {
  const BASE_URL =
    process.env.NEXT_PUBLIC_API_BASE_URL || "http://localhost:8080";
  const requestURL = `${BASE_URL}/rooms`;

  const fetcher = useCallback(
    async (url: string, { arg }: { arg: PostRoomRequest }) => {
      const validatedData = postRoomSchema.parse(arg);
      console.log("部屋作成中:", validatedData);
      try {
        const response = await axios.post(url, validatedData, {
          headers: { "Content-Type": "application/x-www-form-urlencoded" },
        });
        console.log("部屋作成成功:", response.data);
        return response.data;
      } catch (error: any) {
        if (axios.isAxiosError(error)) {
          throw new Error(
            `作成エラー: ${
              error.response?.data?.message || "不明なエラーが発生しました"
            }`
          );
        } else {
          // その他のエラー
          throw new Error(`未知のエラーが発生しました: ${error.message}`);
        }
      }
    },
    []
  );

  const { data, error, isMutating, trigger } = useSWRMutation<
    UserInfo,
    Error,
    string,
    PostRoomRequest
  >(requestURL, fetcher);

  const createRoom = async (
    requestData: PostRoomRequest
  ): Promise<UserInfo> => {
    try {
      const result = await trigger(requestData);
      return result;
    } catch (err) {
      throw err;
    }
  };

  return { data, error, isLoading: isMutating, createRoom };
}
