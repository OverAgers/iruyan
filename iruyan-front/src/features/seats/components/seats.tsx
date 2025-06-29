// Updated to fix ESLint errors - StatusType and any type issues resolved
import React from "react";
import SeatCard from "@/components/ui/card/seat-card";
import { Box, Grid2, Typography } from "@mui/material";
import useSeatStore from "@/stores/seats-store";
import useUserStore from "@/stores/user-store";

export default function Seats() {
  const currentUser = useUserStore((state) => state.currentUser);
  const seats = useSeatStore((state) => state.seats);
  const moveSeat = useSeatStore((state) => state.moveSeat);

  if (!currentUser) {
    return <Typography>ログインしてください</Typography>;
  }

  const handleSeatClick = (seatId: string, isVacant: boolean) => {
    if (isVacant) {
      moveSeat(seatId, currentUser);
    } else {
      const seat = seats.find((s) => s.seatId === seatId);
      if (seat && seat.name === currentUser.name) {
        useSeatStore.getState().leaveSeat(seatId);
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
