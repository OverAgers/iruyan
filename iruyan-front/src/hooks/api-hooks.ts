import { useCallback, useState } from "react";
import useSWR, { SWRConfiguration } from "swr";
import { z } from "zod";
import apiClient, { ApiError } from "@/lib/api-client";
import type { ApiResponse } from "@/types/api";

// Hook options interface
interface UseAPIOptions<T> extends SWRConfiguration {
  method?: "GET" | "POST" | "PUT" | "DELETE";
  schema?: z.ZodSchema<T>;
  enabled?: boolean; // Allow conditional fetching
}

// Enhanced error type for API hooks
export interface ApiHookError {
  message: string;
  status?: number;
  isValidationError: boolean;
  originalError?: unknown;
}

// Main API hook
export function useAPI<T>(
  url: string | null,
  options: UseAPIOptions<T> = {}
) {
  const {
    method = "GET",
    schema,
    enabled = true,
    revalidateOnFocus = false,
    ...swrOptions
  } = options;

  const fetcher = useCallback(
    async (fetchUrl: string): Promise<T> => {
      try {
        const response = await apiClient.request<ApiResponse<T>>({
          url: fetchUrl,
          method,
        });

        // Handle API response format
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        const data = (response.data as any)?.data || response.data;

        // Validate with schema if provided
        if (schema) {
          return schema.parse(data);
        }

        return data as T;
      } catch (error) {
        if (error instanceof z.ZodError) {
          const hookError: ApiHookError = {
            message: "レスポンスの形式が正しくありません",
            isValidationError: true,
            originalError: error,
          };
          throw hookError;
        }

        if (error instanceof ApiError) {
          const hookError: ApiHookError = {
            message: error.message,
            status: error.status,
            isValidationError: false,
            originalError: error,
          };
          throw hookError;
        }

        const hookError: ApiHookError = {
          message: error instanceof Error ? error.message : "予期しないエラーが発生しました",
          isValidationError: false,
          originalError: error,
        };
        throw hookError;
      }
    },
    [method, schema]
  );

  const shouldFetch = enabled && url !== null;
  const swrKey = shouldFetch ? url : null;

  const { data, error, isLoading, mutate, isValidating } = useSWR(
    swrKey,
    fetcher,
    {
      revalidateOnFocus,
      ...swrOptions,
    }
  );

  return {
    data,
    error: error as ApiHookError | undefined,
    isLoading,
    isValidating,
    mutate,
    refetch: () => mutate(),
  };
}

// Specialized hooks for common patterns
export function useAuthenticatedAPI<T>(
  url: string | null,
  options: UseAPIOptions<T> = {}
) {
  const isAuthenticated = typeof window !== "undefined" && !!localStorage.getItem("authToken");

  return useAPI(url, {
    ...options,
    enabled: options.enabled !== false && isAuthenticated,
  });
}

// Hook for conditional API calls
export function useConditionalAPI<T>(
  url: string | null,
  condition: boolean,
  options: UseAPIOptions<T> = {}
) {
  return useAPI(url, {
    ...options,
    enabled: options.enabled !== false && condition,
  });
}

// Hook with automatic retry on error
export function useRetryAPI<T>(
  url: string | null,
  options: UseAPIOptions<T> = {}
) {
  return useAPI(url, {
    ...options,
    errorRetryCount: 3,
    errorRetryInterval: 1000,
  });
}

// Mutation hook for POST/PUT/DELETE operations
export function useMutation<TData, TVariables = unknown>(
  mutationFn: (variables: TVariables) => Promise<TData>
) {
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<ApiHookError | null>(null);
  const [data, setData] = useState<TData | null>(null);

  const mutate = useCallback(
    async (variables: TVariables) => {
      try {
        setIsLoading(true);
        setError(null);
        const result = await mutationFn(variables);
        setData(result);
        return result;
      } catch (err) {
        const hookError: ApiHookError = {
          message: err instanceof Error ? err.message : "Mutation failed",
          isValidationError: false,
          originalError: err,
        };
        setError(hookError);
        throw hookError;
      } finally {
        setIsLoading(false);
      }
    },
    [mutationFn]
  );

  return {
    mutate,
    data,
    error,
    isLoading,
    reset: () => {
      setData(null);
      setError(null);
      setIsLoading(false);
    },
  };
}

export default useAPI;