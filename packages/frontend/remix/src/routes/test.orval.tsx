/* eslint-disable no-console */
import { useGetUser } from "@/openapi/stlaticaInternalApi";

export default function Home() {
  const { data, isLoading, error } = useGetUser("hoge");

  if (error) {
    console.log(data);
  }
  return (
    <div className="flex w-[90vw] flex-col">
      <div>useGetUser mock</div>
      <div>isLoading: {String(isLoading)}</div>
      <div>error: {JSON.stringify(error, null, 2)}</div>
      <div
        // eslint-disable-next-line react/forbid-dom-props
        style={{ lineBreak: "anywhere" }}
      >
        data: {JSON.stringify(data?.data, null, 2)}
      </div>
    </div>
  );
}
