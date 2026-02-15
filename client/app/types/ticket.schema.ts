import { z } from 'zod';
import { userDataSchema, paginateMetaDataSchema } from './common.schema';
import { organizationResponseSchema } from './organization.schema';
import { conversationResponseSchema } from './conversation.schema';

export const createTicketRequestSchema = z.object({
  conversation_id: z.number(),
  name: z.string().min(3, 'Name must be at least 3 characters').max(200, 'Name must not exceed 200 characters'),
});

export const updateTicketRequestSchema = z.object({
  name: z.string().min(3, 'Name must be at least 3 characters').max(200, 'Name must not exceed 200 characters').optional(),
  status: z.enum(['pending', 'in_progress', 'done']).optional(),
});

export const ticketResponseSchema = z.object({
  id: z.number(),
  organization_id: z.number(),
  organization: organizationResponseSchema.optional(),
  conversation_id: z.number(),
  conversation: conversationResponseSchema.optional(),
  created_by_id: z.number(),
  created_by: userDataSchema.optional(),
  ticket_number: z.string(),
  name: z.string(),
  status: z.string(),
  created_at: z.string().or(z.date()),
  updated_at: z.string().or(z.date()),
});

export const ticketListResponseSchema = z.object({
  tickets: z.array(ticketResponseSchema),
  metadata: paginateMetaDataSchema,
});

export type CreateTicketRequest = z.infer<typeof createTicketRequestSchema>;
export type UpdateTicketRequest = z.infer<typeof updateTicketRequestSchema>;
export type TicketResponse = z.infer<typeof ticketResponseSchema>;
export type TicketListResponse = z.infer<typeof ticketListResponseSchema>;
