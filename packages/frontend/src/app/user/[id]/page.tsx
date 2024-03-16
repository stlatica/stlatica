import { LeftUserView } from "@/components/block/LeftUserView";
import { ProfileTab } from "@/components/timelines/ProfileTab";

export default function Page({ params }: { readonly params: { id: string } }) {
  const { id } = params;

  // console.log(id);

  return (
    <main className="h-[100vh]">
      <div className="flex h-full">
        <div className="w-[350px] text-slate-900">
          <LeftUserView userID={id} />
        </div>
        <div className="pr-2" />
        <div className="w-[600px]">
          <ProfileTab />
          di
        </div>
      </div>
    </main>
  );
}
