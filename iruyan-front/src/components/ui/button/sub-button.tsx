import Button from "@mui/joy/Button";
import { SxProps } from "@mui/system"; // SxProps型をインポート

type Props = {
  title: string;
  link?: string;
  size: "xs" | "sm" | "md" | "lg" | "xl";
  onClick?: () => void;
  disabled?: boolean;
  sx?: SxProps; // sxプロパティを追加
};

export default function SubButton(Props: Props) {
  return (
    <Button
      component="a"
      href={Props.link}
      size={Props.size}
      onClick={Props.onClick}
      disabled={Props.disabled}
      sx={{
        color: "#7A8764",
        backgroundColor: "#FFFFFF",
        border: "5px solid #7A8764",
        borderRadius: "8px",
        "&:hover": {
          backgroundColor: "#7A8764",
          color: "#F7F4ED",
        },
        "&:active": {
          backgroundColor: "#353A2B",
          color: "#F7F4ED",
          border: "5px solid #353A2B",
        },
        ...Props.sx, // Props.sx を展開して既存のスタイルにマージ
      }}
    >
      {Props.title}
    </Button>
  );
}
