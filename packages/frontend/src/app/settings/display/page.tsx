import Link from "next/link";

export default function Display() {
  return (
    <div className="flex">
      <div className="flex flex-col">
        <Link href="/settings/muted_user">ミュートしたアカウント</Link>
        <Link href="/settings/font_size">フォントサイズ</Link>
        <Link href="/settings/color">カラーテーマ</Link>
      </div>
    </div>
  );
}
