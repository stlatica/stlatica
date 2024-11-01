import type React from "react";

import { container, textSmall } from "./ProfileView.css";

type LeftUserViewProps = {
  readonly userID: string;
  // children: React.ReactNode;
};

/**
 *
 */
export const LeftUserView: React.FC<LeftUserViewProps> = ({ userID }) => {
  return (
    <div class={container}>
      <div>User Name Hereeeeeeeee</div>
      <div>@{userID}</div>
      <div>
        <hr />
      </div>
      <div class={textSmall}>
        ここに自己紹介をかきまああああああああああああああああああああああああああああああああああああす文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数
      </div>
    </div>
  );
};
