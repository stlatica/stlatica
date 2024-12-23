import { flexColumn, flexContainer } from "@/styles/routes/settings.css";

export default function Page() {
  return (
    <div className={flexContainer}>
      <div className={flexColumn}>
        <a href="/settings/account/user_id">ユーザIDの変更</a>
        <a href="/settings/account/email">メールアドレスの変更</a>
        <a href="/settings/account/password">パスワードの変更</a>
        <a href="/settings/account/audience">Platの非公開設定</a>
      </div>
    </div>
  );
}
