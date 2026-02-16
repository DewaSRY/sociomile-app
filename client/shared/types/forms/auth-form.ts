import  {z} from "zod";

import { RegisterRequestSchema } from "$shared/types/requests/auth-request.schema";

export const SignupFormSchema = RegisterRequestSchema.extend({
  confirmPassword: z.string(),
}).refine((data) => data.password === data.confirmPassword, {
  message: "Passwords do not match",
  path: ["confirmPassword"],
});

export type SignupForm = z.infer<typeof SignupFormSchema>;
