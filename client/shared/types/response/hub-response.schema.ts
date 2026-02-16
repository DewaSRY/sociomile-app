import { z } from "zod";
import { PaginateMetaDataSchema } from "./pagination-response.schema";

export const HubOrganizationResponseSchema = z.object({
  id: z.number().int().nonnegative(),
  name: z.string(),
  ownerName: z.string(),
  ownerId: z.number().int().nonnegative(),
  createdAt: z.coerce.date(),
  updatedAt: z.coerce.date(),
});

export type HubOrganizationResponse = z.infer<
  typeof HubOrganizationResponseSchema
>;

export const HubOrganizationPaginateResponseSchema = z.object({
  data: HubOrganizationResponseSchema.array(),
  metadata: PaginateMetaDataSchema,
});

export type HubOrganizationPaginateResponse = z.infer<
  typeof HubOrganizationPaginateResponseSchema
>;
