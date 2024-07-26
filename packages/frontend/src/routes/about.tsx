import { flexContainer, flexColumn } from "@/styles/routes/about.css";

export default function Page() {
  return (
    <>
      <div className={flexContainer}>
        <div className={flexColumn}>
          <a href="/">リリースノート</a>
          <a href="/">Cookieのポリシー</a>
          <a href="/">プライバシー</a>
          <a href="/">利用規約</a>
          <a href="/">Lisence</a>
        </div>
      </div>
    </>
  );
}
