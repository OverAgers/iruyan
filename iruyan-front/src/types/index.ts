// Common types
export type StatusType = "working" | "resting" | "idle";

// User related types
export type UserInfo = {
  iruyanId: string;
  name: string;
  email: string;
  avatarUrl?: string;
  task?: string;
  note?: string;
  status: StatusType;
  workTime: number;
  restTime: number;
  startTime: number;
  cumulativeTime?: number;
  consecutiveDays?: number;
};

// Room related types
export type RoomInfo = {
  id: string;
  name: string;
  description?: string;
  maxSeats: number;
  currentSeats: number;
  isActive: boolean;
  createdAt: string;
  updatedAt: string;
};

export type RoomsInfo = RoomInfo[];

// Seat related types
export type SeatInfo = {
  id: string;
  roomId: string;
  seatNumber: number;
  isOccupied: boolean;
  userId?: string;
  userName?: string;
  userStatus?: StatusType;
  startTime?: number;
  workTime?: number;
  restTime?: number;
};

export type SeatsInfo = SeatInfo[];

// API Response types
export type ApiResponse<T> = {
  success: boolean;
  data?: T;
  error?: string;
  message?: string;
};

// Form types
export type LoginFormData = {
  email: string;
  password: string;
};

export type RegisterFormData = {
  name: string;
  email: string;
  password: string;
  confirmPassword: string;
}; 