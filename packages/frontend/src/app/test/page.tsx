import Link from "next/link";

export default function Home() {
  return (
    <div className="flex w-full flex-col items-center">
      <Link href="./test/orval" prefetch={false}>
        orval
      </Link>
    </div>
  );
}
