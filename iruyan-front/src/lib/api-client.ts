import axios, { AxiosError, AxiosResponse } from "axios";
import type {
  ApiResponse,
  ApiErrorResponse,
} from "@/types/api";

// Custom error class for API errors
export class ApiError extends Error {
  public status: number;
  public response?: ApiErrorResponse;

  constructor(message: string, status: number, response?: ApiErrorResponse) {
    super(message);
    this.name = "ApiError";
    this.status = status;
    this.response = response;
  }
}

const apiClient = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_BASE_URL || "http://localhost:8080/api/v1",
  headers: {
    "Content-Type": "application/json",
  },
  timeout: 10000,
});

// Request interceptor
apiClient.interceptors.request.use(
  (config) => {
    // Add auth token if available (only in browser environment)
    if (typeof window !== 'undefined') {
      const token = localStorage.getItem("authToken");
      if (token) {
        config.headers.Authorization = `Bearer ${token}`;
      }
    }
    return config;
  },
  (_error: unknown) => {
    console.log(_error);
    return Promise.reject(new ApiError("Request failed", 0, undefined));
  }
);

// Response interceptor
apiClient.interceptors.response.use(
  (response: AxiosResponse) => {
    return response;
  },
  (error: AxiosError<ApiErrorResponse>) => {
    const status = error.response?.status || 0;
    const errorData = error.response?.data;

    if (status === 401) {
      // Handle unauthorized access
      if (typeof window !== 'undefined') {
        localStorage.removeItem("authToken");
      }
      const unauthorizedError = new ApiError(
        "認証が必要です",
        401,
        errorData
      );
      return Promise.reject(unauthorizedError);
    }

    // Create a standardized error
    const apiError = new ApiError(
      errorData?.error || error.message || "API エラーが発生しました",
      status,
      errorData
    );

    return Promise.reject(apiError);
  }
);

// Type-safe API call wrapper
export const makeApiCall = async <T>(
  method: string,
  endpoint: string,
  data?: unknown
): Promise<ApiResponse<T>> => {
  try {
    const response = await apiClient.request({
      method,
      url: endpoint,
      data,
    });

    // Return standardized success response
    return {
      success: true,
      data: response.data,
      status: response.status,
    };
  } catch (error) {
    if (error instanceof ApiError) {
      return {
        success: false,
        error: error.message,
        status: error.status,
        details: error.response,
      };
    }

    // Fallback for unexpected errors
    return {
      success: false,
      error: "予期しないエラーが発生しました",
      status: 0,
    };
  }
};

// Helper functions for common API operations with improved typing
export const apiGet = async <T>(url: string): Promise<ApiResponse<T>> => {
  return makeApiCall<T>("GET", url);
};

export const apiPost = async <T>(url: string, data?: unknown): Promise<ApiResponse<T>> => {
  return makeApiCall<T>("POST", url, data);
};

export const apiPut = async <T>(url: string, data?: unknown): Promise<ApiResponse<T>> => {
  return makeApiCall<T>("PUT", url, data);
};

export const apiDelete = async <T>(url: string): Promise<ApiResponse<T>> => {
  return makeApiCall<T>("DELETE", url);
};

// Utility function to build URLs with parameters
export const buildApiUrl = (template: string, params: Record<string, string | number>): string => {
  let url = template;
  Object.entries(params).forEach(([key, value]) => {
    url = url.replace(`:${key}`, String(value));
  });
  return url;
};

export default apiClient;