import { CreateID } from "@/features//acl/CreateID";

import type { Plat } from "@/openapi/stlaticaInternalApi.schemas";

/**
 *
 */
export const GenerateDummyPlat = (): Plat => {
  return {
    content: "test content",
    created_at: new Date().toISOString(),
    plat_id: CreateID(),
    images: [],
  } satisfies Plat;
};
