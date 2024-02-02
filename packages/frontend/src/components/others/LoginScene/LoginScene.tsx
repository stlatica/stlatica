import React from "react";

import { LoginButton } from "@/components/button/LoginButton";
import { TextEditor } from "@/components/common/TextEditor";

type LoginSceneProps = {
  // children: React.ReactNode;
};

/**
 *
 */
export const LoginScene: React.FC<LoginSceneProps> = () => {
  return (
    <div className="grid h-full w-full place-content-center space-y-3 bg-gray-800">
      <div className="flex justify-center text-5xl">Stlatica</div>
      <div className="h-24 w-80 bg-gray-800 p-6 text-white">
        <TextEditor defaultValue="" maxlength={32} title="mail address" />
      </div>
      <div className="h-24 w-80 bg-gray-800 p-6 text-white">
        <TextEditor defaultValue="" maxlength={32} title="password" />
      </div>
      <div className="flex justify-center">
        <LoginButton
          onClick={() => {
            console.log("clicked login button");
          }}
        />
      </div>
    </div>
  );
};
