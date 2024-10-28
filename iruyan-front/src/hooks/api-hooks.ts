import { useCallback } from "react";
import useSWR from "swr";
import apiClient from "@/lib/api-client";
import z from "zod";

type UseAPIOptions<T> = {
  method?: "GET" | "POST" | "PUT" | "DELETE";
  schema?: z.ZodSchema<T>;
  revalidateOnFocus?: boolean;
}

export default function useAPI<T>(url: string, options: UseAPIOptions<T> = {}) {
  const { method = 'GET', schema, revalidateOnFocus = false } = options;

  const fetcher = useCallback(
    async () => {
      try {
        const response = await apiClient.request<T>({
          url,
          method,
        });

        if (schema) {
          return schema.parse(response.data);
        }
        return response.data;
      } catch (error) {
        if (error instanceof z.ZodError) {
          throw new Error('Response validation error');
        }
        throw new Error("API Error: ${error.message}");
      }
    }, [url, method, schema]
  );

  const { data, error, isLoading } = useSWR(url, fetcher, { revalidateOnFocus });
  return { data, error, isLoading };
}