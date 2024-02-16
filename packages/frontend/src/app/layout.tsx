import "@mantine/core/styles.css";

import "./global.css";

import { ColorSchemeScript, MantineProvider } from "@mantine/core";

export default function RootLayout({ children }: { readonly children: React.ReactNode }) {
  return (
    <html lang="ja">
      <head>
        <ColorSchemeScript defaultColorScheme="dark" />
      </head>
      <body>
        <MantineProvider defaultColorScheme="dark">{children}</MantineProvider>
      </body>
    </html>
  );
}
