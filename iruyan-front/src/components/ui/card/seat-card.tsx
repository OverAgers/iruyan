import React from "react";
import Typography from "@mui/joy/Typography";
import Avatar from "@mui/joy/Avatar";
import Card from "@mui/joy/Card";
import Box from "@mui/joy/Box";

type SeatCardProps = {
  seatId: string;
  isVacant: boolean;
  name?: string;
  image?: string;
  note?: string;
  task?: string;
  isCurrentUser: boolean; // ←追加
  onClick: () => void;
};

const SeatCard: React.FC<SeatCardProps> = ({
  isVacant,
  name,
  image,
  note,
  task,
  onClick,
}) => {
  return (
    <Card
      sx={{
        padding: 0,
        margin: 1,
        width: 370,
        backgroundColor: "#f5f5f5",
        cursor: "pointer",
      }}
      onClick={onClick}
    >
      <Box display="flex" justifyContent="space-between" sx={{ px: 4, py: 2 }}>
        <Box display="flex" flexDirection="column" alignItems="center">
          <Typography level="title-lg" fontWeight="bold">
            {isVacant ? "空席" : name || "（名前なし）"}
          </Typography>
          <Avatar src={image} sx={{ width: 48, height: 48, marginTop: 1 }} />
        </Box>
        {!isVacant && (
          <Box display="flex" flexDirection="column" alignItems="flex-start">
            <Typography level="title-lg" fontWeight="bold" color="primary">
              {task || ""}
            </Typography>
            <Typography level="title-lg" mt={1}>
              {note || ""}
            </Typography>
          </Box>
        )}
      </Box>
    </Card>
  );
};

export default SeatCard;
