/**
 * 型安全なAPIサービス層
 * 各エンドポイントに対応する関数を提供
 */

import { makeApiCall, buildApiUrl } from "@/lib/api-client";
import type {
  ApiResponse,
  LoginRequest,
  RegisterRequest,
  AuthResponse,
  UserProfile,
  UserStats,
  UpdateUserRequest,
  Room,
  CreateRoomRequest,
  UpdateRoomRequest,
  SeatStatusResponse,
  TakeSeatRequest,
  LeaveSeatRequest,
  WorkTimeEntry,
  StartWorkRequest,
  UpdateWorkStatusRequest,
} from "@/types/api";

// ===== Authentication Services =====
export const authService = {
  async login(credentials: LoginRequest): Promise<ApiResponse<AuthResponse>> {
    return makeApiCall("POST", "/auth/login", credentials);
  },

  async register(userData: RegisterRequest): Promise<ApiResponse<AuthResponse>> {
    return makeApiCall("POST", "/auth/register", userData);
  },

  async logout(): Promise<ApiResponse<{ message: string }>> {
    return makeApiCall("POST", "/auth/logout");
  },
};

// ===== User Services =====
export const userService = {
  async getProfile(): Promise<ApiResponse<UserProfile>> {
    return makeApiCall("GET", "/users/profile");
  },

  async updateProfile(data: UpdateUserRequest): Promise<ApiResponse<UserProfile>> {
    return makeApiCall("PUT", "/users/profile", data);
  },

  async getStats(): Promise<ApiResponse<UserStats>> {
    return makeApiCall("GET", "/users/stats");
  },
};

// ===== Room Services =====
export const roomService = {
  async getRooms(): Promise<ApiResponse<Room[]>> {
    return makeApiCall("GET", "/rooms");
  },

  async getRoom(roomId: string): Promise<ApiResponse<Room>> {
    const url = buildApiUrl("/rooms/:roomId", { roomId });
    return makeApiCall("GET", url);
  },

  async createRoom(roomData: CreateRoomRequest): Promise<ApiResponse<Room>> {
    return makeApiCall("POST", "/rooms", roomData);
  },

  async updateRoom(roomId: string, roomData: UpdateRoomRequest): Promise<ApiResponse<Room>> {
    const url = buildApiUrl("/rooms/:roomId", { roomId });
    return makeApiCall("PUT", url, roomData);
  },

  async deleteRoom(roomId: string): Promise<ApiResponse<{ message: string }>> {
    const url = buildApiUrl("/rooms/:roomId", { roomId });
    return makeApiCall("DELETE", url);
  },
};

// ===== Seat Services =====
export const seatService = {
  async getSeatStatus(roomId: string): Promise<ApiResponse<SeatStatusResponse[]>> {
    const url = buildApiUrl("/rooms/:roomId/seats/status", { roomId });
    return makeApiCall("GET", url);
  },

  async takeSeat(
    roomId: string,
    seatNumber: number,
    data: TakeSeatRequest
  ): Promise<ApiResponse<{ message: string }>> {
    const url = buildApiUrl("/rooms/:roomId/seats/:seatNumber/take", {
      roomId,
      seatNumber,
    });
    return makeApiCall("PUT", url, data);
  },

  async leaveSeat(
    roomId: string,
    seatNumber: number,
    data: LeaveSeatRequest
  ): Promise<ApiResponse<{ message: string }>> {
    const url = buildApiUrl("/rooms/:roomId/seats/:seatNumber/leave", {
      roomId,
      seatNumber,
    });
    return makeApiCall("PUT", url, data);
  },
};

// ===== Work Time Services =====
export const workTimeService = {
  async getCurrentWorkTime(): Promise<ApiResponse<WorkTimeEntry | null>> {
    return makeApiCall("GET", "/work-time/current");
  },

  async startWork(data: StartWorkRequest): Promise<ApiResponse<WorkTimeEntry>> {
    return makeApiCall("POST", "/work-time/start", data);
  },

  async updateStatus(data: UpdateWorkStatusRequest): Promise<ApiResponse<WorkTimeEntry>> {
    return makeApiCall("PUT", "/work-time/status", data);
  },

  async stopWork(): Promise<ApiResponse<WorkTimeEntry>> {
    return makeApiCall("POST", "/work-time/stop");
  },
};

// ===== Combined API Service Export =====
export const apiService = {
  auth: authService,
  user: userService,
  room: roomService,
  seat: seatService,
  workTime: workTimeService,
} as const;

export default apiService;