import  Button  from "@mui/joy/Button";

type Props = {
  title: string;
  type: "submit" | "button" | "reset" | undefined;
  fullWidth?: boolean;
  maxWidth?: string;
  width?: string;
};

export default function MainButton(Props: Props) {
  return (
    <Button
      component = "a"
      type={Props.type}
      size="lg"
      fullWidth={Props.fullWidth}
      sx={{
        backgroundColor: "#7A8764", fontWeight: "bold", p: 2, maxWidth: Props.maxWidth,
        width: Props.width,
        "&:active": {
        backgroundColor: "#353A2B",
      },
      }}
    >
      {Props.title}
    </Button>

  );
}