import { z } from "zod";

// ✅ ユーザー入力時のバリデーション
export const loginFormSchema = z.object({
  iruyanId: z.string().min(1, "入力してください"),
  password: z.string().min(1, "入力してください"),
});

export type LoginForm = z.infer<typeof loginFormSchema>;

// ✅ API リクエスト用のスキーマ（同じでも定義分けると明確）
export const postLoginRequestSchema = loginFormSchema;
export type PostLoginRequest = z.infer<typeof postLoginRequestSchema>;

// ✅ API レスポンス用のスキーマ
export const postLoginResponseSchema = z.object({
  message: z.string(),
  user: z.object({
    iruyanId: z.string(),
    userName: z.string(),
    email: z.string().email(),
  }),
});

export type PostLoginResponse = z.infer<typeof postLoginResponseSchema>;
