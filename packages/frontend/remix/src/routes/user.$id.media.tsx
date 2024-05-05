import { LoaderFunctionArgs, redirect } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";

import { Timeline } from "@/components/timelines/Timeline";

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
      <Timeline />
    </>
  );
}
