import { z } from "zod";

// ✅ ユーザー入力時のバリデーション
export const loginFormSchema = z.object({
  iruyanId: z
    .string()
    .min(1, "ユーザーIDを入力してください")
    .max(50, "ユーザーIDは50文字以内で入力してください")
    .regex(/^[a-zA-Z0-9_-]+$/, "ユーザーIDは英数字、ハイフン、アンダースコアのみ使用できます"),
  password: z
    .string()
    .min(1, "パスワードを入力してください")
    .min(4, "パスワードは4文字以上で入力してください")
    .max(128, "パスワードは128文字以内で入力してください"),
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
