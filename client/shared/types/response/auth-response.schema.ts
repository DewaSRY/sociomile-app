import { z } from "zod";

import { OrganizationResponseSchema } from "$shared/types/response/organization-response.schema";

export const UserDataSchema = z.object({
  id: z.number(),
  email: z.string(),
  name: z.string(),
});

export type UserData = z.infer<typeof UserDataSchema>;

export const UserProfileDataSchema = z.object({
  id: z.number(),
  email: z.string(),
  name: z.string(),
  RoleName: z.string(),
  organization: OrganizationResponseSchema.optional(),
});

export type UserProfileData = z.infer<typeof UserProfileDataSchema>;

export const AuthResponseSchema = z.object({
  token: z.string(),
  user: UserProfileDataSchema,
});

export type AuthResponse = z.infer<typeof AuthResponseSchema>;
