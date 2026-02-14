import { z } from 'zod';
import { userDataSchema, paginateMetaDataSchema } from './common.schema';

// Request Schemas
export const createOrganizationRequestSchema = z.object({
  name: z.string().min(3, 'Name must be at least 3 characters').max(100, 'Name must not exceed 100 characters'),
  owner_id: z.number(),
});

export const updateOrganizationRequestSchema = z.object({
  name: z.string().min(3, 'Name must be at least 3 characters').max(100, 'Name must not exceed 100 characters'),
});

// Response Schemas
export const organizationResponseSchema = z.object({
  id: z.number(),
  name: z.string(),
  owner_id: z.number(),
  owner: userDataSchema.optional(),
  created_at: z.string().or(z.date()),
  updated_at: z.string().or(z.date()),
});

export const organizationListResponseSchema = z.object({
  organizations: z.array(organizationResponseSchema),
  metadata: paginateMetaDataSchema,
});

// Type exports
export type CreateOrganizationRequest = z.infer<typeof createOrganizationRequestSchema>;
export type UpdateOrganizationRequest = z.infer<typeof updateOrganizationRequestSchema>;
export type OrganizationResponse = z.infer<typeof organizationResponseSchema>;
export type OrganizationListResponse = z.infer<typeof organizationListResponseSchema>;
