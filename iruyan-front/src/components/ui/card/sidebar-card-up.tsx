import {
  Avatar,
  Box,
  Button,
  Card,
  List,
  ListItem,
  Typography,
} from "@mui/joy";
import SubButton from "../button/sub-button";
import { useEffect, useState } from "react";
import { TextField } from "@mui/material";
import useUserStore from "@/stores/user-store";

export default function SidebarCardUp() {
  const { currentUser, setUser } = useUserStore();

  const [name, setName] = useState(currentUser?.name || "名前");
  const [startTime, setStartTime] = useState("12:10");
  const [totalTime, setTotalTime] = useState(1);
  const [stay, setStay] = useState(1);
  const [task, setTask] = useState(currentUser?.task || "");
  const [note, setNote] = useState(currentUser?.note || "");

  const handleUpdateTask = () => {
    if (!currentUser) return;
    setUser({ ...currentUser, task });
  };

  const handleUpdateNote = () => {
    if (!currentUser) return;
    setUser({ ...currentUser, note });
    localStorage.setItem("note", note);
  };

  useEffect(() => {
    if (!currentUser) return;
    setName(currentUser.name);
  }, [currentUser]);


  return (
    <Card size="md" sx={{ mb: 2, height: "351px", width: "339px" }}>
      <Box>
        <Box
          justifyContent={"space-between"}
          display={"flex"}
          alignItems={"center"}
        >
          <Box display={"flex"} alignItems={"center"}>
            <Avatar />
            <Typography ml={2}>{name}</Typography>
          </Box>
          <SubButton title="来店記録" link="/user" size="lg" />
        </Box>
        <List marker="disc" size="sm">
          <ListItem>入店時間: {startTime}</ListItem>
          <ListItem>累計集中時間: {totalTime}時間</ListItem>
          <ListItem>居る連チャン: {stay}日</ListItem>
        </List>
        <Box mb={2}>
          <Typography fontWeight="bold" mb={1}>
            作業内容を追加
          </Typography>
          <Box display="flex" alignItems="center">
            <TextField
              variant="outlined"
              placeholder="テスト勉強"
              size="small"
              sx={{ flex: 1, mr: 1 }}
              value={task}
              onChange={(e) => setTask(e.target.value)}
            />
            <Button
              color="success"
              sx={{ borderRadius: 1, padding: "6px 12px" }}
              onClick={handleUpdateTask}
            >
              更新
            </Button>
          </Box>
        </Box>
        <Box>
          <Typography fontWeight="bold" mb={1}>
            つぶやきを追加
          </Typography>
          <Box display="flex" alignItems="center">
            <TextField
              variant="outlined"
              placeholder="自分のつぶやき"
              size="small"
              fullWidth
              value={note}
              onChange={(e) => setNote(e.target.value)}
            />
            <Button
              color="success"
              sx={{ borderRadius: 1, padding: "6px 12px" }}
              onClick={handleUpdateNote}
            >
              OK
            </Button>
          </Box>
        </Box>
      </Box>
    </Card>
  );
}
