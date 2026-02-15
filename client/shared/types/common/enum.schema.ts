import * as z from "zod";

export const ButtonVariantTypeSchema = z.enum([
  "text",
  "flat",
  "elevated",
  "outlined",
  "plain",
  "tonal",
]);

export type ButtonVariantType = z.infer<typeof ButtonVariantTypeSchema>;


export const ConversationStatusSchema = z.enum([
  "pending",
  "in_progress",
  "done",
]);

export type ConversationStatus = z.infer<typeof ConversationStatusSchema>;
