import { z } from "zod";

/**
 * 🔹 共通スキーマ：パスワード確認を含む（フロント用）
 */
const baseRegisterSchema = z.object({
  iruyanId: z
    .string()
    .min(1, "ユーザーIDを入力してください")
    .min(6, "ユーザーIDは6文字以上で入力してください")
    .max(50, "ユーザーIDは50文字以内で入力してください")
    .regex(/^[a-zA-Z0-9_-]+$/, "英数字、ハイフン、アンダースコアのみ使用できます")
    .regex(/^(?=.*[a-zA-Z])(?=.*\d)/, "英字と数字をそれぞれ少なくとも1つ含めてください"),
  name: z
    .string()
    .min(1, "ユーザー名を入力してください")
    .max(50, "ユーザー名は50文字以内で入力してください")
    .regex(/^[a-zA-Z0-9ひらがなカタカナ漢字\u3000-\u303F\u3040-\u309F\u30A0-\u30FF\u4E00-\u9FAF_-\s]+$/, "使用できない文字が含まれています"),
  email: z
    .string()
    .min(1, "メールアドレスを入力してください")
    .email("有効なメールアドレスを入力してください")
    .max(254, "メールアドレスが長すぎます"),
  password: z
    .string()
    .min(1, "パスワードを入力してください")
    .min(8, "パスワードは8文字以上で入力してください")
    .max(128, "パスワードは128文字以内で入力してください")
    .regex(/[0-9]/, "少なくとも1つの数字を含めてください")
    .regex(/[a-zA-Z]/, "少なくとも1つの英字を含めてください"),
  passwordConfirm: z.string().min(1, "パスワード確認を入力してください"),
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
