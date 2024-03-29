import Link from "next/link";

export default function Home() {
  return (
    <main>
      <div className="flex flex-col">
        <Link href="/sample/zod">zod sample</Link>
        <Link href="/EditProfile">EditProfile</Link>
        <Link href="/user/sample_user">User Page</Link>
        <Link href="/test">実験用ページ</Link>
        <Link href="/test/CropIcon">実験用ページ / アイコン画像のトリミング</Link>
        <Link href="/login">ログイン画面</Link>
      </div>
    </main>
  );
}
