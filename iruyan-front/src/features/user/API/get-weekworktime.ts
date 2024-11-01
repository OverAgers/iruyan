import axios from "axios";
import { useCallback } from "react";
import z from "zod";
import useSWR from "swr";

// レスポンススキーマ
export const getWeeklyWorkSchema = z.object({
  message: z.string(),
  work_info: z.array(
    z.object({
      date: z.string(), // 日付（ISO 8601などの文字列として受け取ります）
      duration: z.string(), // 作業時間 (例: "HH:mm:ss" 形式の文字列)
      rest: z.string(), // 休憩時間 (例: "HH:mm:ss" 形式の文字列)
      task: z.string(), // タスクの内容
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
        // クエリパラメータとして roomId を付加
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
