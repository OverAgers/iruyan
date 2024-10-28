import { Button } from "@mui/material";

type Props = {
  title: string;
  link: string;
};

export default function ChangeAuthorizationButton(Props: Props) {
  return (
    <Button
      variant="contained"
      href={Props.link}
      sx={{
        color: "#7A8764",
        backgroundColor: "#F7F4ED",
        fontWeight: "bold",
        p: 2,
        height: "4rem",
        width: "12rem",
        border: "5px solid #7A8764",
        borderRadius: "8px",
      }}
    >
      {Props.title}
    </Button>
  );
}
