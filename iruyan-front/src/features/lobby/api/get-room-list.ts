import axios from "axios";
import { useCallback } from "react";
import z from "zod";
import useSWR from "swr";

export const getRoomListSchema = z.object({
  message: z.string(),
  rooms: z.array(
    z.object({
      roomId: z.string().uuid(),
      roomName: z.string(),
      seats: z.array(
        z.object({
          seatId: z.string().uuid(),
          roomId: z.string().uuid(),
          seatNumber: z.number(),
        })
      ),
    })
  ),
});

export type GetRoomListRequest = z.infer<typeof getRoomListSchema>;

export default function useGetRoomList() {
  const BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL || "http://localhost:8080";
  const requestURL = `${BASE_URL}/rooms`;

  const fetcher = useCallback(
    async (url: string): Promise<GetRoomListRequest> => {
      console.log("データ取得中:", url);
      try {
        const res = await axios.get(url);
        console.log("データ取得成功:", res.data);
        return res.data;
      } catch (error: any) {
        if (axios.isAxiosError(error)) {
          throw new Error(
            `ルーム一覧取得エラー: ${
              error.response?.data?.message || error.message
            }`
          );
        } else if (error instanceof z.ZodError) {
          throw new Error(
            `データ検証エラー: ${error.errors.map((e) => e.message).join(", ")}`
          );
        } else {
          throw new Error(`未知のエラーが発生しました: ${error.message}`);
        }
      }
    },
    []
  );

  const { data, error, isLoading } = useSWR<GetRoomListRequest, Error>(
    requestURL,
    fetcher,);

  return { data, error, isLoading };
}
