import React from "react";

import { CancelButton } from "@/components/button/CancelButton";
import { SubmitButton } from "@/components/button/SubmitButton";

import { IconEditor } from "./IconEditor";
import { NameEditor } from "./NameEditor";
import { SelfIntroductionEditor } from "./SelfIntroductionEditor";

export default function Home() {
  return (
    <main>
      <div className="flex p-12 text-white bg-gray-800">
        <div className="w-1/2">
          <IconEditor />
        </div>
        <div className="w-1/2 ml-12">
          <div>
            <NameEditor />
          </div>
          <div className="mt-12">
            <SelfIntroductionEditor />
          </div>
          <div className="flex justify-around mt-12">
            {/* Assuming SubmitButton and CancelButton are styled components, you would need to pass the class names instead of the style prop. */}
            <SubmitButton className="w-[85px]">Save</SubmitButton>{" "}
            {/* Adjust the width as needed */}
            <CancelButton className="w-[85px]">Cancel</CancelButton>{" "}
            {/* Adjust the width as needed */}
          </div>
        </div>
      </div>
    </main>
  );
}
