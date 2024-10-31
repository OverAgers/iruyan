export type StatusType = "working" | "resting" | "idle";

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
};
