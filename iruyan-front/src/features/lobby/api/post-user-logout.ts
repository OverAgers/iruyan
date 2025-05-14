import axios from 'axios';
import { useCallback } from 'react';
import useSWRMutation from 'swr/mutation';
import { z } from 'zod';

export const postLogoutRequestSchema = z.object({
  iruyanId: z.string(),
});

export type PostLogoutRequest = z.infer<typeof postLogoutRequestSchema>;

export default function usePostLogoutRequest() {
  const baseURL = process.env.NEXT_PUBLIC_API_BASE_URL || '';
  const requestURL = baseURL + `/logout`;

  const fetcher = useCallback(
    (url: string, { arg }: { arg: PostLogoutRequest }) => {
      return axios
        .post(url, arg, {
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
          },
        })
        .then(async (res) => {
          const result = res.data;
          return result;
        })
        .catch((error) => {
          throw error;
        });
    },
    [requestURL]
  );

  const { data, error, isMutating, trigger } = useSWRMutation(requestURL, fetcher);

  return { data, error, isMutating, trigger };
}
