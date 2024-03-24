import Link from "next/link";

export default function Account() {
  return (
    <div className="flex">
      <div className="flex flex-col">
        <Link href="/settings/account/user_id">ユーザIDの変更</Link>
        <Link href="/settings/account/email">メールアドレスの変更</Link>
        <Link href="/settings/account/password">パスワードの変更</Link>
        <Link href="/settings/account/audience">Platの非公開設定</Link>
      </div>
    </div>
  );
}
