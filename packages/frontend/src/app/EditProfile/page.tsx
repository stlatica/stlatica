import React from "react";

import { CancelButton } from "@/components/button/CancelButton";
import { SubmitButton } from "@/components/button/SubmitButton";
import { TextEditor } from "@/components/common/TextEditor";

import { IconEditor } from "./IconEditor";

export default function Home() {
  return (
    <main>
      <div className="flex bg-gray-800 p-12 text-white">
        <div className="w-1/2">
          <IconEditor />
        </div>
        <div className="ml-12 w-1/2">
          <div>
            <TextEditor componentId="Name" defaultValue="" maxlength={16} title="Name" />
          </div>
          <div className="mt-12">
            <TextEditor
              componentId="SI"
              defaultValue=""
              maxlength={512}
              title="Self-Intorduction"
            />
          </div>
          <div className="mt-12 flex justify-around">
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
