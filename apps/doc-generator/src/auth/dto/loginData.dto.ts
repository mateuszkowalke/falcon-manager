import { z } from "zod";

export const loginDataSchema = z.object({
  username: z.string(),
  password: z.string()
});

export type LoginData = z.infer<typeof loginDataSchema>;
