/**
 * API エンドポイントとレスポンス型定義
 * バックエンドAPIとの型安全な通信を実現
 */

// Base API Response
export interface BaseApiResponse<T = unknown> {
  success: boolean;
  data?: T;
  error?: string;
  message?: string;
  status?: number;
}

// Error Response
export interface ApiErrorResponse extends Record<string, unknown> {
  success: false;
  error: string;
  message?: string;
  status: number;
  details?: Record<string, unknown>;
}

// Success Response
export interface ApiSuccessResponse<T> {
  success: true;
  data: T;
  message?: string;
  status: number;
}

// Union type for all responses
export type ApiResponse<T> = ApiSuccessResponse<T> | ApiErrorResponse;

// ===== Authentication API Types =====
export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  name: string;
  email: string;
  password: string;
}

export interface AuthResponse {
  token: string;
  user: {
    iruyanId: string;
    name: string;
    email: string;
    avatarUrl?: string;
  };
  expiresAt: string;
}

// ===== User API Types =====
export interface UserProfile {
  iruyanId: string;
  name: string;
  email: string;
  avatarUrl?: string;
  createdAt: string;
  updatedAt: string;
}

export interface UserStats {
  cumulativeTime: number;
  consecutiveDays: number;
  todaysWorkTime: number;
  todaysRestTime: number;
  weeklyStats: Array<{
    day: string;
    workTime: number;
    restTime: number;
  }>;
}

export interface UpdateUserRequest {
  name?: string;
  avatarUrl?: string;
  task?: string;
  note?: string;
}

// ===== Room API Types =====
export interface Room {
  id: string;
  name: string;
  description?: string;
  maxSeats: number;
  currentOccupiedSeats: number;
  isActive: boolean;
  createdAt: string;
  updatedAt: string;
}

export interface CreateRoomRequest {
  name: string;
  description?: string;
  maxSeats: number;
}

export interface UpdateRoomRequest {
  name?: string;
  description?: string;
  maxSeats?: number;
  isActive?: boolean;
}

// ===== Seat API Types =====
export interface Seat {
  seatId: string;
  roomId: string;
  seatNumber: number;
  isVacant: boolean;
  // Occupied seat info
  iruyanId?: string;
  userName?: string;
  userImage?: string;
  task?: string;
  note?: string;
  startTime?: string;
  status?: 'working' | 'resting' | 'idle';
}

export interface SeatStatusResponse {
  seat_number: number;
  iruyan_id?: string;
  user_name?: string;
  user_image?: string;
  task?: string;
  note?: string;
  start_time?: string;
  status?: 'working' | 'resting' | 'idle';
}

export interface TakeSeatRequest {
  iruyanId: string;
}

export interface LeaveSeatRequest {
  iruyanId: string;
}

// ===== Work Time API Types =====
export interface WorkTimeEntry {
  id: string;
  userId: string;
  roomId: string;
  seatNumber: number;
  startTime: string;
  endTime?: string;
  workDuration: number;
  restDuration: number;
  task?: string;
  note?: string;
  status: 'active' | 'completed';
}

export interface StartWorkRequest {
  roomId: string;
  seatNumber: number;
  task?: string;
  note?: string;
}

export interface UpdateWorkStatusRequest {
  status: 'working' | 'resting' | 'idle';
  task?: string;
  note?: string;
}

// ===== API Endpoints Map =====
export interface ApiEndpoints {
  // Auth
  'POST /auth/login': {
    request: LoginRequest;
    response: AuthResponse;
  };
  'POST /auth/register': {
    request: RegisterRequest;
    response: AuthResponse;
  };
  'POST /auth/logout': {
    request: never;
    response: { message: string };
  };

  // User
  'GET /users/profile': {
    request: never;
    response: UserProfile;
  };
  'PUT /users/profile': {
    request: UpdateUserRequest;
    response: UserProfile;
  };
  'GET /users/stats': {
    request: never;
    response: UserStats;
  };

  // Rooms
  'GET /rooms': {
    request: never;
    response: Room[];
  };
  'GET /rooms/:roomId': {
    request: never;
    response: Room;
  };
  'POST /rooms': {
    request: CreateRoomRequest;
    response: Room;
  };
  'PUT /rooms/:roomId': {
    request: UpdateRoomRequest;
    response: Room;
  };
  'DELETE /rooms/:roomId': {
    request: never;
    response: { message: string };
  };

  // Seats
  'GET /rooms/:roomId/seats/status': {
    request: never;
    response: SeatStatusResponse[];
  };
  'PUT /rooms/:roomId/seats/:seatNumber/take': {
    request: TakeSeatRequest;
    response: { message: string };
  };
  'PUT /rooms/:roomId/seats/:seatNumber/leave': {
    request: LeaveSeatRequest;
    response: { message: string };
  };

  // Work Time
  'GET /work-time/current': {
    request: never;
    response: WorkTimeEntry | null;
  };
  'POST /work-time/start': {
    request: StartWorkRequest;
    response: WorkTimeEntry;
  };
  'PUT /work-time/status': {
    request: UpdateWorkStatusRequest;
    response: WorkTimeEntry;
  };
  'POST /work-time/stop': {
    request: never;
    response: WorkTimeEntry;
  };
}

// Helper type to extract request/response types
export type ApiRequestType<T extends keyof ApiEndpoints> = ApiEndpoints[T]['request'];
export type ApiResponseType<T extends keyof ApiEndpoints> = ApiEndpoints[T]['response'];

// Generic API call type
export type ApiCall<T extends keyof ApiEndpoints> = (
  params?: ApiRequestType<T>
) => Promise<ApiResponse<ApiResponseType<T>>>;