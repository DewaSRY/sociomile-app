import { z } from "zod";

export const RegisterRequestSchema = z.object({
  email: z.string(),
  password: z.string(),
  name: z.string(),
});

export type RegisterRequest = z.infer<typeof RegisterRequestSchema>;

export const LoginRequestSchema = z.object({
  email: z.string(),
  password: z.string(),
  name: z.string(),
});

export type LoginRequest = z.infer<typeof LoginRequestSchema>;

export const RefreshTokenRequestSchema = z.object({
  token: z.string(),
});

export type RefreshTokenRequest = z.infer<typeof RefreshTokenRequestSchema>;
