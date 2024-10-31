import { Box, Typography } from "@mui/joy";
import SubButton from "@/components/ui/button/sub-button";
import { useEffect } from "react";
import { FormatTime } from "@/utils/format-time";
import useUserStore from "@/stores/user-store";

export default function WorkTime() {
  const {
    currentUser,
    setStatus,
    incrementWorkTime,
    incrementRestTime,
    setStartTime,
    resetTimes,
  } = useUserStore();

  useEffect(() => {
    if (!currentUser) return;

    const { status, startTime } = currentUser;

    if (status !== "idle" && startTime) {
      const now = Date.now();
      const elapsed = Math.floor((now - startTime) / 1000);
      if (status === "working") {
        incrementWorkTime(elapsed);
      } else if (status === "resting") {
        incrementRestTime(elapsed);
      }
      setStartTime(now);
    }
  }, []);

  useEffect(() => {
    if (!currentUser) return;

    let interval: NodeJS.Timeout | null = null;

    if (currentUser.status !== "idle") {
      interval = setInterval(() => {
        if (currentUser.status === "working") {
          incrementWorkTime(1);
        } else if (currentUser.status === "resting") {
          incrementRestTime(1);
        }
      }, 1000);
    }

    return () => {
      if (interval) clearInterval(interval);
    };
  }, [currentUser?.status]);

  const handleStart = (newStatus: "working" | "resting") => {
    setStatus(newStatus);
    setStartTime(Date.now());
  };

  const handleStop = () => {
    setStatus("idle");
    setStartTime(0);
  };

  if (!currentUser) {
    return <div>ログインしてください。</div>;
  }
  return (
    <>
      <Box display="flex" flexDirection="column" alignItems="center" mb={3}>
        <Box
          display="flex"
          justifyContent="space-between"
          width="100%"
          mb={1}
          alignItems={"center"}
        >
          <Typography level="body-md">作業時間</Typography>
          <Typography level="h1">{FormatTime(currentUser.workTime)}</Typography>
        </Box>
        <Box
          display="flex"
          justifyContent="space-between"
          width="100%"
          alignItems={"center"}
        >
          <Typography level="body-md">休憩時間</Typography>
          <Typography level="h2">{FormatTime(currentUser.restTime)}</Typography>
        </Box>
      </Box>
      <Box display="flex" justifyContent="space-around" mb={2}>
        <SubButton
          title={"鬼集中 🔥"}
          size={"lg"}
          onClick={() => handleStart("working")}
        />
        <SubButton
          title={"休憩 😴"}
          size={"lg"}
          onClick={() => handleStart("resting")}
        />
      </Box>
      <Box display="flex" justifyContent="center">
        <SubButton
          title={"席を離れる 👋"}
          size={"lg"}
          onClick={handleStop}
        />
      </Box>
    </>
  );
}
