import axios from "axios";
import useSWRMutation from "swr/mutation";
import { z } from "zod";

export const postCreateRoomRequestSchema = z.object({
  roomName: z.string(),
});

export const postCreateRoomResponseSchema = z.object({
  message: z.string(),
  roomId: z.string().uuid(),
  roomName: z.string(),
});

export type PostCreateRoomRequest = z.infer<typeof postCreateRoomRequestSchema>;
export type PostCreateRoomResponse = z.infer<typeof postCreateRoomResponseSchema>;

export default function usePostCreateRoomRequest() {
  const requestURL = `${process.env.NEXT_PUBLIC_API_URL}/rooms`;

  const fetcher = async (
    _: string,
    { arg }: { arg: PostCreateRoomRequest }
  ): Promise<PostCreateRoomResponse> => {
    const res = await axios.post(requestURL, arg, {
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
    });
    return postCreateRoomResponseSchema.parse(res.data);
  };

  const { trigger, data, error, isMutating } = useSWRMutation(requestURL, fetcher);

  // 外部から使いやすくするため trigger をラップ
  const createRoom = (arg: PostCreateRoomRequest) => trigger(arg);

  return { createRoom, data, error, isMutating };
}
