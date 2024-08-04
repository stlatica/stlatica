import { SampleImageBrown400 } from "@/features/sample/sample-image/SampleImageBrown400";
import { useGetPlat } from "@/openapi/stlaticaInternalApi";

import { PlatCell } from "./PlatCell";

type Props = {
  id: string;
};

export const PlatCellWithFetch: React.FC<Props> = ({ id }) => {
  const { data } = useGetPlat(id, {
    // 認証ヘッダー仮置き
    axios: {
      headers: {
        Authorization: "Bearer aaaa",
      },
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
      date={data.data.created_at}
      userIcon={SampleImageBrown400}
    />
  );
};
