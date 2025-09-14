import { Typography } from "@mui/joy";
import { Grid2 } from "@mui/material";
import Sidebar from "@/features/sidebar/components/sidebar";
import OptimizedBackground from "@/components/ui/background/optimized-background";

type Props = {
  children: React.ReactNode;
};

export default function RoomLayout({ children }: Props) {
  return (
    <OptimizedBackground
      src="/bg-image/bg_room.jpg"
      alt="Room background"
      priority={true}
      sx={{
        height: "100vh",
        bgcolor: "#3C2800",
      }}
    >
      <Grid2
        container
        sx={{
          height: "100vh",
        }}
        justifyContent={"space-between"}
      >
      <Grid2 container size={7}>
        <Grid2 container justifyContent={"space-between"} width={"100%"}>
          <Grid2 padding={4}>
            <Typography level="title-md" textColor={"#F7F4ED"}>
              IRUYAN
            </Typography>
            <Typography level="h3" textColor={"#F7F4ED"}>
              居る家ん
            </Typography>
          </Grid2>
          <Grid2 padding={4} size={9}>
              {/* <Typography>通知</Typography> */}
          </Grid2>
        </Grid2>
        <Grid2>{children}</Grid2>
      </Grid2>
      <Grid2 size={4}>
        <Sidebar></Sidebar>
      </Grid2>
      </Grid2>
    </OptimizedBackground>
  );
}
