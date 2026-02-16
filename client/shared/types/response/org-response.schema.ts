import { z } from "zod";
import { UserDataSchema } from "./auth-response.schema";
import { PaginateMetaDataSchema } from "./pagination-response.schema";
import { OrganizationResponseSchema } from "./organization-response.schema";

export const OrganizationStaffRecordSchema = z.object({
  id: z.number(),
  name: z.string(),
  roleName: z.string(),
  email: z.string(),
});

export type OrganizationStaffRecord = z.infer<
  typeof OrganizationStaffRecordSchema
>;

export const OrganizationStaffPaginationSchema = z.object({
  data: OrganizationStaffRecordSchema.array(),
  metadata: PaginateMetaDataSchema,
});

export type OrganizationStaffPagination = z.infer<
  typeof OrganizationStaffPaginationSchema
>;

export const TicketResponseSchema = z.object({
  id: z.number().int().nonnegative(),

  organizationId: z.number().int().nonnegative(),
  organization: OrganizationResponseSchema.optional(),

  conversationId: z.number().int().nonnegative(),

  createdById: z.number().int().nonnegative(),
  createdBy: UserDataSchema.optional(),

  ticketNumber: z.string(),
  name: z.string(),
  status: z.string(),

  createdAt: z.string().datetime(),
  updatedAt: z.string().datetime(),
});

export type TicketResponse = z.infer<typeof TicketResponseSchema>;

export const TicketListResponseSchema = z.object({
  tickets: z.array(TicketResponseSchema),
  metadata: PaginateMetaDataSchema,
});

export type TicketListResponse = z.infer<typeof TicketListResponseSchema>;