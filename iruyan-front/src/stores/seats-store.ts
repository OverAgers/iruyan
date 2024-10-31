// src/stores/seat-store.ts
import { create } from "zustand";
import { SeatsInfo } from "@/types/seats-info";
import { UserInfo } from "@/types/user-info";
import { occupySeat, vacateSeat } from "@/utils/seat-utils";

export type SeatStore = {
  seats: SeatsInfo[];
  sitOnSeat: (id: number, userInfo: UserInfo) => void;
  leaveSeat: (id: number) => void;
};

const initialSeats: SeatsInfo[] = [
  { id: 1, isVacant: true, name: "空席", image: "", note: "", task: "" },
  { id: 2, isVacant: true, name: "空席", image: "", note: "", task: "" },
  { id: 3, isVacant: true, name: "空席", image: "", note: "", task: "" },
  { id: 4, isVacant: true, name: "空席", image: "", note: "", task: "" },
  { id: 5, isVacant: true, name: "空席", image: "", note: "", task: "" },
  { id: 6, isVacant: true, name: "空席", image: "", note: "", task: "" },
  { id: 7, isVacant: true, name: "空席", image: "", note: "", task: "" },
  { id: 8, isVacant: true, name: "空席", image: "", note: "", task: "" },
  { id: 9, isVacant: true, name: "空席", image: "", note: "", task: "" },
  { id: 10, isVacant: true, name: "空席", image: "", note: "", task: "" },
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
