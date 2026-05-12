import { valibotResolver } from "@hookform/resolvers/valibot";
import { Button, Stack, TextInput } from "@mantine/core";
import * as v from "valibot";

import { useSafeForm } from "@/features/type/useForm";
import { usePostPlat } from "@/openapi/stlaticaInternalApi";

const schema = v.object({
  content: v.string(),
});

type SchemaType = v.InferOutput<typeof schema>;

export default function LoginScene() {
  const { trigger } = usePostPlat({});

  const { register, handleSubmit } = useSafeForm<SchemaType>({
    resolver: valibotResolver(schema),
    defaultValues: {
      content: "",
    },
  });

  const onSubmit = handleSubmit(async (val) => {
    await trigger({
      content: val.content,
    });
  });

  return (
    <main>
      <form onSubmit={onSubmit}>
        <Stack>
          <TextInput {...register("content")} label="content" />
          <Button type="submit">Plat</Button>
        </Stack>
      </form>
    </main>
  );
}
