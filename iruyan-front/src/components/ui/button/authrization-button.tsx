import { Button } from "@mui/material";

type Props = {
  title: string;
  link: string;
};

export default function AuthorizationButton(Props: Props) {
  return (
    <Button
      variant="contained"
      href={Props.link}
      sx={{backgroundColor: "#7A8764", fontWeight: "bold", p:2, height: "5rem", width: "31rem"}}
    >
      {Props.title}
    </Button>

  );
}