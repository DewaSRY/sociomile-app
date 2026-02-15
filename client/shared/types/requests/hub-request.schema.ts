import {  z } from "zod";

export const RegisterOrganizationRequestSchema = z.object({
  name : z.string(),
  email : z.string(),
  ownerName : z.string(),
  password : z.string(),
});

export type RegisterOrganizationRequest = z.infer<
  typeof RegisterOrganizationRequestSchema
>;

