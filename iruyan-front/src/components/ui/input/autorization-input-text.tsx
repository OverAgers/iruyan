import { TextField, Typography } from "@mui/material";

type Props = {
  title: string
  type: string;
  placeholder: string;
}

export default function AuthInputText(Props: Props) {
  return (
    <>
      <Typography variant="h6" sx={{ fontWeight: "bold", color: "#3C2800" }}>
        {Props.title}
      </Typography>
      <TextField
        variant="outlined"
        type={Props.type}
        placeholder={Props.placeholder}
        sx={{
          borderRadius: "8px",
          width: "31rem",
          mb: 4,
          bgcolor: "#ffffff",
          "& .MuiOutlinedInput-root": {
            borderRadius: "8px",
            "& fieldset": {
              border: "none",
            },
          },
        }}
      />
    </>
  );
}