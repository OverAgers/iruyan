// src/features/register/api/post-register.ts
import axios from "axios";
import { useState, useCallback } from "react";
import { postRegisterResponseSchema, PostRegisterRequest, PostRegisterResponse } from "@/schema/register-form-schema";

export default function UsePostRegisterRequest() {
  const [data, setData] = useState<PostRegisterResponse | null>(null);
  const [error, setError] = useState<Error | null>(null);
  const [isLoading, setIsLoading] = useState(false);

  const register = useCallback(async (props: PostRegisterRequest) => {
    setIsLoading(true);
    try {
      const res = await axios.post(
        `${process.env.NEXT_PUBLIC_API_URL}/register`,
        new URLSearchParams({
          iruyanId: props.iruyanId,
          userName: props.name,
          email: props.email,
          password: props.password,
        }),
        {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
        }
      );
      const result = postRegisterResponseSchema.parse(res.data);
      setData(result);
      return result;
    } catch (err) {
      setError(err as Error);
      throw err;
    } finally {
      setIsLoading(false);
    }
  }, []);

  return { data, error, isLoading, register };
}
