import axios, { AxiosError, AxiosResponse } from "axios";
import type { ApiResponse } from "@/types";

const apiClient = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_BASE_URL || "http://localhost:8080",
  headers: {
    "Content-Type": "application/json",
  },
  timeout: 10000,
});

// Request interceptor
apiClient.interceptors.request.use(
  (config) => {
    // Add auth token if available
    const token = localStorage.getItem("authToken");
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Response interceptor
apiClient.interceptors.response.use(
  (response: AxiosResponse) => {
    return response;
  },
  (error: AxiosError) => {
    if (error.response?.status === 401) {
      // Handle unauthorized access
      localStorage.removeItem("authToken");
      window.location.href = "/login";
    }
    return Promise.reject(error);
  }
);

// Helper functions for common API operations
export const apiGet = async <T>(url: string): Promise<ApiResponse<T>> => {
  try {
    const response = await apiClient.get<ApiResponse<T>>(url);
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const apiPost = async <T>(url: string, data?: any): Promise<ApiResponse<T>> => {
  try {
    const response = await apiClient.post<ApiResponse<T>>(url, data);
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const apiPut = async <T>(url: string, data?: any): Promise<ApiResponse<T>> => {
  try {
    const response = await apiClient.put<ApiResponse<T>>(url, data);
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const apiDelete = async <T>(url: string): Promise<ApiResponse<T>> => {
  try {
    const response = await apiClient.delete<ApiResponse<T>>(url);
    return response.data;
  } catch (error) {
    throw error;
  }
};

export default apiClient;