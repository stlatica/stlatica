import { useEffect } from "react";
import useSwr from "swr";

import { IsNull } from "@/utilities/utilities";

/**
 *
 */
const useStaticSWR = <T>(key: string, initialData: T) => {
  const { data = initialData, mutate } = useSwr<T>(key, null, {
    fallbackData: initialData,
    revalidateOnFocus: false,
    revalidateOnMount: false,
    revalidateOnReconnect: false,
    revalidateIfStale: false,
  });

  useEffect(() => {
    if (!IsNull(data)) {
      return;
    }

    /**
     * 初期値をセット
     */
    mutate((data) => {
      return data ?? initialData;
    }, false).catch((e: unknown) => {
      // TODO: #437 フロント用共通エラー処理関数を作る
      // eslint-disable-next-line no-console
      console.error(e);
    });
  }, [data, initialData, mutate]);

  return { data, mutate };
};

export default useStaticSWR;
