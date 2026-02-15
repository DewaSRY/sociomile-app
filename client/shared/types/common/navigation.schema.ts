import { z } from "zod";
import { ButtonVariantTypeSchema } from "./enum.schema";

export const LandingPageAuthNavSchema = z.object({
  href: z.string(),
  name: z.string(),
  variant: ButtonVariantTypeSchema.optional(),
  icon: z.string(),
  color: z.string().optional(),
});
export type LandingPageAuthNav = z.infer<typeof LandingPageAuthNavSchema>;

export const AruskuSubMenuSchema = z.object({
  name: z.string(),
  icon: z.string(),
  href: z.string(),
});
export type AruskuSubMenu = z.infer<typeof AruskuSubMenuSchema>;

export const AruskuMenuSchema = z.object({
  name: z.string(),
  subMenu: AruskuSubMenuSchema.array(),
});
export type AruskuMenu = z.infer<typeof AruskuMenuSchema>;

export const BreadcrumbItemSchema = z.object({
  title: z.string(),
  to: z.string().optional(),
  href: z.string().optional(),
  disabled: z.boolean().optional(),
});
export type BreadcrumbItem = z.infer<typeof BreadcrumbItemSchema>;
