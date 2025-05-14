import { Box, Button, Modal, Typography } from '@mui/joy';
import { useEffect, useRef, useState } from 'react';
import Webcam from 'react-webcam';
import useUserStore from '@/stores/user-store';

export type Props = {
  isCameraOpen: boolean;
  setIsCameraOpen: (value: boolean) => void;
};

export default function CameraModal(Props: Props) {
  const setAvatarUrl = useUserStore((state) => state.setAvatarUrl);
  const webcamRef = useRef<Webcam>(null);

  const handleCapture = () => {
    if (webcamRef.current) {
      const imageSrc = webcamRef.current.getScreenshot();
      if (imageSrc) {
        setAvatarUrl(imageSrc);
        Props.setIsCameraOpen(false);
      }
    }
  };

  return (
    <Modal
      open={Props.isCameraOpen}
      onClose={() => Props.setIsCameraOpen(false)}
      aria-labelledby="camera-modal-title"
      aria-describedby="camera-modal-description"
    >
      <Box
        sx={{
          width: { xs: '90%', sm: 400 },
          bgcolor: 'background.paper',
          borderRadius: 2,
          boxShadow: 24,
          p: 4,
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
        }}
      >
        <Typography id="camera-modal-title" level="body-md" component="h2" sx={{ mb: 2 }}>
          カメラで写真を撮る
        </Typography>
        <Webcam
          audio={false}
          ref={webcamRef}
          screenshotFormat="image/jpeg"
          videoConstraints={{
            facingMode: 'user',
          }}
          style={{ width: '100%', borderRadius: '8px' }}
        />
        <Box sx={{ display: 'flex', gap: 2, mt: 2 }}>
          <Button color="primary" onClick={handleCapture}>
            撮影
          </Button>
          <Button variant="outlined" color="warning" onClick={() => Props.setIsCameraOpen(false)}>
            キャンセル
          </Button>
        </Box>
      </Box>
    </Modal>
  );
}
