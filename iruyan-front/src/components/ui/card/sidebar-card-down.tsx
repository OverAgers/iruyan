import {
  Avatar,
  Box,
  Button,
  ButtonGroup,
  Card,
  List,
  ListItem,
  ToggleButtonGroup,
  Typography,
} from "@mui/joy";
import SubButton from "../button/sub-button";
import { useState } from "react";

export default function SidebarCardUp() {
  const [workTime, setWorkTime] = useState("12:10");
  const [restTime, setRestTime] = useState("12:10");
  return (
    <Card size="md" sx={{ mb: 2, height: "333px", width: "339px" }}>
      <Box>
        <Box display={"flex"} justifyContent={"center"}>
          <ToggleButtonGroup variant="outlined">
            <SubButton title={"作業時間"} size={"sm"}></SubButton>
            <SubButton title={"タイマー"} size={"sm"}></SubButton>
          </ToggleButtonGroup>
        </Box>
        <Box>
          <Box mb={2}>
            <Box display={"flex"}>
              <Typography>作業時間</Typography>
              <Typography>{workTime}</Typography>
            </Box>
            <Box display={"flex"}>
              <Typography>休憩時間</Typography>
              <Typography>{restTime}</Typography>
            </Box>
          </Box>
          <Box
            display={"flex"}
            flexDirection={"column"}
            justifyContent={"center"}
          >
            <Box display={"flex"}>
              <SubButton title={"鬼集中"} size={"md"}></SubButton>
              <SubButton title={"休憩"} size={"md"}></SubButton>
            </Box>
            <SubButton title={"席を離れる"} size={"md"}></SubButton>
          </Box>
        </Box>
      </Box>
    </Card>
  );
}
