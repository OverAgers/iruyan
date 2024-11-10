import { z } from 'zod';

export const loginFormSchema = z.object({
  iruyanId: z.string().min(1, '入力してください'),
  password: z.string().min(1, '入力してください'),
});

export type LoginForm = z.infer<typeof loginFormSchema>;
