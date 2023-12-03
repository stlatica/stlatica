"use client";
import { useGetUser } from "@/openapi/user/user";

export default function Home() {
  const { data, isLoading, error } = useGetUser("hoge");

  // console.log(data);
  return (
    <div className="flex w-full flex-col items-center">
      <div>useGetUser mock</div>
      <div>isLoading: {String(isLoading)}</div>
      <div>error: {JSON.stringify(error)}</div>
      <div>data: {JSON.stringify(data?.data)}</div>
    </div>
  );
}
