import { z } from "zod";

export const CreateConversationRequestSchema = z.object({
  organizationId: z.number(),
});

export type CreateConversationRequest = z.infer<
  typeof CreateConversationRequest
>;

export const CreateConversationMessageSchema = z.object({
  conversationId: z.number(),
  message: z.string()
});

export type CreateConversationMessage = z.infer<
  typeof CreateConversationMessageSchema
>;