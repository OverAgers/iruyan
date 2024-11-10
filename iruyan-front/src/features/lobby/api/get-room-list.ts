import axios from 'axios';
import { useCallback } from 'react';
import z from 'zod';
import useSWR from 'swr';

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
  const baseURL = process.env.NEXT_PUBLIC_API_BASE_URL || '';
  const requestURL = baseURL + '/rooms';

  const fetcher = useCallback(() => {
    return axios
      .get(requestURL, {
        headers: {
          'Content-Type': 'Application/json',
        },
      })
      .then(async (res) => {
        const result = res.data;
        return getRoomListSchema.parse(result);
      })
      .catch((error) => {
        throw error;
      });
  }, [requestURL]);

  const { data, error, isLoading, mutate } = useSWR<GetRoomListRequest, Error>(
    requestURL,
    fetcher
  );

  return { data, error, isLoading, mutate };
}
