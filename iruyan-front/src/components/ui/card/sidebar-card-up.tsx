import { Avatar, Box, Button, List, ListItem, Typography } from "@mui/joy";
import SubButton from "../button/sub-button";
import { useEffect, useState } from "react";
import { TextField } from "@mui/material";
import useUserStore from "@/stores/user-store";
import { FormatTime } from "@/utils/format-time";

export default function SidebarCardUp() {
  const { currentUser, setUser } = useUserStore();

  const [name, setName] = useState(currentUser?.name || "名前");
  const [avatarUrl, setAvatarUrl] = useState(currentUser?.avatarUrl || "");
  const [entryTime, setEntryTime] = useState<Date | null>(null);
  const [cumulativeTime, setCumulativeTime] = useState<number>(0);
  const [consecutiveDays, setConsecutiveDays] = useState<number>(0);
  const [task, setTask] = useState(currentUser?.task || "");
  const [note, setNote] = useState(currentUser?.note || "");

  useEffect(() => {
    if (!currentUser) return;
    setName(currentUser.name);
    setAvatarUrl(currentUser.avatarUrl || "");
    setTask(currentUser.task || "");
    setNote(currentUser.note || "");

    // 入店時間を設定（ここでは初回マウント時の時間を使用）
    if (!entryTime) {
      const now = new Date();
      setEntryTime(now);
    }

    // サーバーから累計集中時間と連続入室日数を取得
    fetchUserStats(currentUser.iruyanId)
      .then((stats) => {
        setCumulativeTime(stats.cumulativeTime);
        setConsecutiveDays(stats.consecutiveDays);
      })
      .catch((error) => {
        console.error("ユーザーステータスの取得に失敗しました:", error);
      });
  }, [currentUser]);

  const handleUpdateTask = () => {
    if (!currentUser) return;
    setUser({ ...currentUser, task });
  };

  const handleUpdateNote = () => {
    if (!currentUser) return;
    setUser({ ...currentUser, note });
  };

  // サーバーからユーザーステータスを取得する関数（ダミー実装）
  const fetchUserStats = async (iruyanID: string) => {
    console.log("ユーザーステータスを取得します:", iruyanID);
    // 実際のAPIコールに置き換えてください
    // 例:
    // const response = await fetch(`/api/users/${userId}/stats`);
    // const data = await response.json();
    // return data;

    // ダミーデータを返す
    return new Promise<{ cumulativeTime: number; consecutiveDays: number }>(
      (resolve) => {
        setTimeout(() => {
          resolve({
            cumulativeTime: 3600, // 例: 1時間（秒単位）
            consecutiveDays: 3, // 例: 3日連続
          });
        }, 500);
      }
    );
  };

  return (
    <Box
      sx={{
        p: 2,
        borderRadius: "8px",
        backgroundColor: "#f3f0e9",
        mb: 3,
      }}
    >
      <Box>
        <Box
          justifyContent={"space-between"}
          display={"flex"}
          alignItems={"center"}
        >
          <Box display={"flex"} alignItems={"center"}>
            <Avatar src={avatarUrl} />
            <Typography ml={2}>{name}</Typography>
          </Box>
          <SubButton title="来店記録" link="/user" size="lg" />
        </Box>
        <List marker="disc" size="sm">
          <ListItem>
            入店時間: {entryTime ? entryTime.toLocaleTimeString() : "未設定"}
          </ListItem>
          <ListItem>累計集中時間: {FormatTime(cumulativeTime)}</ListItem>
          <ListItem>連続入室日数: {consecutiveDays}日</ListItem>
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
    </Box>
  );
}