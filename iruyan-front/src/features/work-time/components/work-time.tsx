import { Box, Typography } from "@mui/joy";
import SubButton from "@/components/ui/button/sub-button";

export default function Timer() {
  const workTimeLeft = "30:56";
  const restTimeLeft = "5:30";

  return (
    <>
      <Box display="flex" flexDirection="column" alignItems="center" mb={3}>
        <Box display="flex" justifyContent="space-between" width="100%" mb={1}>
          <Typography fontSize="16px">作業時間</Typography>
          <Typography fontSize="32px" fontWeight="bold" color="primary">
            {workTimeLeft}
          </Typography>
        </Box>
        <Box display="flex" justifyContent="space-between" width="100%">
          <Typography fontSize="16px">休憩時間</Typography>
          <Typography fontSize="24px" fontWeight="bold" color="primary">
            {restTimeLeft}
          </Typography>
        </Box>
      </Box>
      <Box display="flex" justifyContent="space-around" mb={2}>
        <SubButton title={"鬼集中 🔥"} size={"lg"} />
        <SubButton title={"休憩 😴"} size={"lg"} />
      </Box>
      <Box display="flex" justifyContent="center">
        <SubButton title={"席を離れる 👋"} size={"lg"} />
      </Box>
    </>
  );
}
