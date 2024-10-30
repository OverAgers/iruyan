import SidebarCardDown from "@/components/ui/card/sidebar-card-down";
import SidebarCardUp from "@/components/ui/card/sidebar-card-up";
import { Typography } from "@mui/joy";
import { Box } from "@mui/material";
import { useEffect, useState } from "react";

export default function Sidebar() {
  const [currentTime, setCurrentTime] = useState(new Date());

  useEffect(() => {
    const timer = setInterval(() => {
      setCurrentTime(new Date());
    }, 1000);
    return () => clearInterval(timer);
  }, []);

  return (
    <Box width={"385px"} justifySelf={"flex-end"}>
      <SidebarCardUp />
      <SidebarCardDown />
      <Box>
        <Typography level="h2" mb={2}>
          現在時刻
        </Typography>
        <Typography fontSize={"128px"} lineHeight={1}>
          {currentTime.toLocaleTimeString([], {
            hour: "2-digit",
            minute: "2-digit",
          })}
        </Typography>
      </Box>
    </Box>
  );
}
