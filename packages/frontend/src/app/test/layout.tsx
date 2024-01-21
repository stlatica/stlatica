import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "Stlatica 開発用ページ",
};

export default function RootLayout({ children }: { readonly children: React.ReactNode }) {
  return (
    <html lang="ja">
      <body>{children}</body>
    </html>
  );
}
