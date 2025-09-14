import { usePathname } from "next/navigation";
import SubButton from "@/components/ui/button/sub-button";
import { Typography, Grid2 } from "@mui/material";
import OptimizedBackground from "@/components/ui/background/optimized-background";

type Props = {
  children: React.ReactNode;
};

export default function AuthLayout({ children }: Props) {
  const pathname = usePathname();
  const isLoginPage = pathname === "/login";
  const changeButtonTitle = isLoginPage ? "新規登録" : "入店する";
  const changeLink = isLoginPage ? "/register" : "/login";
  return (
    <OptimizedBackground
      src="/bg-image/bg_login.jpg"
      alt="Login background"
      priority={true}
      sx={{
        bgcolor: "#F7F4ED",
        height: "100vh",
      }}
    >
      <Grid2
        container
        justifyContent={"center"}
        height={"100vh"}
        alignItems={"flex-start"}
      >
      <Grid2 container padding={6} size={12} justifyContent={"flex-end"}>
        <SubButton title={changeButtonTitle} link={changeLink} size="lg" />
      </Grid2>
      <Grid2 container size={5}>
        <Grid2
          container
          flexDirection={"column"}
          justifyContent={"center"}
          size={12}
        >
          <Typography variant="subtitle2" textAlign={"center"}>
            IRUYAN
          </Typography>
          <Typography variant="h3" textAlign={"center"} mb={2}>
            居る家ん
          </Typography>
          {children}
        </Grid2>
      </Grid2>
      </Grid2>
    </OptimizedBackground>
  );
}
