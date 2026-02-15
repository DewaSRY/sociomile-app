import { z } from "zod";
import { UserDataSchema } from "./auth-response.schema";
import {PaginateMetaDataSchema} from "./pagination-response.schema"
import {OrganizationResponseSchema} from "./organization-response.schema"

export const ConversationResponseSchema = z.object({
  id: z.number().int().nonnegative(),

  organizationId: z.number().int().nonnegative(),

  organization: OrganizationResponseSchema.nullable().optional(),

  guestId: z.number().int().nonnegative(),

  guest: UserDataSchema.nullable().optional(),

  organizationStaffId: z.number().int().nonnegative().nullable().optional(),

  organizationStaff: UserDataSchema.nullable().optional(),

  status: z.string(),

  createdAt: z.string().datetime(),

  updatedAt: z.string().datetime(),
});

export const ConversationListResponseSchema = z.object({
  conversations: z.array(ConversationResponseSchema),
  metadata: PaginateMetaDataSchema,
});