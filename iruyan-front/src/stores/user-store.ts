import { create } from "zustand";
import { persist } from "zustand/middleware";
import { UserInfo, StatusType } from "@/types/user-info";

type UserStore = {
  currentUser: UserInfo | null;
  setUser: (user: UserInfo) => void;
  clearUser: () => void;
  setStatus: (status: StatusType) => void;
  incrementWorkTime: (seconds: number) => void;
  incrementRestTime: (seconds: number) => void;
  setStartTime: (time: number) => void;
  resetTimes: () => void;
};

const useUserStore = create<UserStore>()(
  persist(
    (set, get) => ({
      currentUser: null,

      // ユーザー情報を設定する関数
      setUser: (user: UserInfo) =>
        set({
          currentUser: {
            ...user,
            status: "idle",
            workTime: 0,
            restTime: 0,
            startTime: 0,
          },
        }),

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
      incrementWorkTime: (seconds) => {
        const user = get().currentUser;
        if (user) {
          set({
            currentUser: {
              ...user,
              workTime: user.workTime + seconds,
            },
          });
        }
      },

      // 休憩時間を増加させる関数
      incrementRestTime: (seconds) => {
        const user = get().currentUser;
        if (user) {
          set({
            currentUser: {
              ...user,
              restTime: user.restTime + seconds,
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
    }),
    {
      name: "user-store", // ストレージのキー名
    }
  )
);

export default useUserStore;
