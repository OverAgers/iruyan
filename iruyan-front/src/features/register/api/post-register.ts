import axios from 'axios';
import { useCallback } from 'react';
import useSWRMutation from 'swr/mutation';

import z from 'zod';

export const postRegisterRequestSchema = z.object({
  iruyanId: z.string(),
  userName: z.string(),
  email: z.string(),
  password: z.string(),
});

export const postRegisterResponseSchema = z.object({
  message: z.string(),
  user: z.object({
    iruyanId: z.string().uuid(),
    userName: z.string(),
    email: z.string().email(),
  }),
});

export type PostRegisterRequest = z.infer<typeof postRegisterRequestSchema>;
export type PostRegisterResponse = z.infer<typeof postRegisterResponseSchema>;

export default function UsePostRegisterRequest(Props: PostRegisterRequest) {
  const baseURL = process.env.NEXT_PUBLIC_API_URL;
  const requestURL = baseURL + '/register';

  const fetcher = useCallback(() => {
    return axios
      .post(requestURL, Props, {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
      })
      .then(async (res) => {
        const result = res.data;
        return postRegisterResponseSchema.parse(result);
      })
      .catch((error) => {
        throw error;
      });
  }, [Props, requestURL]);

  const { data, error, isMutating } = useSWRMutation(requestURL, fetcher);

  return { data, error, isMutating };
}
