import React from "react";

import { LoginButton } from "@/components/button/LoginButton";
import { TextEditor } from "@/components/common/TextEditor";

export default function LoginScene() {
  return (
    <main className="flex h-[100vh] w-[100vw] justify-center bg-gray-800">
      <div className="flex w-[50%] min-w-[400px] max-w-[600px] flex-col">
        <div className="mt-[50%] space-y-6">
          <div className="flex justify-center text-5xl">Stlatica</div>
          <div className="bg-gray-800 text-white">
            <TextEditor defaultValue="" maxlength={32} title="mail address" />
          </div>
          <div className="bg-gray-800 text-white">
            <TextEditor defaultValue="" maxlength={32} title="password" />
          </div>
          <div className="flex justify-center">
            <LoginButton />
          </div>
        </div>
      </div>
    </main>
  );
}
