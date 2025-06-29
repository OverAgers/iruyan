import { z } from "zod";

/**
 * 🔹 共通スキーマ：パスワード確認を含む（フロント用）
 */
const baseRegisterSchema = z.object({
  iruyanId: z
    .string()
    .min(6, "6文字以上で入力してください")
    .regex(/^(?=.*[a-zA-Z])(?=.*\d)/, "英字と数字をそれぞれ少なくとも1つ含めてください"),
  name: z.string().min(1, "ユーザー名を入力してください"),
  email: z.string().email("有効なメールアドレスを入力してください"),
  password: z
    .string()
    .min(8, "パスワードは8文字以上で入力してください")
    .regex(/[0-9]/, "少なくとも1つの数字を含めてください")
    .regex(/[a-zA-Z]/, "少なくとも1つの英字を含めてください"),
  passwordConfirm: z.string(),
});

/**
 * ✅ フロントバリデーション用：パスワード一致確認付き
 */
export const registerFormSchema = baseRegisterSchema.refine(
  (data) => data.password === data.passwordConfirm,
  {
    message: "パスワードが一致しません。",
    path: ["passwordConfirm"],
  }
);

export type RegisterForm = z.infer<typeof registerFormSchema>;

/**
 * ✅ API送信用：passwordConfirm を除いたバージョン
 */
export const postRegisterRequestSchema = baseRegisterSchema.omit({
  passwordConfirm: true,
});

export type PostRegisterRequest = z.infer<typeof postRegisterRequestSchema>;

/**
 * ✅ APIレスポンス用スキーマ
 */
export const postRegisterResponseSchema = z.object({
  message: z.string(),
  user: z.object({
    iruyanId: z.string(),
    userName: z.string(),
    email: z.string().email(),
  }),
});

export type PostRegisterResponse = z.infer<typeof postRegisterResponseSchema>;
