import z from "zod";

export const PlatSchema = z.object({
  plat: z.string().min(1).max(140),
});

export type PlatSchemaType = z.infer<typeof PlatSchema>;
