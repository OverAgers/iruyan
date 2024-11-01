import axios from "axios";
import { useCallback } from "react";
import z from "zod";
import useSWR from "swr";


export const getWeeklyWorkSchema = z.object({
  message: z.string(),
  work_info: z.array(
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
export default function useGetWeeklyWork(user_id: string, room_id: string) {
  const requestURL = `http://localhost:8080/user/${user_id}/work_info`;

  const fetcher = useCallback(
    async (url: string) => {
      try {
        const res = await axios.get(url, { params: { room_id: room_id} });
        const data = getWeeklyWorkSchema.parse(res.data);
        console.log("data", data);
        return data;
      } catch (error: any) {
        throw new Error(
          `作業情報取得エラー: ${error.response?.data?.message || error.message}`
        );
      }
    },
    [room_id]
  );

  const { data, error, isLoading } = useSWR(requestURL, fetcher);

  return { data, error, isLoading };
}
