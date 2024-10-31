import { Box, Typography, LinearProgress, IconButton } from "@mui/joy";
import SubButton from "@/components/ui/button/sub-button";
import AddIcon from "@mui/icons-material/Add";
import RemoveIcon from "@mui/icons-material/Remove";
import { useTimer } from "@/hooks/timer-hooks";
import { FormatTime } from "@/utils/format-time";

export default function Timer() {
  const {
    timeLeft,
    isRunning,
    mode,
    totalTime,
    progress,
    reset,
    stop,
    start,
    switchMode,
    adjustTime,
  } = useTimer({ initialWorkTime: 1500, initialBreakTime: 300 });

  return (
    <Box sx={{ p: 2, borderRadius: "8px", backgroundColor: "#f3f0e9" }}>
      <Box
        display="flex"
        justifyContent="space-between"
        alignItems="center"
        mb={2}
      >
        <Typography fontSize="32px" fontWeight="bold" color="primary">
          {FormatTime(timeLeft)}
        </Typography>
        <Box textAlign="right">
          <Typography fontSize="14px">
            {mode === "work" ? "集中モード" : "休憩モード"}
          </Typography>
          <Typography fontSize="14px">
            設定時間: {Math.floor(totalTime / 60)}分
          </Typography>
        </Box>
      </Box>
      <LinearProgress determinate value={progress} sx={{ mb: 3 }} />
      <Box display="flex" justifyContent="center" alignItems="center" mb={2}>
        <IconButton
          variant="outlined"
          color="primary"
          onClick={() => adjustTime(-300)} // 5分減らす
          sx={{ mx: 1 }}
          disabled={isRunning}
        >
          <RemoveIcon />
        </IconButton>
        <Typography fontSize="14px">時間を調整（5分単位）</Typography>
        <IconButton
          variant="outlined"
          color="primary"
          onClick={() => adjustTime(300)} // 5分増やす
          sx={{ mx: 1 }}
          disabled={isRunning}
        >
          <AddIcon />
        </IconButton>
      </Box>
      <Box display="flex" justifyContent="space-around" mb={2}>
        <SubButton title={"リセット"} size={"lg"} onClick={reset} />
        {isRunning ? (
          <SubButton title={"ストップ"} size={"lg"} onClick={stop} />
        ) : (
          <SubButton title={"スタート"} size={"lg"} onClick={start} />
        )}
      </Box>
      <Box display="flex" justifyContent="center">
        <SubButton
          title={
            mode === "work" ? "休憩モードに切り替え" : "集中モードに切り替え"
          }
          size={"lg"}
          onClick={switchMode}
        />
      </Box>
    </Box>
  );
}
