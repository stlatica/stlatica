import React from "react";

import { CancelButton } from "@/components/button/CancelButton";
import { SubmitButton } from "@/components/button/SubmitButton";

import { IconEditor } from "./IconEditor";
import { NameEditor } from "./NameEditor";
import { SelfIntroductionEditor } from "./SelfIntroductionEditor";

export default function Home() {
  return (
    <main>
      <div
        style={{ display: "flex", padding: "48px", color: "#FFFFFF", backgroundColor: "#373737" }}
      >
        <div style={{ width: "50%" }}>
          <IconEditor />
        </div>
        <div style={{ width: "50%", marginLeft: "48px" }}>
          <div>
            <NameEditor />
          </div>
          <div style={{ marginTop: "48px" }}>
            <SelfIntroductionEditor />
          </div>
          <div style={{ marginTop: "48px", display: "flex", justifyContent: "space-around" }}>
            <SubmitButton style={{ width: "100px" }}>Save</SubmitButton>
            <CancelButton style={{ width: "100px" }}>Cancel</CancelButton>
          </div>
        </div>
      </div>
    </main>
  );
}
