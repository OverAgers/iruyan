import * as React from "react";
import Modal from "@mui/material/Modal";
import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";
import MainButton from "@/components/ui/button/main-button";

interface ErrorModalProps {
  open: boolean;
  onClose: () => void;
  errorMessage: string;
}

export default function ErrorModal({ open, onClose, errorMessage }: ErrorModalProps) {
  const handleClose = React.useCallback(() => {
    onClose();
  }, [onClose]);

  return (
    <Modal
      aria-labelledby="error-modal-title"
      aria-describedby="error-modal-desc"
      open={open}
      onClose={onClose}
      sx={{ display: "flex", justifyContent: "center", alignItems: "center" }}
    >
      <Box
        sx={{
          width: "90%",
          maxWidth: 500,
          borderRadius: "16px",
          p: 3,
          boxShadow: "0 10px 25px rgba(0,0,0,0.2)",
          border: "2px solid #ef4444", // 赤い枠線
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          gap: 2,
          backgroundColor: "#fef2f2", // 薄い赤背景
        }}
      >
        <Typography
          component="h2"
          id="error-modal-title"
          variant="h4"
          sx={{
            fontWeight: "bold",
            color: "#dc2626", // 赤色
            textAlign: "center",
            mb: 1,
          }}
        >
          ⚠️ エラー
        </Typography>
        
        <Typography
          id="error-modal-desc"
          variant="body1"
          sx={{
            color: "#991b1b", // 濃い赤色
            textAlign: "center",
            mb: 2,
            lineHeight: 1.5,
            whiteSpace: "pre-line", // 改行を保持
          }}
        >
          {errorMessage}
        </Typography>
        
        <MainButton
          component="button"
          title="閉じる"
          type="button"
          onClick={handleClose}
          maxWidth="200px"
        />
      </Box>
    </Modal>
  );
} 