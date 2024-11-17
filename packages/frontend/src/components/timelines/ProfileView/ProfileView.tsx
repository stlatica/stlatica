import type React from "react";

import { container, textSmall } from "./ProfileView.css";

type ProfileViewProps = {
  readonly userID: string;
  // children: React.ReactNode;
};

/**
 *
 */
export const ProfileView: React.FC<ProfileViewProps> = ({ userID }) => {
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
