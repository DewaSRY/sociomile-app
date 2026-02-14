import { z } from 'zod';
import { userDataSchema, paginateMetaDataSchema } from './common.schema';
import { conversationResponseSchema } from './conversation.schema';

// Request Schemas
export const createConversationMessageRequestSchema = z.object({
  conversation_id: z.number(),
  message: z.string().min(1, 'Message is required').max(5000, 'Message must not exceed 5000 characters'),
});

// Response Schemas
export const conversationMessageResponseSchema = z.object({
  id: z.number(),
  organization_id: z.number(),
  conversation_id: z.number(),
  conversation: conversationResponseSchema.optional(),
  created_by_id: z.number(),
  created_by: userDataSchema.optional(),
  message: z.string(),
  created_at: z.string().or(z.date()),
  updated_at: z.string().or(z.date()),
});

export const conversationMessageListResponseSchema = z.object({
  messages: z.array(conversationMessageResponseSchema),
  metadata: paginateMetaDataSchema,
});

// Type exports
export type CreateConversationMessageRequest = z.infer<typeof createConversationMessageRequestSchema>;
export type ConversationMessageResponse = z.infer<typeof conversationMessageResponseSchema>;
export type ConversationMessageListResponse = z.infer<typeof conversationMessageListResponseSchema>;
