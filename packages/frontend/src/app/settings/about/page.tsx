import Link from "next/link";

export default function About() {
  return (
    <div className="flex">
      <div className="flex w-40 flex-col">
        <Link href="/">リリースノート</Link>
        <Link href="/">Cookieのポリシー</Link>
        <Link href="/">プライバシー</Link>
        <Link href="/">利用規約</Link>
        <Link href="/">Lisence</Link>
      </div>
    </div>
  );
}
