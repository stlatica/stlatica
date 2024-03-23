import Link from "next/link";

export default function Account() {
  return (
    <div className="flex">
      <div className="flex flex-col">
        <Link href="/settings/user_id">ユーザIDの変更</Link>
        <Link href="/settings/email">メールアドレスの変更</Link>
        <Link href="/settings/password">パスワードの変更</Link>
        <Link href="/settings/audience">Platの非公開設定</Link>
      </div>
    </div>
  );
}
