import axios from 'axios';
import { useCallback } from 'react';
import z from 'zod';
import useSWR from 'swr';

export const getFocusRankingSchema = z.object({
  message: z.string(),
});

export type GetWeeklyWorkRequest = z.infer<typeof getFocusRankingSchema>;

export default function useGetFocusRanking(userId: string, roomId: string) {
  const requestURL = `http://localhost:8080/user/${userId}/ranking`;

  const fetcher = useCallback(
    async (url: string) => {
      try {
        const res = await axios.get(url, { params: { roomId: roomId } });
        const data = getFocusRankingSchema.parse(res.data);
        console.log('data', data);
        return data;
      } catch (error: any) {
        throw new Error(`集中時間ランキング情報取得エラー: ${error.response?.data?.message || error.message}`);
      }
    },
    [roomId]
  );

  const { data, error, isLoading } = useSWR(requestURL, fetcher);

  return { data, error, isLoading };
}
