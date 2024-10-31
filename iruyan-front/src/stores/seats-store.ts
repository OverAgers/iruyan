// src/stores/seatStore.ts
import { create } from "zustand";
import { SeatsInfo } from "@/types/seats-info";
import useUserStore from "./user-store";

export type SeatStore = {
  seats: SeatsInfo[];
  sitOnSeat: (id: number) => void;
  leaveSeat: (id: number) => void;
};

const useSeatStore = create<SeatStore>((set) => ({
  seats: [
    {
      id: 1,
      isVacant: true,
      name: "空席",
      image: "",
      note: "つぶやき",
      task: "",
    },
    {
      id: 2,
      isVacant: true,
      name: "空席",
      image: "",
      note: "つぶやき",
      task: "",
    },
    {
      id: 3,
      isVacant: true,
      name: "空席",
      image: "",
      note: "つぶやき",
      task: "",
    },
    {
      id: 4,
      isVacant: true,
      name: "空席",
      image: "",
      note: "つぶやき",
      task: "",
    },
    {
      id: 5,
      isVacant: true,
      name: "空席",
      image: "",
      note: "つぶやき",
      task: "",
    },
    {
      id: 6,
      isVacant: true,
      name: "空席",
      image: "",
      note: "つぶやき",
      task: "",
    },
    {
      id: 7,
      isVacant: true,
      name: "空席",
      image: "",
      note: "つぶやき",
      task: "",
    },
    {
      id: 8,
      isVacant: true,
      name: "空席",
      image: "",
      note: "つぶやき",
      task: "",
    },
  ],

  sitOnSeat: (id: number) =>
    set((state) => {
      const currentUser = useUserStore.getState().currentUser;
      return {
        seats: state.seats.map((seat) =>
          seat.id === id && seat.isVacant && currentUser
            ? {
              ...seat,
              isVacant: false,
              name: currentUser.name,
              image: currentUser.avatarUrl || "",
              note: currentUser.note || "",
            }
            : seat
        ),
      };
    }),

  leaveSeat: (id: number) =>
    set((state) => {
      return {
      seats: state.seats.map((seat) =>
        seat.id === id && !seat.isVacant
          ? {
            ...seat,
            isVacant: true,
            name: "空席",
            image: "",
            note: "つぶやき",
            task: "",
          }
          : seat
      ),
    };
    }),
}));

export default useSeatStore;
