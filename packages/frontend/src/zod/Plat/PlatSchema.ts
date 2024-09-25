import * as v from "valibot";

export const PlatSchema = v.object({
  plat: v.pipe(v.string(), v.minLength(1), v.maxLength(140)),
});

export type PlatSchemaType = v.InferOutput<typeof PlatSchema>;
