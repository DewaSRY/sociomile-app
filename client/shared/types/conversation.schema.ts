import { z } from 'zod';
import { userDataSchema, paginateMetaDataSchema } from './common.schema';
import { organizationResponseSchema } from './organization.schema';
// Request Schemas
export const createConversationRequestSchema = z.object({
  organization_id: z.number(),
});

export const updateConversationRequestSchema = z.object({
  status: z.enum(['pending', 'in_progress', 'done']),
});

export const assignConversationRequestSchema = z.object({
  organization_staff_id: z.number(),
});

export const conversationResponseSchema = z.object({
  id: z.number(),
  organization_id: z.number(),
  organization: organizationResponseSchema.optional(),
  guest_id: z.number(),
  guest: userDataSchema.optional(),
  organization_staff_id: z.number().nullable().optional(),
  organization_staff: userDataSchema.optional(),
  status: z.string(),
  created_at: z.string().or(z.date()),
  updated_at: z.string().or(z.date()),
});

export const conversationListResponseSchema = z.object({
  conversations: z.array(conversationResponseSchema),
  metadata: paginateMetaDataSchema,
});

// Type exports
export type CreateConversationRequest = z.infer<typeof createConversationRequestSchema>;
export type UpdateConversationRequest = z.infer<typeof updateConversationRequestSchema>;
export type AssignConversationRequest = z.infer<typeof assignConversationRequestSchema>;
export type ConversationResponse = z.infer<typeof conversationResponseSchema>;
export type ConversationListResponse = z.infer<typeof conversationListResponseSchema>;
