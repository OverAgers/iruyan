// src/stores/seat-store.ts
import { create } from "zustand";
import { SeatsInfo } from "@/types/seats-info";
import { UserInfo } from "@/types/user-info";
import { occupySeat, vacateSeat } from "@/utils/seat-utils";

export type SeatStore = {
  seats: SeatsInfo[];
  sitOnSeat: (id: string, userInfo: UserInfo) => void;
  leaveSeat: (id: string) => void;
};

const initialSeats: SeatsInfo[] = [

];

const useSeatStore = create<SeatStore>((set) => ({
  seats: initialSeats,

  sitOnSeat: (id, userInfo) =>
    set((state) => ({
      seats: occupySeat(state.seats, id, userInfo),
    })),

  leaveSeat: (id) =>
    set((state) => ({
      seats: vacateSeat(state.seats, id),
    })),
}));

export default useSeatStore;
