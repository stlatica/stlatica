import { MetaFunction, Outlet } from "@remix-run/react";

export const meta: MetaFunction = () => {
  return [{ title: "test page" }];
};

export default function Home() {
  return (
    // <div className="flex h-screen w-full ">
    <>
      {/* eslint-disable-next-line react/forbid-dom-props */}
      <div className="flex h-full flex-row gap-4">
        <div className="flex h-full w-[8em] flex-col gap-1">
          <div>
            <a href="/test/orval">orval</a>
          </div>
          <div>
            <a href="/test/crop_icon">crop-icon</a>
          </div>
        </div>
        {/* <div className="w-auto"> */}
        <Outlet />
        {/* </div> */}
      </div>
    </>
  );
}
