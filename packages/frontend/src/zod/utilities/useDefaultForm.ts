// from: https://zenn.dev/yuitosato/articles/292f13816993ef#1.-useform%E3%82%92%E3%83%A9%E3%83%83%E3%83%97%E3%81%97%E3%81%A6%E3%82%BF%E3%82%A4%E3%83%97%E3%82%BB%E3%83%BC%E3%83%95%E3%81%AB%E3%81%99%E3%82%8B

import { type UseFormProps, type UseFormReturn, useForm } from "react-hook-form";

/**
 * useFormを型安全に利用するためのユーティリティ
 */
const useDefaultForm = <FORM_TYPE extends Record<string, unknown>>(
  props: UseFormProps<FORM_TYPE> & {
    defaultValues: FORM_TYPE;
  },
): UseFormReturn<FORM_TYPE> => {
  return useForm(props);
};

export { useDefaultForm };
