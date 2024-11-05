import type React from "react";

import { container } from "./LeftUserView.css";
import { LeftUserViewTop } from "./LeftUserViewTop";

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
      <LeftUserViewTop />
      <div className="text-xl">User Name Hereeeeeeeee</div>
      <div className="text-gray-400">@{userID}</div>
      <div className="flex justify-center ">
        <hr className="m-3 w-4/5 border-black" />
      </div>
      <div className="text-sm">
        ここに自己紹介をかきまああああああああああああああああああああああああああああああああああああす文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数
      </div>
    </div>
  );
};
