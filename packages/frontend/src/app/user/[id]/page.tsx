const Page = ({ params }: { readonly params: { id: string } }) => {
  console.log(params.id);

  return (
    <main>
      <div className="flex bg-gray-600">
        <div className="w-[400px] bg-gray-100 text-slate-900">
          <div>{params.id}</div>
        </div>
        <div className="h-full bg-red-100 text-slate-900">タイムラインをここに</div>
      </div>
    </main>
  );
};

export default Page;
