import axios from "axios";
import { useCallback } from "react";
import useSWRMutation from "swr/mutation";
import { z } from "zod";
import { UserInfo } from "@/types/user-info";

export const postRoomEnterSchema = z.object({
  iruyanID: z.string(),
  task: z.string(),
});

export type PostRoomEnterRequest = z.infer<typeof postRoomEnterSchema>;



export default function usePostRoomEnterRequest(roomId: string) {
  const BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL || "http://localhost:8080";
  const requestURL = `${BASE_URL}/${roomId}/enter`;

  const fetcher = useCallback(
    async (url: string, { arg }: { arg: PostRoomEnterRequest }) => {
      try {
        const validatedData = postRoomEnterSchema.parse(arg);
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
  } = useSWRMutation<UserInfo, Error, string, PostRoomEnterRequest>(
    requestURL,
    fetcher
  );

  const entry = async (requestData: PostRoomEnterRequest): Promise<UserInfo> => {
    if (!roomId) {
      throw new Error("部屋IDが指定されていません。");
    }
    return await trigger(requestData);
  };

  return { data, error, isLoading, entry };
}
