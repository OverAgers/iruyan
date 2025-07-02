// Updated to fix ESLint errors - StatusType and any type issues resolved
import React, { useEffect } from "react";
import SeatCard from "@/components/ui/card/seat-card";
import { Box, Grid2, Typography } from "@mui/material";
import useSeatStore from "@/stores/seats-store";
import useUserStore from "@/stores/user-store";
import { useParams } from "next/navigation";

const API_URL = process.env.NEXT_PUBLIC_API_URL;

export default function Seats() {
  const params = useParams();
  const roomId = params?.roomId as string;

  const currentUser = useUserStore((state) => state.currentUser);
  const seats = useSeatStore((state) => state.seats);
  const moveSeat = useSeatStore((state) => state.moveSeat);
  const leaveSeat = useSeatStore((state) => state.leaveSeat);
  const setSeatUser = useSeatStore((state) => state.setSeatUser); // 追加：座席にユーザーを割り当てる関数

  useEffect(() => {
    if (!roomId || !currentUser) return;

    const fetchSeatStatus = async () => {
      try {
        const res = await fetch(`${API_URL}/rooms/${roomId}/seats/status`);
        if (!res.ok) throw new Error("座席情報の取得に失敗しました");

        const seatStatus = await res.json();
        console.log("seatStatus:", seatStatus); // 中身確認

        if (!Array.isArray(seatStatus)) {
          throw new Error("座席情報の形式が不正です");
        }

        seatStatus.forEach((seat: {
          seat_number: number;
          iruyan_id?: string;
          user_name?: string;
        }) => {
          setSeatUser(seat.seat_number, {
            seatId: String(seat.seat_number), // seatId は文字列
            seatNumber: seat.seat_number,
            roomId: roomId,
            isVacant: false,
            iruyanId: seat.iruyan_id || "",
            userName: seat.user_name || "",
            userImage: "", // 現時点でAPIにないので空文字
            note: "",
            task: "",
          });
        });


      } catch (err: unknown) {
        const errorMessage = err instanceof Error ? err.message : 'Unknown error occurred';
        alert(`座席情報取得エラー: ${errorMessage}`);
      }
    };

    fetchSeatStatus();
  }, [roomId, currentUser, setSeatUser]);

  if (!currentUser) {
    return <Typography>ログインしてください</Typography>;
  }

  if (!roomId) {
    return <Typography>部屋IDが取得できません</Typography>;
  }

  const handleSeatClick = async (seatId: string, isVacant: boolean) => {
    const seat = seats.find((s) => s.seatId === seatId);
    const seatNumber = seat?.seatNumber;

    if (!seatNumber) {
      alert("座席番号が取得できません");
      return;
    }

    const url = `${API_URL}/rooms/${roomId}/seats/${seatNumber}/${isVacant ? "take" : "leave"}`;
    const formData = new URLSearchParams({ iruyanId: currentUser.iruyanId });

    try {
      const res = await fetch(url, {
        method: "PUT",
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
        },
        body: formData,
      });

      const result = await res.json();

      if (!res.ok) {
        throw new Error(result.message || "リクエストに失敗しました");
      }

      if (isVacant) {
        moveSeat(seatId, currentUser); // 着席処理
      } else {
        leaveSeat(seatId, currentUser); // 離席処理
      }
    } catch (error: unknown) {
      const errorMessage = error instanceof Error ? error.message : 'Unknown error occurred';
      alert(`操作に失敗しました: ${errorMessage}`);
    }
  };

  return (
    <Box sx={{ width: 800 }}>
      <Grid2 container spacing={2}>
        {seats.map((seat) => {
          const isCurrentUser = seat.iruyanId === currentUser.iruyanId;

          return (
            <Grid2 key={seat.seatId}>
              <SeatCard
                seatId={`${seat.seatNumber}`}
                isVacant={seat.isVacant}
                onClick={() => handleSeatClick(seat.seatId, seat.isVacant)}
                name={seat.userName || ""}
                image={seat.userImage || ""}
                note={seat.note || ""}
                task={seat.task || ""}
                isCurrentUser={isCurrentUser}
              />
            </Grid2>
          );
        })}
      </Grid2>
    </Box>
  );
}
