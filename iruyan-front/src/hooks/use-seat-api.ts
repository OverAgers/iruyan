/**
 * 座席API操作のための型安全なカスタムフック
 */

import { useCallback } from "react";
import { useAPI, useMutation } from "./api-hooks";
import { seatService } from "@/services/api";
import type {
  SeatStatusResponse,
} from "@/types/api";

// 座席状態を取得するフック
export function useSeatStatus(roomId: string) {
  return useAPI<SeatStatusResponse[]>(
    roomId ? `/rooms/${roomId}/seats/status` : null,
    {
      refreshInterval: 5000, // 5秒ごとに自動更新
      revalidateOnFocus: true,
    }
  );
}

// 座席を取るためのミューテーション
export function useTakeSeat() {
  return useMutation(async ({ roomId, seatNumber, iruyanId }: {
    roomId: string;
    seatNumber: number;
    iruyanId: string;
  }) => {
    const response = await seatService.takeSeat(roomId, seatNumber, { iruyanId });
    if (!response.success) {
      throw new Error(response.error || "座席の取得に失敗しました");
    }
    return response.data;
  });
}

// 座席を離れるためのミューテーション
export function useLeaveSeat() {
  return useMutation(async ({ roomId, seatNumber, iruyanId }: {
    roomId: string;
    seatNumber: number;
    iruyanId: string;
  }) => {
    const response = await seatService.leaveSeat(roomId, seatNumber, { iruyanId });
    if (!response.success) {
      throw new Error(response.error || "座席から離れることに失敗しました");
    }
    return response.data;
  });
}

// 座席操作の統合フック
export function useSeatOperations() {
  const takeSeat = useTakeSeat();
  const leaveSeat = useLeaveSeat();

  const handleSeatAction = useCallback(async ({
    action,
    roomId,
    seatNumber,
    iruyanId,
  }: {
    action: "take" | "leave";
    roomId: string;
    seatNumber: number;
    iruyanId: string;
  }) => {
    try {
      if (action === "take") {
        await takeSeat.mutate({ roomId, seatNumber, iruyanId });
      } else {
        await leaveSeat.mutate({ roomId, seatNumber, iruyanId });
      }
    } catch (error) {
      console.error(`Seat ${action} operation failed:`, error);
      throw error;
    }
  }, [takeSeat, leaveSeat]);

  return {
    handleSeatAction,
    isLoading: takeSeat.isLoading || leaveSeat.isLoading,
    error: takeSeat.error || leaveSeat.error,
  };
}