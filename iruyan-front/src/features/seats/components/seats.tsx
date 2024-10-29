import SeatCard from "@/components/ui/seat-card/seat-card";
import { Box, Grid2 } from "@mui/material";
import { useState } from "react";

export default function Seats() {

  const initialSeats = [
    {
      id: 1,
      isVacant: true,
      name: "空席",
      image: "",
      tweet: "つぶやき",
      status: null,
    },
    {
      id: 2,
      isVacant: true,
      name: "空席",
      image: "",
      tweet: "つぶやき",
      status: null,
    },
    {
      id: 3,
      isVacant: true,
      name: "空席",
      image: "",
      tweet: "つぶやき",
      status: null,
    },
    {
      id: 4,
      isVacant: true,
      name: "空席",
      image: "",
      tweet: "つぶやき",
      status: null,
    },
    {
      id: 5,
      isVacant: true,
      name: "空席",
      image: "",
      tweet: "つぶやき",
      status: null,
    },
    {
      id: 6,
      isVacant: true,
      name: "空席",
      image: "",
      tweet: "つぶやき",
      status: null,
    },
    {
      id: 7,
      isVacant: true,
      name: "空席",
      image: "",
      tweet: "つぶやき",
      status: null,
    },
    {
      id: 8,
      isVacant: true,
      name: "空席",
      image: "",
      tweet: "つぶやき",
      status: null,
    },
  ];

  const [seats, setSeats] = useState(initialSeats);

  const currentUser = {
    name: "山田太郎",
    image: "https://source.unsplash.com/random",
    tweet: "今日も一日がんばるぞい！",
  };

  const handleSeatClick = (id:number) => {
    setSeats((prevSeats) =>
      prevSeats.map((seat) =>
        seat.id === id && seat.isVacant
          ? {
              ...seat,
              isVacant: false,
              name: currentUser.name,
              image: currentUser.image,
              tweet: currentUser.tweet,
            }
          : seat
      )
    );
  };
  return (
    <Box sx={{ backgroundColor: "#8B4513", width: 800}}>
      <Grid2 container spacing={2}>
        {seats.map((seat) => (
          <Grid2 key={seat.id}>
            <SeatCard {...seat} onClick={() => handleSeatClick(seat.id)} />
          </Grid2>
        ))}
      </Grid2>
    </Box>
  );
}