import axios from "axios";
import { useCallback } from "react";
import useSWRMutation from "swr/mutation";
import { z } from "zod";
import { UserInfo } from "@/types/user-info";

export const postRoomEnterSchema = z.object({
  user_id: z.string(),
  room_id: z.string(),
});

export type PostRoomEnterRequest = z.infer<typeof postRoomEnterSchema>;

export default function usePostRoomEnterRequest() {
  const BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL || "http://localhost:8080";
  // const requestURL = `${BASE_URL}/${roomId}/enter`;

  const fetcher = useCallback(
    async (url: string, { arg }: { arg: PostRoomEnterRequest }) => {
      try {
        const idurl = `${BASE_URL}/${arg.room_id}/enter/${arg.user_id}`;
        // const validatedData = postRoomEnterSchema.parse(arg);
        const response = await axios.post(idurl, arg.user_id);
        console.log("入室成功:", response.data);
        return response.data;
      } catch (error: any) {
        if (axios.isAxiosError(error)) {
          // throw new Error(
          //   `入室エラー: ${error.response?.data?.message || error.message}`
          // );
        } else if (error instanceof z.ZodError) {
          throw new Error(
            `入力エラー: ${error.errors.map((e) => e.message).join(", ")}`
          );
        } else {
          throw new Error(`未知のエラーが発生しました: ${error.message}`);
        }
      }
    },
    []
  );

  const {
    data,
    error,
    isMutating: isLoading,
    trigger,
  } = useSWRMutation<UserInfo, Error, string, PostRoomEnterRequest>(
    BASE_URL,
    fetcher
  );

  const entry = async (requestData: PostRoomEnterRequest): Promise<UserInfo> => {
    return await trigger(requestData);
  };

  return { data, error, isLoading, entry };
}
