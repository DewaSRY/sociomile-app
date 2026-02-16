import { z } from "zod";

export const UserDataSchema = z.object({
  id: z.number(),
  email: z.string(),
  name: z.string(),
});

export type UserData = z.infer<typeof UserDataSchema>;

export const AuthResponseSchema = z.object({
  token: z.string(),
  user: UserDataSchema,
});

export type AuthResponse = z.infer<typeof AuthResponseSchema>;
