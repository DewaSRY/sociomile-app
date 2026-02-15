import z from "zod";

export const FiltersSchema = z.object({
  page: z.number().default(1),
  limit: z.number().default(20),
});

export type Filters = z.infer<typeof FiltersSchema>;