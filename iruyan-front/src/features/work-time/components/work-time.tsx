import { Box, Typography, LinearProgress } from "@mui/joy";
import SubButton from "@/components/ui/button/sub-button";
import { useEffect } from "react";
import { FormatTime } from "@/utils/format-time";
import useUserStore from "@/stores/user-store";
import { useTimer } from "@/hooks/timer-hooks";

export default function WorkTime() {
  const { currentUser, setStatus, incrementWorkTime, incrementRestTime } =
    useUserStore();

  const {
    timeLeft,
    isRunning,
    mode,
    progress,
    reset,
    stop,
    start,
    switchMode,
  } = useTimer({
  initialWorkTime: 1500,
  initialBreakTime: 300,
  onTick: () => {
    if (mode === "work") {
      incrementWorkTime();
    } else if (mode === "break") {
      incrementRestTime();
    }
  },
});

  useEffect(() => {
    if (!isRunning) return;

    const interval = setInterval(() => {
      if (mode === "work") {
        incrementWorkTime();
      } else if (mode === "break") {
        incrementRestTime();
      }
    }, 1000);

    return () => clearInterval(interval);
  }, [isRunning, mode]);

  if (!currentUser) {
    return <div>ログインしてください。</div>;
  }

  return (
    <>
      <Box display="flex" flexDirection="column" alignItems="center" mb={3}>
        <Typography fontSize="32px" fontWeight="bold" color="primary">
          {FormatTime(timeLeft)}
        </Typography>
        <LinearProgress
          determinate
          value={progress}
          sx={{ width: "100%", mb: 2 }}
        />
        <Box
          display="flex"
          justifyContent="space-between"
          width="100%"
          mb={1}
          alignItems="center"
        >
          <Typography level="body-md">累計作業時間</Typography>
          <Typography level="h1">{FormatTime(currentUser.workTime)}</Typography>
        </Box>
        <Box
          display="flex"
          justifyContent="space-between"
          width="100%"
          alignItems="center"
        >
          <Typography level="body-md">累計休憩時間</Typography>
          <Typography level="h2">{FormatTime(currentUser.restTime)}</Typography>
        </Box>
      </Box>
      <Box display="flex" justifyContent="space-around" mb={2}>
        <SubButton
          title={"鬼集中 🔥"}
          size={"lg"}
          onClick={() => {
            setStatus("working");
            if (mode !== "work") switchMode();
            start();
          }}
        />
        <SubButton
          title={"休憩 😴"}
          size={"lg"}
          onClick={() => {
            setStatus("resting");
            if (mode !== "break") switchMode();
            start();
          }}
        />
      </Box>
      <Box display="flex" justifyContent="center" mb={2}>
        <SubButton
          title={isRunning ? "一時停止" : "再開"}
          size={"lg"}
          onClick={() => {
            if (isRunning) {
              stop();
            } else {
              start();
            }
          }}
        />
      </Box>
      <Box display="flex" justifyContent="center">
        <SubButton
          title={"リセット"}
          size={"lg"}
          onClick={() => {
            reset();
            setStatus("idle");
          }}
        />
      </Box>
    </>
  );
}
