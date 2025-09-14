import { create } from "zustand";
import { persist } from "zustand/middleware";
import { devtools } from "zustand/middleware";
import type { StateCreator } from "zustand";
import { UserInfo, StatusType } from "@/types/user-info";
import { incrementTime } from "@/utils/time-utils";

interface UserState {
  currentUser: UserInfo | null;
}

interface UserActions {
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
}

type UserStore = UserState & UserActions;

type UserStoreSlice = StateCreator<UserStore, [], [], UserStore>;

const createUserStore: UserStoreSlice = (set, get) => ({
  // State
  currentUser: null,

  // Actions
  setUser: (user: UserInfo): void => {
    set({ currentUser: user }, false);
  },

  clearUser: (): void => {
    set({ currentUser: null }, false);
  },

  setStatus: (status: StatusType): void => {
    const user = get().currentUser;
    if (!user) return;

    set(
      {
        currentUser: {
          ...user,
          status,
        },
      },
      false
    );
  },

  incrementWorkTime: (): void => {
    const user = get().currentUser;
    if (!user) return;

    set(
      {
        currentUser: {
          ...user,
          workTime: incrementTime(user.workTime),
        },
      },
      false
    );
  },

  incrementRestTime: (): void => {
    const user = get().currentUser;
    if (!user) return;

    set(
      {
        currentUser: {
          ...user,
          restTime: incrementTime(user.restTime),
        },
      },
      false
    );
  },

  setStartTime: (time: number): void => {
    const user = get().currentUser;
    if (!user) return;

    set(
      {
        currentUser: {
          ...user,
          startTime: time,
        },
      },
      false
    );
  },

  resetTimes: (): void => {
    const user = get().currentUser;
    if (!user) return;

    set(
      {
        currentUser: {
          ...user,
          workTime: 0,
          restTime: 0,
        },
      },
      false
    );
  },

  setTask: (task: string): void => {
    const user = get().currentUser;
    if (!user) return;

    set(
      {
        currentUser: {
          ...user,
          task,
        },
      },
      false
    );
  },

  setNote: (note: string): void => {
    const user = get().currentUser;
    if (!user) return;

    set(
      {
        currentUser: {
          ...user,
          note,
        },
      },
      false
    );
  },

  setAvatarUrl: (avatarUrl: string): void => {
    const user = get().currentUser;
    if (!user) return;

    set(
      {
        currentUser: {
          ...user,
          avatarUrl,
        },
      },
      false
    );
  },

  setCumulativeTime: (cumulativeTime: number): void => {
    const user = get().currentUser;
    if (!user) return;

    set(
      {
        currentUser: {
          ...user,
          cumulativeTime,
        },
      },
      false
    );
  },

  setConsecutiveDays: (consecutiveDays: number): void => {
    const user = get().currentUser;
    if (!user) return;

    set(
      {
        currentUser: {
          ...user,
          consecutiveDays,
        },
      },
      false
    );
  },
});

const useUserStore = create<UserStore>()(
  devtools(
    persist(createUserStore, {
      name: "user-store",
      partialize: (state) => ({ currentUser: state.currentUser }),
    }),
    {
      name: "user-store",
    }
  )
);

export default useUserStore;
