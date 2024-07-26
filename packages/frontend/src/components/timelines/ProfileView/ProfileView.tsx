import React from "react";

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
    <div className={container}>
      <div>User Name Hereeeeeeeee</div>
      <div>@{userID}</div>
      <div>
        <hr />
      </div>
      <div className={textSmall}>
        ここに自己紹介をかきまああああああああああああああああああああああああああああああああああああす文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数
      </div>
    </div>
  );
};
