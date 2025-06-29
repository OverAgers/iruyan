// src/stores/seat-store.ts
import { create } from "zustand";
import { SeatsInfo } from "@/types/seats-info";
import { UserInfo } from "@/types/user-info";
import { occupySeat, vacateSeat } from "@/utils/seat-utils";

export type SeatStore = {
  seats: SeatsInfo[];
  sitOnSeat: (id: string, userInfo: UserInfo) => void;
  leaveSeat: (id: string, userInfo: UserInfo) => void;
  moveSeat: (newSeatId: string, userInfo: UserInfo) => void;
  setSeatUser: (seatNumber: number, userInfo: SeatsInfo) => void;
};

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

  moveSeat: (newSeatId, userInfo) =>
    set((state) => {
      const vacatedSeats = state.seats.map((seat) =>
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

      const updatedSeats = vacatedSeats.map((seat) =>
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
    }),

  setSeatUser: (seatNumber, userInfo) =>
    set((state) => ({
      seats: state.seats.map((seat) =>
        seat.seatNumber === seatNumber
          ? {
              ...seat,
              ...userInfo,
              isVacant: false,
            }
          : seat
      ),
    })),
}));

export default useSeatStore;
