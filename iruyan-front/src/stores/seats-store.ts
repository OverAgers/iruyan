import { create } from "zustand";
import { devtools } from "zustand/middleware";
import type { StateCreator } from "zustand";
import { SeatsInfo } from "@/types/seats-info";
import { UserInfo } from "@/types/user-info";
import { occupySeat, vacateSeat } from "@/utils/seat-utils";

interface SeatState {
  seats: SeatsInfo[];
}

interface SeatActions {
  sitOnSeat: (id: string, userInfo: UserInfo) => void;
  leaveSeat: (id: string) => void;
  moveSeat: (newSeatId: string, userInfo: UserInfo) => void;
  setSeatUser: (seatNumber: number, userInfo: SeatsInfo) => void;
}

export type SeatStore = SeatState & SeatActions;

type SeatStoreSlice = StateCreator<SeatStore, [], [], SeatStore>;

const initialSeats: SeatsInfo[] = [
  { seatId: "1", roomId: "aaa", seatNumber: 1, isVacant: true },
  { seatId: "2", roomId: "aaa", seatNumber: 2, isVacant: true },
  { seatId: "3", roomId: "aaa", seatNumber: 3, isVacant: true },
  { seatId: "4", roomId: "aaa", seatNumber: 4, isVacant: true },
  { seatId: "5", roomId: "aaa", seatNumber: 5, isVacant: true },
  { seatId: "6", roomId: "aaa", seatNumber: 6, isVacant: true },
  { seatId: "7", roomId: "aaa", seatNumber: 7, isVacant: true },
  { seatId: "8", roomId: "aaa", seatNumber: 8, isVacant: true },
  { seatId: "9", roomId: "aaa", seatNumber: 9, isVacant: true },
  { seatId: "10", roomId: "aaa", seatNumber: 10, isVacant: true },
];

const createSeatStore: SeatStoreSlice = (set) => ({
  // State
  seats: initialSeats,

  // Actions
  sitOnSeat: (id: string, userInfo: UserInfo): void => {
    set(
      (state) => ({
        seats: occupySeat(state.seats, id, userInfo),
      }),
      false
    );
  },

  leaveSeat: (id: string): void => {
    set(
      (state) => ({
        seats: vacateSeat(state.seats, id),
      }),
      false
    );
  },

  moveSeat: (newSeatId: string, userInfo: UserInfo): void => {
    set(
      (state) => {
        const vacatedSeats: SeatsInfo[] = state.seats.map((seat) =>
          seat.iruyanId === userInfo.iruyanId
            ? {
                ...seat,
                isVacant: true,
                iruyanId: "",
                userName: "",
                userImage: "",
                task: "",
                note: "",
              }
            : seat
        );

        const updatedSeats: SeatsInfo[] = vacatedSeats.map((seat) =>
          seat.seatId === newSeatId
            ? {
                ...seat,
                isVacant: false,
                iruyanId: userInfo.iruyanId,
                userName: userInfo.name,
                userImage: userInfo.avatarUrl ?? "",
                task: userInfo.task ?? "",
                note: userInfo.note ?? "",
              }
            : seat
        );

        return { seats: updatedSeats };
      },
      false
    );
  },

  setSeatUser: (seatNumber: number, userInfo: SeatsInfo): void => {
    set(
      (state) => ({
        seats: state.seats.map((seat) =>
          seat.seatNumber === seatNumber
            ? {
                ...seat,
                ...userInfo,
                isVacant: false,
              }
            : seat
        ),
      }),
      false
    );
  },
});

const useSeatStore = create<SeatStore>()(
  devtools(createSeatStore, {
    name: "seat-store",
  })
);

export default useSeatStore;
