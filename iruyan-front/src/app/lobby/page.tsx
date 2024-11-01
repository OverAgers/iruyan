"use client";

import React, { useState, useRef } from "react";
import { Box, Typography, Button, Modal } from "@mui/joy";
import TextField from "@mui/material/TextField";
import CameraIcon from "@mui/icons-material/CameraAlt";
import MainButton from "@/components/ui/button/main-button";
import SubButton from "@/components/ui/button/sub-button";
import useUserStore from "@/stores/user-store";
import Webcam from "react-webcam";

export default function Lobby() {
  const currentUser = useUserStore((state) => state.currentUser);
  const { setTask, setNote, setAvatarUrl } = useUserStore();

  const [task, setTaskInput] = useState(currentUser?.task || "");
  const [note, setNoteInput] = useState(currentUser?.note || "");
  const [isCameraOpen, setIsCameraOpen] = useState(false);
  const webcamRef = useRef<Webcam>(null);

  const handleUpdateUser = () => {
    setTask(task);
    setNote(note);
    window.location.href = "/rooms/:aaa";
  };

  const handleCapture = () => {
    const imageSrc = webcamRef.current?.getScreenshot();
    if (imageSrc) {
      setAvatarUrl(imageSrc);
      setIsCameraOpen(false);
    }
  };

  console.log("現在のユーザー情報:", currentUser);

  if (!currentUser) {
    return <Typography>ログインしてください。</Typography>;
  }

  return (
    <Box
      sx={{
        position: "relative",
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        justifyContent: "center",
        minHeight: "100vh",
        bgcolor: "#F7F4ED",
        p: 2,
      }}
    >
      <SubButton
        title="ログアウト"
        size="lg"
        sx={{
          position: "absolute",
          top: "37px",
          right: "200px",
        }}
      />
      <SubButton
        title="来店記録"
        size="lg"
        sx={{
          position: "absolute",
          top: "37px",
          right: "41px",
        }}
      />
      <Typography level="h4" sx={{ mb: 1, fontSize: "48px" }}>
        ご案内用紙
      </Typography>
      <Typography level="h4" sx={{ mb: 3, color: "text.secondary" }}>
        {currentUser?.name}さん、ごゆっくりどうぞ
      </Typography>
      <Box
        sx={{
          display: "flex",
          justifyContent: "center",
          gap: "8%",
          width: "100%",
          maxWidth: 840,
        }}
      >
        <Box
          sx={{
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
            bgcolor: "background.default",
            flexDirection: "column",
            mr: 2,
          }}
        >
          <Typography level="h4">今日のアイコン</Typography>
          <Button
            variant="plain"
            sx={{
              width: "240px",
              height: "240px",
              borderRadius: "50%",
              bgcolor: "white",
              color: "text.primary",
              boxShadow: 1,
              "&:hover": {
                bgcolor: "#f0f0f0",
              },
            }}
            onClick={() => setIsCameraOpen(true)}
          >
            {currentUser?.avatarUrl ? (
              <img
                src={currentUser.avatarUrl}
                alt="Avatar"
                style={{ width: "100%", height: "100%", borderRadius: "50%" }}
              />
            ) : (
              <CameraIcon fontSize="large" />
            )}
          </Button>
        </Box>
        <Box
          sx={{
            display: "flex",
            alignItems: "flex-start",
            justifyContent: "center",
            flexDirection: "column",
            gap: "44px",
          }}
        >
          <Box
            sx={{
              display: "flex",
              flexDirection: "column",
              justifyContent: "flex-start",
              width: "30vw",
            }}
          >
            <Typography level="h4" sx={{ mb: 1 }}>
              ・作業内容
            </Typography>
            <TextField
              placeholder="勉強"
              fullWidth
              value={task}
              onChange={(e) => setTaskInput(e.target.value)}
              sx={{
                bgcolor: "white",
                width: "100%",
              }}
            />
          </Box>
          <Box
            sx={{
              display: "flex",
              justifyContent: "center",
              flexDirection: "column",
              width: "30vw",
            }}
          >
            <Typography level="h4" sx={{ mb: 1 }}>
              ・今日のやる気 / つぶやき
            </Typography>
            <TextField
              placeholder="課題やばい、、よ"
              fullWidth
              value={note}
              onChange={(e) => setNoteInput(e.target.value)}
              sx={{
                bgcolor: "white",
                width: "100%",
              }}
            />
          </Box>
        </Box>
      </Box>
      <Box
        sx={{
          paddingTop: 7,
        }}
      ></Box>
      <MainButton
        title="席に着く"
        type="button"
        maxWidth="240px"
        width="50%"
        component={"div"}
        onClick={handleUpdateUser}
      />
      <Modal
        open={isCameraOpen}
        onClose={() => setIsCameraOpen(false)}
        aria-labelledby="camera-modal-title"
        aria-describedby="camera-modal-description"
      >
        <Box
          sx={{
            position: "absolute",
            top: "50%",
            left: "50%",
            transform: "translate(-50%, -50%)",
            width: { xs: "90%", sm: 400 },
            bgcolor: "background.paper",
            borderRadius: 2,
            boxShadow: 24,
            p: 4,
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          <Typography
            id="camera-modal-title"
            level="body-md"
            component="h2"
            sx={{ mb: 2 }}
          >
            カメラで写真を撮る
          </Typography>
          <Webcam
            audio={false}
            ref={webcamRef}
            screenshotFormat="image/jpeg"
            videoConstraints={{
              facingMode: "user",
            }}
            style={{ width: "100%", borderRadius: "8px" }}
          />
          <Box sx={{ display: "flex", gap: 2, mt: 2 }}>
            <Button color="primary" onClick={handleCapture}>
              撮影
            </Button>
            <Button
              variant="outlined"
              color="warning"
              onClick={() => setIsCameraOpen(false)}
            >
              キャンセル
            </Button>
          </Box>
        </Box>
      </Modal>
    </Box>
  );
}
