import axios from "axios";
import { useCallback } from "react";
import z from "zod";
import useSWR from "swr";

export const getFocusRankingSchema = z.object({
  message: z.string(),
});

export type GetWeeklyWorkRequest = z.infer<typeof getFocusRankingSchema >;

export default function useGetFocusRanking(user_id: string, room_id: string) {
  const requestURL = `http://localhost:8080/user/${user_id}/ranking`;

  const fetcher = useCallback(
    async (url: string) => {
      try {
        const res = await axios.get(url, { params: { room_id: room_id} });
        const data = getFocusRankingSchema.parse(res.data);
        console.log("data", data);
        return data;
      } catch (error: any) {
        throw new Error(
          `集中時間ランキング情報取得エラー: ${error.response?.data?.message || error.message}`
        );
      }
    },
    [room_id]
  );

  const { data, error, isLoading } = useSWR(requestURL, fetcher);

  return { data, error, isLoading };
}
