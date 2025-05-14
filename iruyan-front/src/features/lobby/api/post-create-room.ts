import axios from 'axios';
import { useCallback } from 'react';
import useSWRMutation from 'swr/mutation';
import { z } from 'zod';

export const postCreateRoomRequestSchema = z.object({
  userName: z.string(),
});

export const postCreateRoomResponseSchema = z.object({
  message: z.string(),
  roomId: z.string().uuid(),
  roomName: z.string(),
});

export type PostCreateRoomRequest = z.infer<typeof postCreateRoomRequestSchema>;
export type PostCreateRoomResponse = z.infer<typeof postCreateRoomResponseSchema>;

export default function usePostCreateRoomRequest() {
  const baseURL = process.env.NEXT_PUBLIC_API_BASE_URL || '';
  const requestURL = baseURL + '/rooms';

  const fetcher = useCallback(
    (url: string, { arg }: { arg: PostCreateRoomRequest }) => {
      return axios
        .post(url, arg, {
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
          },
        })
        .then(async (res) => {
          const result = res.data;
          return postCreateRoomRequestSchema.parse(result);
        })
        .catch((error) => {
          throw error;
        });
    },
    [requestURL]
  );

  const { data, error, isMutating, trigger } = useSWRMutation<PostCreateRoomRequest, Error>(requestURL, fetcher);
  return { data, error, isMutating, trigger };
}
