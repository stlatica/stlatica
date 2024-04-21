export default function Home() {
  return (
    <main>
      <div className="flex flex-col">
        <a href="/edit_profile">EditProfile</a>
        <a href="/user/sample_user">User Page</a>
        <a href="/test">実験用ページ</a>
        <a href="/login">ログイン画面</a>
        <a href="/settings">設定画面(サンプル)</a>
      </div>
    </main>
  );
}
