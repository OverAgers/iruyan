import SeatCard from "@/components/ui/card/seat-card";
import { Box, Grid2 } from "@mui/material";
import useSeatStore from "@/stores/seats-store";

export default function Seats() {
  const { seats, sitOnSeat, leaveSeat } = useSeatStore();

  return (
    <Box sx={{width: 800 }}>
      <Grid2 container spacing={2}>
        {seats.map((seat) => (
          <Grid2 key={seat.id}>
            <SeatCard
              {...seat}
              onClick={() =>
                seat.isVacant ? sitOnSeat(seat.id) : leaveSeat(seat.id)
              }
            />
          </Grid2>
        ))}
      </Grid2>
    </Box>
  );
}
