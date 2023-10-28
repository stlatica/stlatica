const Page = ({ params }: { readonly params: { id: string } }) => {
  console.log(params.id);

  return (
    <main>
      <div>
        <div>{params.id}</div>
      </div>
    </main>
  );
};

export default Page;
