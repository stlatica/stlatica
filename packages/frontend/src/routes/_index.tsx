import { container } from "@/styles/routes/index.css";

export default function Home() {
  return (
    <main>
      <div class={container}>
        <a href="/edit_profile">EditProfile</a>
        <a href="/user/sample_user">User Page</a>
        <a href="/test">実験用ページ</a>
        <a href="/login">ログイン画面</a>
        <a href="/settings">設定画面(サンプル)</a>
        <a href="/post">plat投稿画面(試作版)</a>
      </div>
    </main>
  );
}
