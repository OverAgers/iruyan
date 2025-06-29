"use client";

import React, { useState, useRef, useEffect } from "react";
import { Box, Typography, Button, Modal } from "@mui/joy";
import TextField from "@mui/material/TextField";
import CameraIcon from "@mui/icons-material/CameraAlt";
import MainButton from "@/components/ui/button/main-button";
import SubButton from "@/components/ui/button/sub-button";
import useUserStore from "@/stores/user-store";
import Webcam from "react-webcam";
import useGetRoomList from "@/features/lobby/api/get-room-list";
import usePostRoomEnterRequest from "@/features/lobby/api/post-room-enter";
// import usePostLogoutRequest from "@/features/lobby/api/post-user-logout";

export default function Lobby() {
  const currentUser = useUserStore((state) => state.currentUser);
  const { setTask, setNote, setAvatarUrl, clearUser } = useUserStore();

  const roomList = useGetRoomList();
  const roomEntry = usePostRoomEnterRequest();

  const [taskInput, setTaskInput] = useState<string>(currentUser?.task || "");
  const [noteInput, setNoteInput] = useState<string>(currentUser?.note || "");
  const [isCameraOpen, setIsCameraOpen] = useState<boolean>(false);
  const [selectedRoomId, setSelectedRoomId] = useState<string | null>(null);

  const webcamRef = useRef<Webcam>(null);

  useEffect(() => {
    const userInfo = localStorage.getItem("user-store");
    if (!userInfo) {
      console.log("ユーザー情報が存在しないため、ログインページにリダイレクトします。");
      window.location.href = "/login";
      return;
    }
  }, []);

  if (!currentUser) {
    return <Typography>ログインしてください。</Typography>;
  }

  const handleUpdateUser = async () => {
    if (!selectedRoomId) {
      alert("入室する部屋を選択してください。");
      return;
    }

    try {
      setTask(taskInput);
      setNote(noteInput);

      await roomEntry.entry({
        iruyanId: currentUser.iruyanId,
        roomId: selectedRoomId,
      });

      window.location.href = `/rooms/:${selectedRoomId}`;
    } catch (error) {
      console.error("部屋へのエントリーに失敗しました:", error);
      alert("入室に失敗しました");
    }
  };

  const handleLogout = async () => {
    localStorage.removeItem("user-store");
    clearUser();
    window.location.href = "/login";
  };

  const handleCapture = async () => {
    if (webcamRef.current) {
      const imageSrc = webcamRef.current.getScreenshot();
      if (imageSrc) {
        setAvatarUrl(imageSrc);
        setIsCameraOpen(false);
      }
    }
  };

  return (
    <Box
      sx={{
        position: "relative",
        display: "flex",
        alignItems: "flex-end",
        justifyContent: "center",
        height: "100vh",
        backgroundImage: "url('/bg-image/bg_lobby.jpg')",
        backgroundSize: "cover",
      }}
    >
      <Box sx={({
        position: "absolute",
        top: 0,
        right: 0,
        display: "flex",
        gap: "16px",
        padding: "32px"
      })}>
        <SubButton
          title="ログアウト"
          size="lg"
          onClick={handleLogout}
        />
        <SubButton
          title="来店記録"
          size="lg"
          onClick={() => (window.location.href = "/user")}
        />
      </Box>

      <Box sx={({
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        paddingBottom: "40px",
      })}>
        <Typography level="h4" sx={{ mb: 1, fontSize: "40px" }}>
          ご案内用紙
        </Typography>

        {/* 部屋一覧表示 */}
        {roomList.data && roomList.data.rooms.length > 0 && (
          <Box
            sx={{
              mt: 4,
              p: 2,
              bgcolor: "white",
              borderRadius: 2,
              boxShadow: 2,
              width: "80%",
            }}
          >
            <Typography sx={{ mb: 2 }}>現在の部屋一覧</Typography>
            {roomList.data.rooms.map((room) => (
              <Box
                key={room.roomId}
                sx={{
                  p: 1,
                  border: "1px solid #ccc",
                  borderRadius: 1,
                  mb: 1,
                  cursor: "pointer",
                  backgroundColor: selectedRoomId === room.roomId ? "#e0f7fa" : "#ffffff",
                  "&:hover": { backgroundColor: "#f0f0f0" },
                }}
                onClick={() => setSelectedRoomId(room.roomId)}
              >
                <Typography>{room.roomName}</Typography>
              </Box>
            ))}
          </Box>
        )}

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
          {/* アバターとタスク・メモ欄 */}
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
                "&:hover": { bgcolor: "#f0f0f0" },
              }}
              onClick={() => setIsCameraOpen(true)}
            >
              {currentUser?.avatarUrl ? (
                <img
                  src={currentUser.avatarUrl}
                  alt="Avatar"
                  style={{
                    width: "100%",
                    height: "100%",
                    borderRadius: "50%",
                    objectFit: "cover",
                  }}
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
            <Box sx={{ display: "flex", flexDirection: "column", width: "30vw" }}>
              <Typography level="h4" sx={{ mb: 1 }}>
                作業内容
              </Typography>
              <TextField
                placeholder="勉強"
                fullWidth
                value={taskInput}
                onChange={(e) => setTaskInput(e.target.value)}
                sx={{ bgcolor: "white" }}
              />
            </Box>
            <Box sx={{ display: "flex", flexDirection: "column", width: "30vw" }}>
              <Typography level="h4" sx={{ mb: 1 }}>
                今日のやる気 / つぶやき
              </Typography>
              <TextField
                placeholder="課題やばい、、よ"
                fullWidth
                value={noteInput}
                onChange={(e) => setNoteInput(e.target.value)}
                sx={{ bgcolor: "white" }}
              />
            </Box>
          </Box>
        </Box>

        <MainButton
          title="席に着く"
          type="button"
          maxWidth="240px"
          width="50%"
          component={"div"}
          onClick={handleUpdateUser}
          disabled={roomList.isLoading}
        />
      </Box>

      {/* カメラ撮影モーダル */}
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
          <Typography id="camera-modal-title" level="body-md" component="h2" sx={{ mb: 2 }}>
            カメラで写真を撮る
          </Typography>
          <Webcam
            audio={false}
            ref={webcamRef}
            screenshotFormat="image/jpeg"
            videoConstraints={{ facingMode: "user" }}
            style={{ width: "100%", borderRadius: "8px" }}
          />
          <Box sx={{ display: "flex", gap: 2, mt: 2 }}>
            <Button color="primary" onClick={handleCapture}>
              撮影
            </Button>
            <Button variant="outlined" color="warning" onClick={() => setIsCameraOpen(false)}>
              キャンセル
            </Button>
          </Box>
        </Box>
      </Modal>
    </Box >
  );
}
