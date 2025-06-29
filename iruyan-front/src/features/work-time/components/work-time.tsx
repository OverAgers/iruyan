import { Box, Typography } from "@mui/joy";
import SubButton from "@/components/ui/button/sub-button";
import { useEffect } from "react";
import { formatTime } from "@/utils/time-utils";
import useUserStore from "@/stores/user-store";
import { useTimer } from "@/hooks/timer-hooks";

export default function WorkTime() {
  const { currentUser, setStatus, incrementWorkTime, incrementRestTime, clearUser } =
    useUserStore();

  const {
    isRunning,
    mode,
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

    const handleLogout = async () => {
      // try {
      console.log(currentUser);
      // await logout.logout({ iruyanID: currentUser.iruyanID });
      localStorage.removeItem("user-store");
      clearUser();
      window.location.href = "/login";
      setStatus("idle");
      stop();
      // } catch (error) {
      // console.error("ログアウトに失敗しました:", error);
      // }
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
          alignItems="center"
        >
          <Typography level="body-md">累計作業時間</Typography>
          <Typography level="h1">{formatTime(currentUser.workTime)}</Typography>
        </Box>
        <Box
          display="flex"
          justifyContent="space-between"
          width="100%"
          alignItems="center"
        >
          <Typography level="body-md">累計休憩時間</Typography>
          <Typography level="h2">{formatTime(currentUser.restTime)}</Typography>
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
      <Box textAlign={"center"}>
        <SubButton
          title={"席を離れる 👋"}
          size={"lg"}
          onClick={handleLogout}
        />
      </Box>
    </>
  );
}
