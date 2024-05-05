import { LoaderFunctionArgs, redirect } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";

export const loader = ({ params }: LoaderFunctionArgs) => {
  const { id } = params;

  if (id === undefined) {
    return redirect("/");
  }

  return { id };
};

export default function Page() {
  return (
    <>
      <div>未実装</div>
    </>
  );
}
