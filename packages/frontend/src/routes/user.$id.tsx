import { Stack } from "@mantine/core";
import { LoaderFunctionArgs, redirect } from "@remix-run/node";
import { Outlet, useLoaderData } from "@remix-run/react";

import { LeftUserView } from "@/components/block/LeftUserView";
import { ProfileTab } from "@/components/timelines/ProfileTab";
import { mainContainer, flexContainer, leftPanel, rightPadding } from "@/styles/routes/user.id.css";

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
        <div className={leftPanel}>
          <LeftUserView userID={id} />
        </div>
        <div className={rightPadding} />
        <Stack>
          <ProfileTab />
          <Outlet />
        </Stack>
      </div>
    </main>
  );
}
