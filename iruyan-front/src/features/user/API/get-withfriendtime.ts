import axios, { AxiosError } from "axios";
import { useCallback } from "react";
import z from "zod";
import useSWR from "swr";

export const getwithFriendTimeSchema = z.object({
  message: z.string(),
});

export type GetWeeklyWorkRequest = z.infer<typeof getwithFriendTimeSchema >;

export default function useGetWithFriendTime(userId: string, roomId: string) {
  const requestURL = `http://localhost:8080/user/${userId}/together`;

  const fetcher = useCallback(
    async (url: string) => {
      try {
        const res = await axios.get(url, { params: { roomId: roomId} });
        const data = getwithFriendTimeSchema .parse(res.data);
        console.log("data", data);
        return data;
      } catch (error) {
        const err = error as AxiosError<{ message?: string }>;
        throw new Error(
          `友達と過ごした時間情報取得エラー: ${err.response?.data?.message || err.message}`
        );
      }
    },
    [roomId]
  );

  const { data, error, isLoading } = useSWR(requestURL, fetcher);

  return { data, error, isLoading };
}
