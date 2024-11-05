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
    <div class="flex flex-col">
      <div class="grid h-[170px] grid-cols-10 ">
        <div class="col-span-2 ">
          <div class="h-full">
            <a href="/">
              <BackButton />
            </a>
          </div>
        </div>
        <div class="relative col-span-4 ">
          <div class="absolute w-full border-2 border-black bg-red-100 pb-[100%]" />
        </div>
        <div class="col-span-4 m-2 flex w-full flex-col">
          <div class="flex w-full justify-center pb-3">
            <div>
              <FollowButton isFollow={true} />
            </div>
          </div>
          <div class="pb-2" />
          <div>9999万 Follows</div>
          <div>9999 Follers</div>
        </div>
      </div>
    </div>
  );
};
