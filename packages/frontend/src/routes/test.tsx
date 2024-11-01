import { type MetaFunction, Outlet } from "@remix-run/react";

import { flexContainer, sidebar } from "@/styles/routes/test.css";

export const meta: MetaFunction = () => {
  return [{ title: "test page" }];
};

export default function Home() {
  return (
    <>
      <div class={flexContainer}>
        <div class={sidebar}>
          <div>
            <a href="/test/orval">orval</a>
          </div>
          <div>
            <a href="/test/crop_icon">crop-icon</a>
          </div>
        </div>
        <Outlet />
      </div>
    </>
  );
}
