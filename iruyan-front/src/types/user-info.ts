export type StatusType = 'working' | 'resting' | 'idle';

export type UserInfo = {
  iruyanId: string;
  userName: string;
  email: string;
  avatarUrl?: string;
  task?: string;
  note?: string;
  status?: StatusType;
  workTime?: number;
  restTime?: number;
  startTime?: number;
  cumulativeTime?: number;
  consecutiveDays?: number;
};
