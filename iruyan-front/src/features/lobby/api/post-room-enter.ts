import axios from 'axios';
import { useCallback } from 'react';
import useSWRMutation from 'swr/mutation';
import { z } from 'zod';

export const postRoomEnterSchema = z.object({
  iruyanId: z.string().uuid(),
  task: z.string(),
});

export type PostRoomEnterRequest = z.infer<typeof postRoomEnterSchema>;

export default function usePostRoomEnterRequest(
  Props: PostRoomEnterRequest,
  roomId: string
) {
  const baseURL = process.env.NEXT_PUBLIC_API_BASE_URL || '';
  const requestURL = baseURL + `/rooms/${roomId}/enter`;

  const fetcher = useCallback(() => {
    return axios
      .post(requestURL, Props, {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
      })
      .then(async (res) => {
        const result = res.data;
        return postRoomEnterSchema.parse(result);
      })
      .catch((error) => {
        throw error;
      });
  }, [requestURL, Props]);

  const { data, error, isMutating } = useSWRMutation(requestURL, fetcher);

  return { data, error, isMutating };
}
