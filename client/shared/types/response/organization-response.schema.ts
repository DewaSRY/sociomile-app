import { z } from "zod";
import { UserDataSchema } from "./auth-response.schema";

export const OrganizationResponseSchema = z.object({
  id: z.number().int().nonnegative(),

  name: z.string(),

  ownerId: z.number().int().nonnegative(),

  owner: UserDataSchema.nullable().optional(),

  createdAt: z.string().datetime(),

  updatedAt: z.string().datetime(),
});
