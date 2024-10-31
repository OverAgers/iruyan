// src/utils/seatUtils.ts
import { SeatsInfo } from "@/types/seats-info";
import { UserInfo } from "@/types/user-info";

export const occupySeat = (
  seats: SeatsInfo[],
  seatId: number,
  userInfo: UserInfo
): SeatsInfo[] => {
  return seats.map((seat) =>
    seat.id === seatId && seat.isVacant
      ? {
          ...seat,
          isVacant: false,
          name: userInfo.name,
          image: userInfo.avatarUrl || "",
          note: userInfo.note || "",
          task: userInfo.task || "",
        }
      : seat
  );
};

export const vacateSeat = (seats: SeatsInfo[], seatId: number): SeatsInfo[] => {
  return seats.map((seat) =>
    seat.id === seatId && !seat.isVacant
      ? {
          ...seat,
          isVacant: true,
          name: "空席",
          image: "",
          note: "",
          task: "",
        }
      : seat
  );
};
