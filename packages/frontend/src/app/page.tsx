import { Stack } from "@mui/material";
import Link from "next/link";

export default function Home() {
  return (
    <main>
      <Stack>
        <Link href="/sample">サンプルページへ</Link>
        <Link href="/sample/zod">zod sample</Link>
        <Link href="/user/sample_user">User Page</Link>
      </Stack>
    </main>
  );
}
