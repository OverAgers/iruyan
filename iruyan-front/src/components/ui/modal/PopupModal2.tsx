import * as React from "react";
import Button from "@mui/joy/Button";
import Modal from "@mui/joy/Modal";
import Typography from "@mui/joy/Typography";
import Sheet from "@mui/joy/Sheet";
import ModalClose from "@mui/joy/ModalClose";
import MainButton from "@/components/ui/button/main-button";

export default function PopupModal2() {
  const [open, setOpen] = React.useState<boolean>(false);
  return (
    <React.Fragment>
      <Button variant="outlined" color="neutral" onClick={() => setOpen(true)}>
        Open modal
      </Button>
      <Modal
        aria-labelledby="modal-title"
        aria-describedby="modal-desc"
        open={open}
        onClose={() => setOpen(false)}
        sx={{ display: "flex", justifyContent: "center", alignItems: "center" }}
      >
        <Sheet
          variant="outlined"
          sx={{
            width: "50%",
            maxWidth: 641,
            borderRadius: "16px",
            align: "center",
            p: 2,
            boxShadow: "lg",
            border: "13px solid #D3AE6F", // 枠線を指定
            display: "flex",
            flexDirection: "column", // 縦に要素を並べる
            alignItems: "center", // 横方向の中央揃え
            paddingTop: 4,
            paddingBottom: 4,
          }}
        >

          <Typography
            component="h2"
            id="modal-title"
            level="h4"
            textColor="inherit"
            fontSize={36}
            sx={{
              fontWeight: "lg",
              mb: 1,
              color: "#3C2800",
              textAlign: "center",
            }}
          >
            作業を終えて、<br />
            受付に戻りますか？
          </Typography>
          <Typography
            id="modal-desc"
            textColor="text.tertiary"
            fontSize={20}
            sx={{
              mb: 1,
              color: "#3C2800",
              paddingTop: 4,
              paddingBottom: 4,
            }}
          >
            1時間30分の作業を記録します。
          </Typography>
          <MainButton
            title="受付に戻る"
            type="button"
            maxWidth="342px"
            width="50%"
            component="a"
          />
          {/* 下部に配置されたModalClose */}
          <Button
            variant="plain"
            color="neutral"
            sx={{ fontSize: 18, color: "#3C2800", paddingTop: "14px", paddingButtom: "0px"}}
            onClick={() => setOpen(false)}
          >
              × 作業を続ける
          </Button>
        </Sheet>
      </Modal>
    </React.Fragment>
  );
}