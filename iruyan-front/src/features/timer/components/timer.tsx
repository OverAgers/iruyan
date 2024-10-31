import { Box, Typography, Button, LinearProgress } from "@mui/joy";
import SubButton from "@/components/ui/button/sub-button";

export default function Timer() {
  const workTimeLeft = "15:10"; // remaining work time
  const focusMode = "集中モード"; // Focus mode label
  const totalWorkTime = "25分"; // Total work time set
  const progressValue = 60; // Assume progress percentage for demo

  return (
    <Box sx={{ p: 2, borderRadius: "8px", backgroundColor: "#f3f0e9" }}>
      <Box
        display="flex"
        justifyContent="space-between"
        alignItems="center"
        mb={2}
      >
        <Typography fontSize="32px" fontWeight="bold" color="primary">
          {workTimeLeft}
        </Typography>
        <Box textAlign="right">
          <Typography fontSize="14px">{focusMode}</Typography>
          <Typography fontSize="14px">設定時間: {totalWorkTime}</Typography>
        </Box>
      </Box>
      <LinearProgress
        determinate
        value={progressValue}
        sx={{ mb: 3 }}
      />
      <Box display="flex" justifyContent="space-around" mb={2}>
        <SubButton title={"リセット"} size={"lg"} />
        <SubButton title={"ストップ"} size={"lg"} />
      </Box>
      <Box display="flex" justifyContent="center">
        <SubButton title={"休憩モードに切り替え"} size={"lg"} />
      </Box>
    </Box>
  );
}
