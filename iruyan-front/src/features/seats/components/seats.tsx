import React from "react";
import SeatCard from "@/components/ui/card/seat-card";
import { Box, Grid2 } from "@mui/material";
import useSeatStore from "@/stores/seats-store";
import useUserStore from "@/stores/user-store";
import { useEffect } from "react";

export default function Seats() {
  const { seats, sitOnSeat, leaveSeat } = useSeatStore();
  const { currentUser } = useUserStore();

  useEffect(() => {
    if (currentUser) {
      const occupiedSeat = seats.find(
        (seat) => !seat.isVacant && seat.name === currentUser.name
      );
      if (occupiedSeat) {
        sitOnSeat(occupiedSeat.seatId, currentUser);
      }
    }
  }, [currentUser, seats]);

  const handleSeatClick = (seatId: string, isVacant: boolean) => {
    if (isVacant) {
      if (currentUser) {
        sitOnSeat(seatId, currentUser);
      } else {
        alert("ログインしてください");
      }
    } else {
      // 座っているのが自分自身かどうかを確認
      const seat = seats.find((s) => s.seatId === seatId);
      if (seat && seat.name === currentUser?.name) {
        leaveSeat(seatId);
      } else {
        alert("この席は他のユーザーが使用中です");
      }
    }
  };

  return (
    <Box sx={{ width: 800 }}>
      <Grid2 container spacing={2}>
        {seats.map((seat) => (
          <Grid2 key={seat.seatId}>
            <SeatCard
              {...seat}
              onClick={() => handleSeatClick(seat.seatId, seat.isVacant)}
            />
          </Grid2>
        ))}
      </Grid2>
    </Box>
  );
}
