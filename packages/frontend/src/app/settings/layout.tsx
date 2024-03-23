import Link from "next/link";

export default function SettingsLayout({ children }: { readonly children: React.ReactNode }) {
  return (
    <div className="flex">
      <div className="flex w-40 flex-col">
        <Link href="/settings/account">アカウント</Link>
        <Link href="/settings/display">表示</Link>
        <Link href="/settings/about">Stlaticaについて</Link>
      </div>
      <div>{children}</div>
    </div>
  );
}
