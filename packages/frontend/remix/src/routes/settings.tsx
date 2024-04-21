import { LoaderFunctionArgs, redirect } from "@remix-run/node";
import { Outlet } from "@remix-run/react";

export const loader = ({ request }: LoaderFunctionArgs) => {
  const url = new URL(request.url);

  // このページを直接開いていたらリダイレクトする
  if (url.pathname === "/settings") {
    return redirect("/settings/account");
  }

  return null;
};

export default function Page() {
  return (
    <div className="flex">
      <div className="flex w-40 flex-col">
        <a href="/settings/account">アカウント</a>
        <a href="/settings/display">表示</a>
        <a href="/about">Stlaticaについて</a>
      </div>
      <div>
        <Outlet />
      </div>
    </div>
  );
}
