import { usePathname } from "next/navigation";
import ChangeAuthorizationButton from "@/components/ui/button/change-authorization-button";
import { Box, Typography} from "@mui/material";

type Props = {
  children: React.ReactNode;
};

export default function AuthLayout({ children }: Props) {
  const pathname = usePathname();
  const isLoginPage = pathname === "/login";
  const changeButtonTitle = isLoginPage ? "新規登録" : "入店する";
  const changeLink = isLoginPage ? "/register" : "/login";
  return (
    <div
      className="flex-col size-full relative"
      style={{ background: "#F7F4ED", height: "100vh" }}
    >
      <div className="inline-block absolute top-4 right-8">
        <ChangeAuthorizationButton
          title={changeButtonTitle}
          link={changeLink}
        />
      </div>
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          justifyContent: "center",
          height: "100%",
        }}
      >
        <Typography variant="subtitle2">IRUYAN</Typography>
        <Typography variant="h3">居る家ん</Typography>
        {children}
      </Box>
    </div>
  );
}
