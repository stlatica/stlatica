import React from "react";

import { LoginButton } from "@/components/button/LoginButton";

type LoginSceneProps = {
  // children: React.ReactNode;
};

/**
 *
 */
export const LoginScene: React.FC<LoginSceneProps> = () => {
  return (
    <div className="grid h-full w-full place-content-center space-y-5 bg-gray-200">
      <div className="text-5xl">Stlatica</div>
      {/* <div className="">
        <TextEditor defaultValue="" maxlength={32} title="mail address" />
      </div> */}
      {/* <div className="">
        <TextEditor defaultValue="" maxlength={32} title="password" />
      </div> */}
      <div className="">
        <LoginButton
          onClick={() => {
            console.log("clicked login button");
          }}
        />
      </div>
    </div>
  );
};
