import { z } from 'zod';

export const userRoleSchema = z.object({
  id: z.number(),
  name: z.string(),
});

export const userDataSchema = z.object({
  id: z.number(),
  email: z.string().email(),
  name: z.string(),
  role_id: z.number(),
  role: userRoleSchema.optional(),
  organization_id: z.number().nullable().optional(),
});

export const errorResponseSchema = z.object({
  message: z.string(),
  code: z.number(),
  error: z.string().optional(),
});

export const commonResponseSchema = z.object({
  message: z.string(),
  data: z.any().optional(),
  code: z.number(),
});

export const successResponseSchema = z.object({
  message: z.string(),
});

export const paginateMetaDataSchema = z.object({
  total: z.number(),
  page: z.number(),
  limit: z.number(),
});

export type UserRole = z.infer<typeof userRoleSchema>;
export type UserData = z.infer<typeof userDataSchema>;
export type ErrorResponse = z.infer<typeof errorResponseSchema>;
export type CommonResponse = z.infer<typeof commonResponseSchema>;
export type SuccessResponse = z.infer<typeof successResponseSchema>;
export type PaginateMetaData = z.infer<typeof paginateMetaDataSchema>;
