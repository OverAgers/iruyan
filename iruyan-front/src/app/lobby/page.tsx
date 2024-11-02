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
import usePostRoomRequest from "@/features/lobby/api/post-room";

export default function Lobby() {
  // ユーザー情報の取得
  const currentUser = useUserStore((state) => state.currentUser);
  const { setTask, setNote, setAvatarUrl, clearUser } = useUserStore();

  const roomList = useGetRoomList();

  // 部屋リストの取得
  // const {
  //   data: roomListData,
  //   error: roomListError,
  //   isLoading: roomListLoading,
  //   mutate: refetchRoomList,
  // } = useGetRoomList();

  // ステートの定義
  const [roomId, setRoomId] = useState<string>(
    "25e6a8db-b950-4d67-a46a-afb5b637575a"
  );
  const [taskInput, setTaskInput] = useState<string>(currentUser?.task || "");
  const [noteInput, setNoteInput] = useState<string>(currentUser?.note || "");
  const [isCameraOpen, setIsCameraOpen] = useState<boolean>(false);

  const webcamRef = useRef<Webcam>(null);

  // カスタムフックの取得
  // const logout = usePostLogoutRequest();
  const postRoom = usePostRoomRequest();
  const roomEntry = usePostRoomEnterRequest();

  // 初期化用 useEffect
  useEffect(() => {
    // console.log("Initial userInfo:", localStorage.getItem("user-store"));
    const userInfo = localStorage.getItem("user-store");
    if (!userInfo) {
      console.log(
        "ユーザー情報が存在しないため、ログインページにリダイレクトします。"
      );
      window.location.href = "/login";
      return;
    }
    try {
      // 部屋の作成
      postRoom.createRoom({ name: "mokumoku" });
      console.log("部屋の作成が成功しました。");
    } catch (error) {
      console.error("初期化中にエラーが発生しました:", error);
    }
  }, []); // 関数を依存配列に追加
  // ログイン確認
  if (!currentUser) {
    return <Typography>ログインしてください。</Typography>;
  }

  // ユーザー更新ハンドラー
  const handleUpdateUser = async () => {
    if (roomList.data) {
      const list = roomList.data.rooms[0].room_id;

      if (list) {
        setRoomId(list);
        console.log("aaa", roomId);
      }
      setTask(taskInput);
      setNote(noteInput);
      // const requestData: PostRoomEnterRequest = {
      //   user_id: currentUser.iruyanID,
      //   task: taskInput,
      // };
      try {
        await roomEntry.entry({ user_id: currentUser.iruyanID, room_id: list });
        if (list) {
          window.location.href = `/rooms/:${list}`;
        }
      } catch (error) {
        console.error("部屋へのエントリーに失敗しました:", error);
      }
    }
  };

  const handleLogout = async () => {
    // try {
    console.log(currentUser);
    // await logout.logout({ iruyanID: currentUser.iruyanID });
    localStorage.removeItem("user-store");
    clearUser();
    window.location.href = "/login";
    // } catch (error) {
    // console.error("ログアウトに失敗しました:", error);
    // }
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
      <SubButton
        title="ログアウト"
        size="lg"
        onClick={handleLogout}
        sx={{
          position: "absolute",
          top: "37px",
          right: "200px",
        }}
      />
      <SubButton
        title="来店記録"
        size="lg"
        onClick={() => (window.location.href = "/user")}
        sx={{
          position: "absolute",
          top: "37px",
          right: "41px",
        }}
      />
      <Box
        display={"flex"}
        flexDirection={"column"}
        alignItems={"center"}
        mb={20}
      >
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
                value={taskInput}
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
                value={noteInput}
                onChange={(e) => setNoteInput(e.target.value)}
                sx={{
                  bgcolor: "white",
                  width: "100%",
                }}
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
