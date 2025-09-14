"use client";

import React from "react";
import {
  Box,
  Typography,
  Avatar,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from "@mui/material";
import SubButton from "@/components/ui/button/sub-button";
import { useRouter } from "next/navigation";
import DynamicBarChart from "@/components/charts/dynamic-bar-chart";

function UserPage() {
  const router = useRouter();
  const chartData = [
    { day: "月", hours: 4 },
    { day: "火", hours: 3 },
    { day: "水", hours: 5 },
    { day: "木", hours: 7 },
    { day: "金", hours: 3 },
    { day: "土", hours: 1 },
    { day: "日", hours: 0 },
  ];

  return (
    <Box
      sx={{
        position: "relative",
        bgcolor: "#F7F4ED",
        backgroundImage: "url('/bg-image/bg_user.jpg')",
        backgroundSize: "contain",
      }}
    >
      <SubButton
        title="ノートを閉じる"
        size="lg"
        onClick={() => router.push("/lobby")}
        sx={{
          position: "absolute",
          top: "37px",
          left: "63px",
        }}
      />
      <Box
        display={"flex"}
        justifyContent={"center"}
        alignItems={"flex-start"}
        pt={20}
        sx={{
          gap: "7%",
        }}
      >
        <Box
          sx={{
            textAlign: "center",
            display: "flex",
            flexDirection: "column",
            alignItems: "flex-start",
            gap: "44px",
            width: "40%",
          }}
        >
          <Box
            sx={{
              textAlign: "center",
              display: "flex",
              alignItems: "center",
              justifyContent: "center",
              gap: "20px",
            }}
          >
            <Avatar
              alt="User Avatar"
              src="/path/to/profile-image.jpg"
              sx={{ width: 160, height: 160, mb: 2 }}
            />
            <Box
              sx={{
                textAlign: "center",
                display: "flex",
                flexDirection: "column",
                alignItems: "flex-start",
                gap: "10px",
              }}
            >
              <Typography variant="h5" sx={{ fontSize: "32px" }}>
                ほしょ
              </Typography>
              <Box
                sx={({
                  display: "flex",
                  alignItems: "baseline",
                  gap: "16px"
                })}
              >
                <Typography sx={({ fontSize: "16px" })}>
                  居る連チャン
                </Typography>
                <Typography sx={{ fontSize: "20px" }}>
                  5Days
                </Typography>
              </Box>
              <Box
                sx={({
                  display: "flex",
                  alignItems: "baseline",
                  gap: "16px"
                })}
              >
                <Typography sx={({ fontSize: "16px" })}>
                  累計集中時間
                </Typography>
                <Typography sx={{ fontSize: "20px" }}>
                  120h
                </Typography>
              </Box>
            </Box>
          </Box>
          <Box
            sx={{
              textAlign: "center",
              display: "flex",
              alignItems: "flex-start",
              justifyContent: "center",
              gap: "53px",
            }}
          >
            <Box
              sx={{
                textAlign: "center",
                display: "flex",
                flexDirection: "column",
                alignItems: "flex-start",
                width: "fit-content",
                gap: "18px",
              }}
            >
              <Typography align="center" sx={{ fontSize: "16px" }}>
                一緒に居た常連さん
              </Typography>
              <Box
                sx={{
                  textAlign: "center",
                  display: "flex",
                  flexDirection: "column",
                  alignItems: "flex-start",
                  width: "100%",
                  gap: "18px",
                }}
              >
                {[...Array(5)].map((_, index) => (
                  <Box
                    key={index}
                    sx={{
                      display: "flex",
                      alignItems: "center",
                      mt: 1,
                      width: "100%",
                      gap: "1vw",
                    }}
                  >
                    <Box sx={{ position: "relative", display: "inline-block" }}>
                      <Box
                        sx={{
                          position: "absolute",
                          top: -8,
                          left: -8,
                          backgroundColor: "#d4a373",
                          color: "#fff",
                          width: "31px",
                          height: "31px",
                          borderRadius: "50%",
                          display: "flex",
                          alignItems: "center",
                          justifyContent: "center",
                          fontSize: "12px",
                          fontWeight: "bold",
                          zIndex: 1,
                        }}
                      >
                        {index + 1}
                      </Box>
                      <Avatar
                        alt="Friend Avatar"
                        src="/path/to/profile-image.jpg"
                        sx={{ width: 56, height: 56 }}
                      />
                    </Box>
                    <Box
                      sx={{
                        display: "flex",
                        justifyContent: "space-between",
                        gap: "3vw",
                      }}
                    >
                      <Typography sx={{ fontSize: "20px" }}>ほしょ</Typography>
                      <Typography sx={{ fontSize: "20px" }}>3.4h</Typography>
                    </Box>
                  </Box>
                ))}
              </Box>
            </Box>
            <Box
              sx={{
                textAlign: "center",
                display: "flex",
                flexDirection: "column",
                alignItems: "flex-start",
                width: "fit-content",
                gap: "18px",
              }}
            >
              <Typography align="center" sx={{ fontSize: "16px" }}>
                今週の集中ランキング
              </Typography>
              <Box
                sx={{
                  textAlign: "center",
                  display: "flex",
                  flexDirection: "column",
                  alignItems: "flex-start",
                  width: "100%",
                  gap: "18px",
                }}
              >
                {[...Array(5)].map((_, index) => (
                  <Box
                    key={index}
                    sx={{
                      display: "flex",
                      alignItems: "center",
                      mt: 1,
                      width: "100%",
                      gap: "1vw",
                    }}
                  >
                    <Box sx={{ position: "relative", display: "inline-block" }}>
                      <Box
                        sx={{
                          position: "absolute",
                          top: -8,
                          left: -8,
                          backgroundColor: "#d4a373",
                          color: "#fff",
                          width: "31px",
                          height: "31px",
                          borderRadius: "50%",
                          display: "flex",
                          alignItems: "center",
                          justifyContent: "center",
                          fontSize: "12px",
                          fontWeight: "bold",
                          zIndex: 1,
                        }}
                      >
                        {index + 1}
                      </Box>
                      <Avatar
                        alt="Friend Avatar"
                        src="/path/to/profile-image.jpg"
                        sx={{ width: 56, height: 56 }}
                      />
                    </Box>
                    <Box
                      sx={{
                        display: "flex",
                        justifyContent: "space-between",
                        gap: "3vw",
                      }}
                    >
                      <Typography sx={{ fontSize: "20px" }}>ほしょ</Typography>
                      <Typography sx={{ fontSize: "20px" }}>3.4h</Typography>
                    </Box>
                  </Box>
                ))}
              </Box>
            </Box>
          </Box>
        </Box>
        <Box
          sx={{
            display: "flex",
            flexDirection: "column",
            alignItems: "flex-start",
            width: "35%",
          }}
        >
          <Typography align="center" sx={{ fontSize: "25px" }}>
            ご来店記録
          </Typography>
          <TableContainer sx={{ minWidth: "40vw" }}>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell>日付</TableCell>
                  <TableCell>作業時間</TableCell>
                  <TableCell>休憩時間</TableCell>
                  <TableCell>作業内容</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {[...Array(5)].map((_, index) => (
                  <TableRow key={index}>
                    <TableCell>1/1(月)</TableCell>
                    <TableCell>13:00</TableCell>
                    <TableCell>2:00</TableCell>
                    <TableCell>勉強</TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
          <Box sx={{ width: 400 }}>
            <DynamicBarChart
              data={chartData}
              width={400}
              height={400}
              color="#D3AE6F"
            />
          </Box>
        </Box>
      </Box>
    </Box >
  );
}

export default UserPage;
