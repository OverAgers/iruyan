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
  Paper,
} from "@mui/material";
import { Button } from "@mui/joy";
import MainButton from "@/components/ui/button/main-button";
import SubButton from "@/components/ui/button/sub-button";
import { BarChart } from "@mui/x-charts/BarChart";

function UserPage() {
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
        minHeight: "100vh",
        bgcolor: "#F7F4ED",
        p: 4,
      }}
    >
      <SubButton
        title="ノートを閉じる"
        size="lg"
        sx={{
          position: "absolute",
          top: "37px",
          left: "63px",
        }}
      />
      <Box
        sx={{
          textAlign: "center",
          display: "flex",
          justifyContent: "center",
          gap: "10%",
          mt: "100px",
        }}
      >
        <Box
          sx={{
            textAlign: "center",
            display: "flex",
            flexDirection: "column",
            alignItems: "flex-start",
            gap: "44px",
          }}
        >
          <Box
            sx={{
              textAlign: "center",
              display: "flex",
              alignItems: "center",
              justifyContent: "center",
              gap: "53px",
            }}
          >
            <Avatar
              alt="User Avatar"
              src="/path/to/profile-image.jpg"
              sx={{ width: 240, height: 240, mb: 2 }}
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
              <Typography variant="h5" sx={{ fontSize: "48px" }}>
                ほしょ
              </Typography>
              <Typography sx={{ fontSize: "25px" }}>
                ・居る連チャン: 5Days
              </Typography>
              <Typography sx={{ fontSize: "25px" }}>
                ・累計集中時間: 120h
              </Typography>
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
              <Typography align="center" sx={{ fontSize: "25px" }}>
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
                          top: -15,
                          left: 0,
                          backgroundColor: "#d4a373",
                          color: "#fff",
                          width: "31px",
                          height: "31px",
                          borderRadius: "50%",
                          display: "flex",
                          alignItems: "center",
                          justifyContent: "center",
                          fontSize: "16px",
                          fontWeight: "bold",
                          zIndex: 1,
                        }}
                      >
                        {index + 1}
                      </Box>
                      <Avatar
                        alt="Friend Avatar"
                        src="/path/to/profile-image.jpg"
                        sx={{ width: 69, height: 69 }}
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
              <Typography align="center" sx={{ fontSize: "25px" }}>
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
                          top: -15,
                          left: 0,
                          backgroundColor: "#d4a373",
                          color: "#fff",
                          width: "31px",
                          height: "31px",
                          borderRadius: "50%",
                          display: "flex",
                          alignItems: "center",
                          justifyContent: "center",
                          fontSize: "16px",
                          fontWeight: "bold",
                          zIndex: 1,
                        }}
                      >
                        {index + 1}
                      </Box>
                      <Avatar
                        alt="Friend Avatar"
                        src="/path/to/profile-image.jpg"
                        sx={{ width: 69, height: 69 }}
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
            gap: "23px",
            width: "fit-content",
          }}
        >
          <Typography align="center" sx={{ fontSize: "25px" }}>
            ご来店記録
          </Typography>
          <TableContainer sx={{ maxWidth: 500 }}>
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
          <Box sx={{ width: 300 }}>
            <BarChart
              series={[
                {
                  data: chartData.map((item) => item.hours),
                  color: "#D3AE6F",
                },
              ]}
              height={315}
              xAxis={[
                {
                  data: chartData.map((item) => item.day),
                  scaleType: "band",
                },
              ]}
              margin={{ top: 61, bottom: 30, left: 40, right: 10 }}
            />
          </Box>
        </Box>
      </Box>
    </Box>
  );
}

export default UserPage;
