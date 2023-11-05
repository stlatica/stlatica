import React from "react";

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
    <div className="h-full w-full bg-gray-200 p-3">
      <LeftUserViewTop />
      <div className="text-xl">User Name Hereeeeeeeee</div>
      <div className="text-gray-400">@{userID}</div>
      <div className="flex justify-center ">
        <hr className="m-3 w-[80%] border-black" />
      </div>
      <div className="text-sm">
        ここに自己紹介をかきまああああああああああああああああああああああああああああああああああああす文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数文字数
      </div>
    </div>
  );
};
