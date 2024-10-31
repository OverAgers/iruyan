import { create } from "zustand";
import { persist } from "zustand/middleware";
import { UserInfo, StatusType } from "@/types/user-info";
import { incrementTime } from "@/utils/increment-time";

type UserStore = {
  currentUser: UserInfo | null;
  setUser: (user: UserInfo) => void;
  clearUser: () => void;
  setStatus: (status: StatusType) => void;
  incrementWorkTime: () => void;
  incrementRestTime: () => void;
  setStartTime: (time: number) => void;
  resetTimes: () => void;
  setTask: (task: string) => void;
  setNote: (note: string) => void;
  setAvatarUrl: (avatarUrl: string) => void;
  setCumulativeTime: (cumulativeTime: number) => void;
  setConsecutiveDays: (consecutiveDays: number) => void;
};

const useUserStore = create<UserStore>()(
  persist(
    (set, get) => ({
      currentUser: null,

      // ユーザー情報を設定する関数
      setUser: (user: UserInfo) => set({ currentUser: user }),

      // ユーザー情報をクリアする関数（ログアウト用）
      clearUser: () => set({ currentUser: null }),

      // ステータスを設定する関数
      setStatus: (status) => {
        const user = get().currentUser;
        if (user) {
          set({
            currentUser: {
              ...user,
              status,
            },
          });
        }
      },

      // 作業時間を増加させる関数
      incrementWorkTime: () => {
        const user = get().currentUser;
        if (user) {
          set({
            currentUser: {
              ...user,
              workTime: incrementTime(user.workTime),
            },
          });
        }
      },

      // 休憩時間を増加させる関数
      incrementRestTime: () => {
        const user = get().currentUser;
        if (user) {
          set({
            currentUser: {
              ...user,
              restTime: incrementTime(user.restTime),
            },
          });
        }
      },

      // 開始時間を設定する関数
      setStartTime: (time) => {
        const user = get().currentUser;
        if (user) {
          set({
            currentUser: {
              ...user,
              startTime: time,
            },
          });
        }
      },

      // 時間をリセットする関数
      resetTimes: () => {
        const user = get().currentUser;
        if (user) {
          set({
            currentUser: {
              ...user,
              workTime: 0,
              restTime: 0,
            },
          });
        }
      },

      setTask: (task) => {
        const user = get().currentUser;
        if (user) {
          set({
            currentUser: {
              ...user,
              task,
            },
          });
        }
      },

      setNote: (note) => {
        const user = get().currentUser;
        if (user) {
          set({
            currentUser: {
              ...user,
              note,
            },
          });
        }
      },

      setAvatarUrl: (avatarUrl) => {
        const user = get().currentUser;
        if (user) {
          set({
            currentUser: {
              ...user,
              avatarUrl,
            },
          });
        }
      },

      setCumulativeTime: (cumulativeTime) => {
        const user = get().currentUser;
        if (user) {
          set({
            currentUser: {
              ...user,
              cumulativeTime,
            },
          });
        }
      },

      setConsecutiveDays: (consecutiveDays) => {
        const user = get().currentUser;
        if (user) {
          set({
            currentUser: {
              ...user,
              consecutiveDays,
            },
          });
        }
      },
    }),
    {
      name: "user-store", // ストレージのキー名
    }
  )
);

export default useUserStore;
