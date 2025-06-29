import axios from "axios";
import { useCallback } from "react";
import z from "zod";
import useSWR from "swr";


export const getWeeklyWorkSchema = z.object({
  message: z.string(),
  workInfo: z.array(
    z.object({
      date: z.string(),
      duration: z.string(),
      rest: z.string(),
      task: z.string(),
    })
  ),
});

export type GetWeeklyWorkRequest = z.infer<typeof getWeeklyWorkSchema>;

// カスタムフック
export default function useGetWeeklyWork(userId: string, roomId: string) {
  const requestURL = `http://localhost:8080/user/${userId}/work_info`;

  const fetcher = useCallback(
    async (url: string) => {
      try {
        const res = await axios.get(url, { params: { roomId: roomId} });
        const data = getWeeklyWorkSchema.parse(res.data);
        console.log("data", data);
        return data;
      } catch (error: unknown) {
        const errorMessage = error instanceof Error 
          ? error.message 
          : 'Unknown error occurred';
        throw new Error(
          `作業情報取得エラー: ${errorMessage}`
        );
      }
    },
    [roomId]
  );

  const { data, error, isLoading } = useSWR(requestURL, fetcher);

  return { data, error, isLoading };
}
