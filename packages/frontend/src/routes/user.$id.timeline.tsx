import { type LoaderFunctionArgs, redirect } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";

import { ProfileTimeline } from "@/components/timelines/Timeline";

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
    <>
      <ProfileTimeline url={`user/${id}`} user_id={id} />
    </>
  );
}
