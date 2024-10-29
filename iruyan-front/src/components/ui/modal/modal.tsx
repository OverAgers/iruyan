import * as React from 'react';
import Button from '@mui/joy/Button';
import Modal from '@mui/joy/Modal';
import Typography from '@mui/joy/Typography';
import Sheet from '@mui/joy/Sheet';
import ModalClose from '@mui/joy/ModalClose';
import MainButton from "@/components/ui/button/main-button";

export default function BasicModal() {
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
        sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center' }}
      >
        <Sheet
          variant="outlined"
          sx={{ 
            width: '50%',
            maxWidth: 641, 
            borderRadius: '16px', 
            align: "center",
            p: 2, 
            boxShadow: 'lg',
            border: '13px solid #D3AE6F', // 枠線を指定
            display: 'flex',
            flexDirection: 'column',  // 縦に要素を並べる
            alignItems: 'center',     // 横方向の中央揃え
            paddingTop: 4,
            paddingBottom: 4
          }}
        >
          <Typography
            component="h2"
            id="modal-title"
            level="h4"
            textColor="inherit"
            fontSize={36}
            sx={{
              fontWeight: 'lg', mb: 1, color: '#3C2800',
            }}
          >
            お疲れ様でした！
          </Typography>
          <Typography id="modal-desc" textColor="text.tertiary"
            fontSize={20}
            sx={{
              mb: 1, color: '#3C2800', paddingTop: 4, paddingBottom: 4
            }}>
            1時間30分の作業、お見事です！<br />
            またお待ちしています。
          </Typography>
          {/* 下部に配置されたModalClose */}
          <MainButton title="受付に戻る" type="button" maxWidth = "342px" width="50%"/>
        </Sheet>
      </Modal>
    </React.Fragment>
  );
}
