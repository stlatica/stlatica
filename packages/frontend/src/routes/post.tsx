import { valibotResolver } from "@hookform/resolvers/valibot";
import { Button, Stack, TextInput } from "@mantine/core";
import * as v from "valibot";

import { useSafeForm } from "@/features/type/useForm";
import { usePostPlat } from "@/openapi/stlaticaInternalApi";

const schema = v.object({
  content: v.string(),
  user_id: v.string(),
});

type schemaType = v.InferOutput<typeof schema>;

export default function LoginScene() {
  const { trigger } = usePostPlat({});

  const { register, handleSubmit } = useSafeForm<schemaType>({
    resolver: valibotResolver(schema),
    defaultValues: {
      content: "",
      user_id: "",
    },
  });

  const onSubmit = handleSubmit(async (val) => {
    await trigger({
      content: val.content,
      user_id: val.user_id,
    });
  });

  return (
    <main>
      <form onSubmit={onSubmit}>
        <Stack>
          <TextInput {...register("content")} label="content" />
          <TextInput {...register("user_id")} label="user_id" />
          <Button type="submit">Plat</Button>
        </Stack>
      </form>
    </main>
  );
}
