import Link from "next/link";

export default function Home() {
  return (
    <main>
      <div className="flex flex-col">
        <Link href="/sample/zod">zod sample</Link>
        <Link href="/user/sample_user">User Page</Link>
      </div>
    </main>
  );
}
