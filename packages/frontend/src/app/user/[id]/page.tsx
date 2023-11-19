import { LeftUserView } from "@/components/block/LeftUserView";

import { DummyTimeLines } from "./DummyTimeLines";

export default function Page({ params }: { readonly params: { id: string } }) {
  const { id } = params;

  // console.log(id);

  return (
    <main className="h-[100vh]">
      <div className="flex h-full">
        <div className="min-w-[350px] text-slate-900">
          <LeftUserView userID={id} />
        </div>
        <DummyTimeLines />
      </div>
    </main>
  );
}
