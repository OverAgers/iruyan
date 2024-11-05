// src/utils/seatUtils.ts
import { SeatsInfo } from "@/types/seats-info";
import { UserInfo } from "@/types/user-info";

export const occupySeat = (
  seats: SeatsInfo[],
  seatId: string,
  userInfo: UserInfo
): SeatsInfo[] => {
  let isChanged = false;
  const updatedSeats = seats.map((seat) => {
    if (seat.seatId === seatId && seat.isVacant) {
      isChanged = true;
      return {
        ...seat,
        isVacant: false,
        name: userInfo.name,
        iruyanId: userInfo.iruyanId,
        image: userInfo.avatarUrl || "",
        note: userInfo.note || "",
        task: userInfo.task || "",
      };
    }
    return seat;
  });
  return isChanged ? updatedSeats : seats;
};

export const vacateSeat = (seats: SeatsInfo[], seatId: string): SeatsInfo[] => {
  let isChanged = false;
  const updatedSeats = seats.map((seat) => {
    if (seat.seatId === seatId && !seat.isVacant) {
      isChanged = true;
      return {
        ...seat,
        isVacant: true,
        name: "空席",
        image: "",
        iruyanId: "",
        note: "",
        task: "",
      };
    }
    return seat;
  });
  return isChanged ? updatedSeats : seats;
};
