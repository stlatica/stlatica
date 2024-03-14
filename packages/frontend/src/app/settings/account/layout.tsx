import Link from "next/link";

export default function AccountLayout({ children }: { readonly children: React.ReactNode }) {
  return (
    <div className="flex">
      <div className="flex w-40 flex-col">
        <Link href="/settings/account/UserId">ユーザID</Link>
        <Link href="/settings/account/Email">メールアドレス</Link>
        <Link href="/settings/account/password">パスワード</Link>
      </div>
      <div>{children}</div>
    </div>
  );
}
