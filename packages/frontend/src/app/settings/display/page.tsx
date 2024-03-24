import Link from "next/link";

export default function Display() {
  return (
    <div className="flex">
      <div className="flex flex-col">
        <Link href="/settings/display/muted_user">ミュートしたアカウント</Link>
        <Link href="/settings/display/font_size">フォントサイズ</Link>
        <Link href="/settings/display/color">カラーテーマ</Link>
      </div>
    </div>
  );
}
