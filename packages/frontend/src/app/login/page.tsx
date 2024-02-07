import React from "react";

import { LoginButton } from "@/components/button/LoginButton";
import { TextEditor } from "@/components/common/TextEditor";

export default function LoginScene() {
  return (
    <main className="h-[100vh]">
      <div className="grid h-full w-full grid-cols-1 content-center space-y-6 bg-gray-800">
        <div className="flex justify-center text-5xl">Stlatica</div>
        <div className="mx-[25%] bg-gray-800 text-white">
          <TextEditor defaultValue="" maxlength={32} title="mail address" />
        </div>
        <div className="mx-[25%] bg-gray-800 text-white">
          <TextEditor defaultValue="" maxlength={32} title="password" />
        </div>
        <div className="flex justify-center">
          <LoginButton />
        </div>
      </div>
    </main>
  );
}
