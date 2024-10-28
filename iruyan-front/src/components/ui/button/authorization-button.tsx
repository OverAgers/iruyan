import { Button } from "@mui/material";

type Props = {
  title: string;
  type : "submit" | "button" | "reset" | undefined;
};

export default function AuthorizationButton(Props: Props) {
  return (
    <Button
      variant="contained"
      type= {Props.type}
      sx={{backgroundColor: "#7A8764", fontWeight: "bold", p:2, height: "5rem", width: "31rem"}}
    >
      {Props.title}
    </Button>

  );
}