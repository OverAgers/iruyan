import axios from 'axios';
import { useCallback, useState } from 'react';
import useSWRMutation from 'swr';

import z from 'zod';

export const postLoginSchema = z.object({
  userID: z.string().min(1,'required'),
  password: z.string().min(1,'required'),
});

export type PostLoginRequest = z.infer<typeof postLoginSchema>;

export default function UseLogin() {
  const [request, setRequest] = useState<PostLoginRequest>();
  const requestURL = process.env.NEXT_PUBLIC_API_BASE_URL +"login";
  const fetcher = useCallback(() => {
    return axios
      .post(requestURL, request)
      .then(async (res) => {
        return res.data;
      })
      .catch((error) => {
        throw new Error(`CardCheckout Error: ${error.message}`);
      });
  }, [request, requestURL]);

  const { data, error,isLoading } = useSWRMutation(requestURL, fetcher);
  return { data, error, isLoading, setRequest };
}