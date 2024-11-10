// src/stores/seat-store.ts
import { create } from 'zustand';
import { SeatsInfo } from '@/types/seats-info';
import { UserInfo } from '@/types/user-info';
import { occupySeat, vacateSeat } from '@/utils/seat-utils';

export type SeatStore = {
  seats: SeatsInfo[];
  sitOnSeat: (id: string, userInfo: UserInfo) => void;
  leaveSeat: (id: string) => void;
  moveSeat: (newSeatId: string, userInfo: UserInfo) => void; // 新しいアクションを追加
};

const initialSeats: SeatsInfo[] = [
  { seatId: '1', roomId: 'aaa', seatNumber: 1, isVacant: true },
  { seatId: '2', roomId: 'aaa', seatNumber: 2, isVacant: true },
  { seatId: '3', roomId: 'aaa', seatNumber: 3, isVacant: true },
  { seatId: '4', roomId: 'aaa', seatNumber: 4, isVacant: true },
  { seatId: '5', roomId: 'aaa', seatNumber: 5, isVacant: true },
  { seatId: '6', roomId: 'aaa', seatNumber: 6, isVacant: true },
  { seatId: '7', roomId: 'aaa', seatNumber: 7, isVacant: true },
  { seatId: '8', roomId: 'aaa', seatNumber: 8, isVacant: true },
  { seatId: '9', roomId: 'aaa', seatNumber: 9, isVacant: true },
  { seatId: '10', roomId: 'aaa', seatNumber: 10, isVacant: true },
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

  moveSeat: (newSeatId, userInfo) => {
    set((state) => {
      const currentSeat = state.seats.find(
        (seat) => seat.iruyanId === userInfo.iruyanId && !seat.isVacant
      );

      let updatedSeats = state.seats;

      if (currentSeat) {
        updatedSeats = vacateSeat(updatedSeats, currentSeat.seatId);
      }

      updatedSeats = occupySeat(updatedSeats, newSeatId, userInfo);

      return { seats: updatedSeats };
    });
  },
}));

export default useSeatStore;
