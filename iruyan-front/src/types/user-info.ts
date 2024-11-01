export type StatusType = "working" | "resting" | "idle";

export type UserInfo = {
  iruyanID: string;
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
