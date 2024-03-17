import React from "react";

import { LoginButton } from "@/components/button/LoginButton";
import { TextEditor } from "@/components/common/TextEditor";

export default function LoginScene() {
  return (
    <main className="flex h-screen w-screen justify-center bg-gray-800">
      <div className="flex w-4/5 min-w-[300px] max-w-[800px] flex-col justify-center gap-6">
        <div className="flex justify-center text-5xl">Stlatica</div>
        <div className="bg-gray-800 text-white">
          <TextEditor defaultValue="" maxLength={32} title="mail address" />
        </div>
        <div className="bg-gray-800 text-white">
          <TextEditor defaultValue="" maxLength={32} title="password" />
        </div>
        <div className="flex justify-center">
          <LoginButton />
        </div>
      </div>
    </main>
  );
}

