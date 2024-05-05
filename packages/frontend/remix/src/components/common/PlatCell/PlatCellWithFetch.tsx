import { useGetPlat } from "@/openapi/stlaticaInternalApi";

import { PlatCell } from "./PlatCell";

type Props = {
  id: string;
};

export const PlatCellWithFetch: React.FC<Props> = ({ id }) => {
  const { data } = useGetPlat(id, {
    swr: {
      // revalidateOnFocus: false,
      // revalidateOnMount: false,
      // revalidateOnReconnect: false,
      // revalidateIfStale: false,
    },
  });

  if (data === undefined) {
    return <></>;
  }

  return (
    <PlatCell
      content={data.data.content}
      favoriteCount={20}
      replyCount={1}
      shareCount={10}
      userId="test"
      userName="テストユーザー"
    />
  );
};
