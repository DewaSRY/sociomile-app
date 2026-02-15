import { z } from "zod";

export const PaginateMetaDataSchema = z.object({
  total: z.number(),
  page: z.number(),
  limit: z.number(),
});

export type PaginateMetaData = z.infer<typeof PaginateMetaDataSchema>;

export const SuccessResponseSchema = z.object({
  message: z.string(),
});

export type SuccessResponse = z.infer<typeof SuccessResponseSchema>;

export const CommonResponseSchema = z.object({
  message: z.string(),
  code: z.number(),
});

export type CommonResponse = z.infer<typeof CommonResponseSchema>;
