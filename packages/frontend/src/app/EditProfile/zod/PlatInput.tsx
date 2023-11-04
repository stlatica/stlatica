"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { Stack, TextField, Typography } from "@mui/material";
import React, { useState } from "react";

import { SubmitButton } from "@/components/button/SubmitButton";
import { PlatSchema, PlatSchemaType } from "@/zod/Plat/PlatSchema";
import { useDefaultForm } from "@/zod/utilities/useDefaultForm";

export const PlatInput: React.FC = () => {
  const [text, setText] = useState("");

  // フォームのhookを作成
  const { register, handleSubmit, formState } = useDefaultForm<PlatSchemaType>({
    defaultValues: {
      plat: "",
    },
    resolver: zodResolver(PlatSchema),
  });

  // エラーを抽出
  // 要調査: https://react-hook-form.com/docs/useformstate/errormessage
  const { errors } = formState;

  // formのsubmit用関数を作成
  // useFormの戻り値handleSubmitを使うことでzodバリデーション後に呼び出されることが保証される
  const onSubmit = handleSubmit((x) => {
    setText(x.plat);
  });

  return (
    <>
      <form onSubmit={onSubmit}>
        <Stack width={400}>
          {/* registerでフォームを制御する */}
          <TextField {...register("plat")} multiline rows={4} />
          {/* エラーメッセージは別途表示の必要あり */}
          {errors.plat?.message}
          <SubmitButton type="submit">Plat</SubmitButton>
        </Stack>
      </form>
      <Typography>{text}</Typography>
    </>
  );
};
