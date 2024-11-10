import { z } from 'zod';

export const registerFormSchema = z
  .object({
    iruyanId: z
      .string()
      .min(6, '6文字以上で入力してください')
      .regex(
        /^(?=.*[a-zA-Z])(?=.*\d)/,
        '英字と数字をそれぞれ少なくとも1つ含めてください'
      ),
    name: z.string().min(1, 'ユーザー名を入力してください'),
    email: z.string().email('有効なメールアドレスを入力してください'),
    password: z.string().min(8, 'パスワードは8文字以上で入力してください'),
    passwordConfirm: z.string(),
  })
  .refine((data) => data.password === data.passwordConfirm, {
    message: 'パスワードが一致しません。',
    path: ['passwordConfirm'],
  });

export type RegisterForm = z.infer<typeof registerFormSchema>;
