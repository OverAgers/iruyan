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

export default function UseGetRoomList() {
  const requestURL = `http://localhost:8080/rooms`;

  const fetcher = useCallback(
    async (url: string) => {
      try {
        const res = await axios.get(url);
        const data = getRoomListSchema.parse(res.data);
        console.log("data", data);
        return data;
      } catch (error: any) {
        throw new Error(
          `ルーム一覧取得エラー: ${error.response?.data?.message || error.message}`
        );
      }
    },
    []
  );

  const { data, error, isLoading } = useSWR(requestURL, fetcher);

  return { data, error, isLoading };
}
