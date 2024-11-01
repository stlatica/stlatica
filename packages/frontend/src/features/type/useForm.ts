import { useForm, type UseFormProps, type UseFormReturn } from "react-hook-form";

/**
 * デフォルト値が型安全な useForm
 */
const useSafeForm = <FORM_TYPE extends Record<string, unknown>>(
  props: UseFormProps<FORM_TYPE> & {
    defaultValues: FORM_TYPE;
  }
): UseFormReturn<FORM_TYPE> => {
  return useForm(props);
};

export { useSafeForm };
