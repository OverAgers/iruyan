import axios from "axios";
import { z } from "zod";

export const postRoomEnterSchema = z.object({
  iruyanId: z.string(),
  roomId: z.string(),
});

export type PostRoomEnterRequest = z.infer<typeof postRoomEnterSchema>;

export default function usePostRoomEnterRequest() {
  const baseURL = process.env.NEXT_PUBLIC_API_URL;

  const entry = async ({
    iruyanId,
    roomId,
  }: PostRoomEnterRequest & { roomId: string }) => {
    const requestURL = `${baseURL}/rooms/${roomId}/enter/${iruyanId}`;

    try {
      await axios.post(
        requestURL,
        new URLSearchParams({
          iruyanId,
          roomId,
        }),
        {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
        }
      );

      // バリデーション（必須ではないなら削除可）
      return postRoomEnterSchema.parse({ iruyanId, roomId });
    } catch (error) {
      console.error("API エラー:", error);
      throw error;
    }
  };

  return { entry };
}
