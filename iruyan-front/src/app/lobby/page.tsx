import React from "react";
import { Box, Typography, Button } from "@mui/joy";
import TextField from "@mui/material/TextField"; // Material UIのTextFieldをインポート
import CameraIcon from "@mui/icons-material/CameraAlt";
import MainButton from "@/components/ui/button/main-button";
import SubButton from "@/components/ui/button/sub-button";

function HomeScreen() {
  return (
    <Box
      sx={{
        position: "relative", // 子要素の絶対配置の基準
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        justifyContent: "center",
        minHeight: "100vh",
        bgcolor: "#F7F4ED", // 背景色を指定
        p: 2,
      }}
    >
      {/* Right Top Button */}
      <SubButton
        title="右上ボタン"
        size="lg"
        sx={{
          position: "absolute", // 絶対位置で配置
          top: "37px", // 上から37px
          right: "41px", // 右から41px
        }}
      />

      {/* Title */}
      <Typography level="h4" sx={{ mb: 1, fontSize: "48px" }}>
        ご案内用紙
      </Typography>
      <Typography level="h4" sx={{ mb: 3, color: "text.secondary" }}>
        ほしょさん、ごゆっくりどうぞ
      </Typography>

      {/* Icon and Form Section */}
      <Box
        sx={{
          display: "flex",
          justifyContent: "center",
          gap: "8%",
          width: "100%",
          maxWidth: 840,
        }}
      >
        {/* Icon */}
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
              width: "240px", // ボタンの幅
              height: "240px", // ボタンの高さ（円形にするために幅と同じ）
              borderRadius: "50%", // 円形にするための設定
              bgcolor: "white", // 背景色を白に設定
              color: "text.primary", // アイコンの色
              boxShadow: 1, // ボタンに軽い影を追加
              "&:hover": {
                bgcolor: "#f0f0f0", // ホバー時に少し変化
              },
            }}
          >
            <CameraIcon fontSize="large" />
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
              sx={{
                bgcolor: "white", // 背景色を白に設定
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
              sx={{
                bgcolor: "white", // 背景色を白に設定
                width: "100%",
              }}
            />
          </Box>
        </Box>
      </Box>

      {/* Task Content and Motivation */}

      {/* Main Button */}
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
      />
    </Box>
  );
}

export default HomeScreen;
