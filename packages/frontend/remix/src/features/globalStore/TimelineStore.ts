import { mutate } from "swr";

import { SimpleSortObject } from "@/utilities/utilities";

import useStaticSWR from "./useStaticSWR";

import type { Plat } from "@/openapi/stlaticaInternalApi.schemas";

const PlatURLKey = (plat_id: string) => {
  return `http://localhost:4010/internal/v1/plats/${plat_id}`;
};

/**
 * useGetPlat のキャッシュを手動で作成する
 */
export const SetPlatCache = (plats: Plat[]) => {
  plats.forEach((plat) => {
    // console.log("set cache: ", PlatURLKey(plat.plat_id));

    // plat-{uuid} の形式で SWR に保存する
    mutate([PlatURLKey(plat.plat_id)], { data: plat }).catch((e: unknown) => {
      // TODO: #437 フロント用共通エラー処理関数を作る
      // eslint-disable-next-line no-console
      console.error(e);
    });
  });
};

type TimelineType = { plat_id: string; created_at: Date };

/**
 * 各タイムラインで利用するフック
 */
export const useTimeline = (url: string) => {
  const { data, mutate: _mutate } = useStaticSWR(`timeline-${url}`, new Array<TimelineType>());

  const mutate = (plats: Plat[]) => {
    SetPlatCache(plats);

    // データ整形
    const r: TimelineType[] = plats.map((x) => {
      return { plat_id: x.plat_id, created_at: new Date(x.created_at) };
    });

    // 新しいデータを合体しつつソート
    const newData = r.concat(data).sort(SimpleSortObject("created_at"));

    _mutate(newData).catch((e: unknown) => {
      // TODO: #437 フロント用共通エラー処理関数を作る
      // eslint-disable-next-line no-console
      console.error(e);
    });
  };

  return {
    data,

    /**
     * platデータをグローバルストアに保管しつつ、このタイムラインを更新する
     */
    mutate,
  };
};

// /**
//  * platの情報を取得する
//  */
// export const useGetPlatFromStore = (platID: string) => {
//   const ret = useSWR(PlatPrefixedKey(platID), {
//     revalidateOnFocus: false,
//     revalidateOnMount: false,
//     revalidateOnReconnect: false,
//     revalidateIfStale: false,
//   });

//   return ret;
// };
