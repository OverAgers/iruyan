import Typography  from "@mui/joy/Typography";
import Avatar from "@mui/joy/Avatar";
import Card from "@mui/joy/Card";
import Box from "@mui/joy/Box";

type Props = {
  id: number;
  isVacant: boolean | true;
  name: string;
  image: string | "";
  note: string | "つぶやき";
  task: React.ReactNode | null;
  onClick: () => void;
}

export default function SeatCard(Props: Props) {

  return (
    <Card
      sx={{
        padding: 0,
        margin: 1,
        width: 370,
        backgroundColor: "#f5f5f5",
        cursor: Props.isVacant ? "pointer" : "default",
      }}
      onClick={Props.onClick}
    >
      <Box display="flex" justifyContent="space-between"sx={{px:4, py:2}}>
        <Box display="flex" flexDirection="column" alignItems="center">
          <Typography level="title-lg" fontWeight="bold">
            {Props.isVacant ? "空席" : Props.name}
          </Typography>
          <Avatar
            src={Props.image}
            sx={{ width: 48, height: 48, marginTop: 1 }}
          />
        </Box>
        <Box display="flex" flexDirection="column" alignItems="flex-start">
          {!Props.isVacant && (
            <>
              <Typography
                startDecorator={Props.task}
                level="title-lg"
                fontWeight="bold"
                color="primary"
              >
                テスト勉強
              </Typography>
              <Typography level="title-lg" mt={1}>
                {Props.note}
              </Typography>
            </>
          )}
        </Box>
      </Box>
    </Card>
  );
}