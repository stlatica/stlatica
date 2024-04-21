import { ColorSchemeScript, MantineProvider } from "@mantine/core";
import { LinksFunction, LoaderFunction, redirect } from "@remix-run/node";
import { Links, Meta, Outlet, Scripts, ScrollRestoration } from "@remix-run/react";

import "@mantine/core/styles.css";
// for tailwind
import stylesheet from "./routes/global.css?url";

export const links: LinksFunction = () => {
  return [
    // for tailwind
    { rel: "stylesheet", href: stylesheet },
    // cropperjs css from cdn
    {
      rel: "stylesheet",
      href: "https://cdnjs.cloudflare.com/ajax/libs/cropperjs/1.6.1/cropper.min.css",
    },
  ];
};

// トレイリングスラッシュ除去
export const loader = (({ request }) => {
  const { pathname, search } = new URL(request.url);
  // https://github.com/remix-run/remix/discussions/6584
  if (pathname.endsWith("/") && pathname !== "/") {
    // Redirect to the same URL without a trailing slash
    return redirect(`${pathname.slice(0, -1)}${search}`, 301);
  }

  return null;
}) satisfies LoaderFunction;

export const Layout = ({ children }: { readonly children: React.ReactNode }) => {
  return (
    <html lang="en">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <Meta />
        <Links />
        <ColorSchemeScript defaultColorScheme="dark" />
      </head>
      <body>
        <MantineProvider defaultColorScheme="dark">{children}</MantineProvider>
        <ScrollRestoration />
        <Scripts />
      </body>
    </html>
  );
};

const App = () => {
  return <Outlet />;
};

export default App;
