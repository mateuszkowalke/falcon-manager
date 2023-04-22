import { z } from "zod";

export const keystoneUserSchema = z.object({
  firstName: z.string(),
  lastName: z.string(),
  email: z.string(),
});

export type KeystoneUser = z.infer<typeof keystoneUserSchema>;
