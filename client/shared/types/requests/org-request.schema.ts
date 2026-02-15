import { z } from "zod";
import { ConversationStatusSchema } from "../common/enum.schema";
export const AssignConversationRequestSchema = z.object({
  organizationStaffId: z.number(),
});

export type AssignConversationRequest = z.infer<
  typeof AssignConversationRequestSchema
>;

export const UpdateConversationRequestSchema = z.object({
  status: ConversationStatusSchema,
});

export type UpdateConversationRequest = z.infer<
  typeof UpdateConversationRequestSchema
>;

export const CreateTicketRequestSchema = z.object({
  conversationId: z.string(),
  name: z.string(),
});

export type CreateTicketRequest = z.infer<typeof CreateTicketRequestSchema>;

export const UpdateTicketRequestSchema = z.object({
  status: ConversationStatusSchema,
  name: z.string(),
});

export type UpdateTicketRequest = z.infer<typeof UpdateTicketRequestSchema>;