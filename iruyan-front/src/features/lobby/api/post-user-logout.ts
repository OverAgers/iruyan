import axios from 'axios';
import { useCallback } from 'react';
import useSWR from 'swr';
import { z } from 'zod';

export const postLogoutRequestSchema = z.object({
  iruyanId: z.string(),
});

export type PostLogoutRequest = z.infer<typeof postLogoutRequestSchema>;

export default function usePostLogoutRequest(Props: PostLogoutRequest) {
  const baseURL = process.env.NEXT_PUBLIC_API_BASE_URL || '';
  const requestURL = baseURL + `/logout`;

  const fetcher = useCallback(() => {
    return axios
      .post(requestURL, Props, {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
      })
      .then(async (res) => {
        const result = res.data;
        return postLogoutRequestSchema.parse(result);
      })
      .catch((error) => {
        throw error;
      });
  }, [requestURL, Props]);

  const { data, error, isLoading } = useSWR(requestURL, fetcher);

  return { data, error, isLoading };
}
