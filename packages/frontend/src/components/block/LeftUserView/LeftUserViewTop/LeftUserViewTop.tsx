import type React from "react";

import { BackButton } from "@/components/button/BackButton";
import { FollowButton } from "@/components/button/FollowButton";

// type LeftUserViewTopProps = {
//   // children: React.ReactNode;
// };

/**
 *
 */
export const LeftUserViewTop: React.FC = () => {
  return (
    <div className="flex flex-col">
      <div className="grid h-[170px] grid-cols-10 ">
        <div className="col-span-2 ">
          <div className="h-full">
            <a href="/">
              <BackButton />
            </a>
          </div>
        </div>
        <div className="relative col-span-4 ">
          <div className="absolute w-full border-2 border-black bg-red-100 pb-[100%]" />
        </div>
        <div className="col-span-4 m-2 flex w-full flex-col">
          <div className="flex w-full justify-center pb-3">
            <div>
              <FollowButton isFollow={true} />
            </div>
          </div>
          <div className="pb-2" />
          <div>9999万 Follows</div>
          <div>9999 Follers</div>
        </div>
      </div>
    </div>
  );
};
