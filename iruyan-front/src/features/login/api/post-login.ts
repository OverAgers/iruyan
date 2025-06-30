import axios from "axios";
import { useCallback, useState } from "react";
import { z } from "zod";

export const postLoginRequestSchema = z.object({
  iruyanId: z.string(),
  password: z.string(),
});

export const postLoginResponseSchema = z.object({
  message: z.string(),
  user: z.object({
    iruyanId: z.string(),
    userName: z.string(),
    email: z.string(),
  }),
});

export type PostLoginRequest = z.infer<typeof postLoginRequestSchema>;
export type PostLoginResponse = z.infer<typeof postLoginResponseSchema>;

export default function usePostLoginRequest() {
  const requestURL = `${process.env.NEXT_PUBLIC_API_URL}/login`;

  const [data, setData] = useState<PostLoginResponse | null>(null);
  const [error, setError] = useState<Error | null>(null);
  const [isMutating, setIsMutating] = useState(false);

  const login = useCallback(async (props: PostLoginRequest) => {
    setIsMutating(true);
    try {
      const res = await axios.post(
        requestURL,
        new URLSearchParams({
          iruyanId: props.iruyanId,
          password: props.password,
        }),
        {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
        }
      );
      const result = postLoginResponseSchema.parse(res.data);

      setData(result);
      return result;
    } catch (err) {
      setError(err as Error);
      throw err;
    } finally {
      setIsMutating(false);
    }
  }, [requestURL]);

  return { data, error, isMutating, login };
}
