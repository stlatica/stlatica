import { UserIconSampleImage } from "@/components/common/UserIcon/UserIconSampleImage";
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
      userIcon={UserIconSampleImage}
      userId="test"
      userName="テストユーザー"
    />
  );
};
