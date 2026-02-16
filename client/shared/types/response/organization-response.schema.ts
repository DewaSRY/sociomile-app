import { z } from "zod";
import { PaginateMetaDataSchema } from "./pagination-response.schema";

export const OrganizationResponseSchema = z.object({
  id: z.number().int().nonnegative(),

  name: z.string(),

  ownerId: z.number().int().nonnegative(),

  createdAt: z.string().datetime(),

  updatedAt: z.string().datetime(),
});

export type OrganizationResponse = z.infer<typeof OrganizationResponseSchema>;

export const OrganizationPaginateResponseSchema = z.object({
  data: OrganizationResponseSchema.array(),
  metadata: PaginateMetaDataSchema,
});

export type OrganizationPaginateResponse = z.infer<
  typeof OrganizationPaginateResponseSchema
>;
