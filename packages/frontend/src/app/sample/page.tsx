import Link from "next/link";
import React from "react";

import { CancelButton } from "@/components/button/CancelButton";

import { X } from "./X";

const TimeLine: React.FC = () => {
  const ar = new Array(100).fill(0).map((_, i) => {
    return i;
  });

  const list = ar.map((x) => {
    return (
      <div key={x} style={{ marginBottom: "0.5em" }}>
        <X />
      </div>
    );
  });

  return (
    <>
      <Link href="/">
        <CancelButton>トップへもどる</CancelButton>
      </Link>
      <div>{list}</div>
    </>
  );
};

export default function Home() {
  return (
    <main>
      <TimeLine />
    </main>
  );
}
