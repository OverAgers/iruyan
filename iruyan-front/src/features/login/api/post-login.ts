import axios from 'axios';
import { useCallback } from 'react';
import useSWRMutation from 'swr/mutation';
import z from 'zod';

export const postLoginRequestSchema = z.object({
  iruyanId: z.string(),
  password: z.string(),
});

export const postLoginResponseSchema = z.object({
  message: z.string(),
  user: z.object({
    iruyanId: z.string(),
    userName: z.string(),
    email: z.string().email(),
  }),
});

export type PostLoginRequest = z.infer<typeof postLoginRequestSchema>;
export type PostLoginResponse = z.infer<typeof postLoginResponseSchema>;

export default function usePostLoginRequest() {
  const baseURL = process.env.NEXT_PUBLIC_API_BASE_URL || '';
  const requestURL = baseURL + '/login';

  const fetcher = useCallback((url: string, { arg }: { arg: PostLoginRequest }) => {
    return axios
      .post(url, arg, {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
      })
      .then(async (res) => {
        const result = res.data;
        return postLoginResponseSchema.parse(result);
      })
      .catch((error) => {
        throw error;
      });
  }, []);

  const { data, error, isMutating, trigger } = useSWRMutation(requestURL, fetcher);

  return { data, error, isMutating, trigger };
}
