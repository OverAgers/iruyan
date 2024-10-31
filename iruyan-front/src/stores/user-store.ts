// src/stores/userStore.ts
import { create } from "zustand";
import { UserInfo } from "@/types/user-info";

type UserStore = {
  currentUser: UserInfo | null;
  setUser: (user: UserInfo) => void;
  clearUser: () => void;
};

const useUserStore = create<UserStore>((set) => ({
  currentUser: null,

  // ユーザー情報を設定する関数
  setUser: (user: UserInfo) => set({ currentUser: user }),

  // ユーザー情報をクリアする関数（ログアウト用）
  clearUser: () => set({ currentUser: null }),
}));

export default useUserStore;
