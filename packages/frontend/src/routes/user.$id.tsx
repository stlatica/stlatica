import { type LoaderFunctionArgs, redirect } from "@remix-run/node";
import { Outlet, useLoaderData } from "@remix-run/react";

import { ProfileTab } from "@/components/timelines/ProfileTab";
import { UserProfilePanel } from "@/components/timelines/UserProfilePanel/UserProfilePanel";
import { flexContainer, mainContainer } from "@/styles/routes/user.id.css";

export const loader = ({ params, request }: LoaderFunctionArgs) => {
  const { id } = params;

  if (id === undefined) {
    return redirect("/");
  }

  // このページを直接開いていたらリダイレクトする
  const lastPath = request.url.split("/").at(-1);
  if (lastPath === id) {
    return redirect(`${request.url}/timeline`);
  }

  return { id };
};

export default function Page() {
  const { id } = useLoaderData<typeof loader>();

  return (
    <main className={mainContainer}>
      <div className={flexContainer}>
        <UserProfilePanel userID={id} />
        <ProfileTab />
        {/* eslint-disable-next-line react/forbid-dom-props */}
        <div style={{ height: "100vh" }}>
          <Outlet />
        </div>
      </div>
    </main>
  );
}
