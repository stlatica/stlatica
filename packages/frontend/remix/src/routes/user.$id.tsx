import { LoaderFunctionArgs, redirect } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";

import { LeftUserView } from "@/components/block/LeftUserView";
import { ProfileTab } from "@/components/timelines/ProfileTab";

export const loader = ({ params }: LoaderFunctionArgs) => {
  const { id } = params;

  if (id === undefined) {
    return redirect("/");
  }

  return { id };
};

export default function Page() {
  const { id } = useLoaderData<typeof loader>();

  return (
    <main className="h-screen">
      <div className="flex h-full">
        <div className="w-[330px] text-slate-900">
          <LeftUserView userID={id} />
        </div>
        <div className="pr-2" />
        <div className="w-[600px]">
          <ProfileTab />
        </div>
      </div>
    </main>
  );
}
