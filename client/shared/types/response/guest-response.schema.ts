import { z } from "zod";
import { UserDataSchema } from "./auth-response.schema";
import { PaginateMetaDataSchema } from "./pagination-response.schema";
import { OrganizationResponseSchema } from "./organization-response.schema";

export const ConversationMessageResponseSchema = z.object({
  id: z.number().int().nonnegative(),
  organizationId: z.number().int().nonnegative(),
  conversationId: z.number().int().nonnegative(),

  createdById: z.number().int().nonnegative(),
  createdBy: UserDataSchema.optional(),

  message: z.string(),

  createdAt: z.coerce.date(),
  updatedAt: z.coerce.date(),
});

export const ConversationMessagePaginateSchema = z.object({
  data: z.array(ConversationMessageResponseSchema),
  metadata: PaginateMetaDataSchema,
});

export const ConversationResponseSchema = z.object({
  id: z.number().int().nonnegative(),

  organizationId: z.number().int().nonnegative(),

  organization: OrganizationResponseSchema.nullable().optional(),

  guestId: z.number().int().nonnegative(),

  guest: UserDataSchema.nullable().optional(),

  organizationStaffId: z.number().int().nonnegative().nullable().optional(),

  organizationStaff: UserDataSchema.nullable().optional(),

  messages: ConversationMessageResponseSchema.array().default([]),
  status: z.string(),

  createdAt: z.coerce.date(),

  updatedAt: z.coerce.date(),
});

export type ConversationResponse = z.infer<typeof ConversationResponseSchema>;

export const ConversationListResponseSchema = z.object({
  conversations: z.array(ConversationResponseSchema),
  metadata: PaginateMetaDataSchema,
});

export type ConversationListResponse = z.infer<
  typeof ConversationListResponseSchema
>;

export type ConversationMessagePaginate = z.infer<
  typeof ConversationMessagePaginateSchema
>;
