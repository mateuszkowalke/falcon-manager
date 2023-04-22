import { z } from "zod";

export const generateDocFormInput = z.object({
  officeId: z.string(),
  falconsIds: z.array(z.string()),
});

export type GenerateDocFormInput = z.infer<typeof generateDocFormInput>;
