import "./global.css";
import { ThemeProvider } from "@mui/material/styles";
import { AppRouterCacheProvider } from "@mui/material-nextjs/v13-appRouter";

import { muiTheme } from "./providers";

export default function RootLayout({ children }: { readonly children: React.ReactNode }) {
  return (
    <html lang="ja">
      <body>
        {/* https://mui.com/material-ui/integrations/nextjs/ */}
        <AppRouterCacheProvider>
          <ThemeProvider theme={muiTheme}>{children}</ThemeProvider>
        </AppRouterCacheProvider>
      </body>
    </html>
  );
}
