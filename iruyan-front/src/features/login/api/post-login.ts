import axios from "axios";
import { useCallback, useState } from "react";
import { z } from "zod";
import { PostLoginRequest, postLoginRequestSchema, postLoginResponseSchema } from "@/schema/login-form-schema";

export type PostLoginResponse = z.infer<typeof postLoginResponseSchema>;
export type { PostLoginRequest } from "@/schema/login-form-schema";

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
      return result; // ← ★ ここを追加
    } catch (err) {
      setError(err as Error);
      throw err; // ← ★ catch しても呼び出し元にエラーを伝える
    } finally {
      setIsMutating(false);
    }
  }, [requestURL]);

  return { data, error, isMutating, login };
}
