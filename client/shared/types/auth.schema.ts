import { z } from 'zod';
import { userDataSchema } from './common.schema';

export const registerRequestSchema = z.object({
  email: z.string().email('Invalid email format'),
  password: z.string().min(6, 'Password must be at least 6 characters'),
  name: z.string().min(2, 'Name must be at least 2 characters'),
});

export const loginRequestSchema = z.object({
  email: z.string().email('Invalid email format'),
  password: z.string().min(1, 'Password is required'),
});

export const refreshTokenRequestSchema = z.object({
  token: z.string().min(1, 'Token is required'),
});

export const authResponseSchema = z.object({
  token: z.string(),
  user: userDataSchema,
});

export type RegisterRequest = z.infer<typeof registerRequestSchema>;
export type LoginRequest = z.infer<typeof loginRequestSchema>;
export type RefreshTokenRequest = z.infer<typeof refreshTokenRequestSchema>;
export type AuthResponse = z.infer<typeof authResponseSchema>;
