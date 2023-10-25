import { ImageResponse } from "next/server";

const font = fetch(
  new URL("../../assets/NotoSansJP-Bold.otf", import.meta.url)
).then((res) => res.arrayBuffer());

export const runtime = "edge";

export async function GET(request: Request) {
  const { searchParams } = new URL(request.url);
  const hasTitle = searchParams.has("title");
  const title = hasTitle ? searchParams.get("title")?.slice(0, 100) : "Sayka";
  const hasUserName = searchParams.has("userName");
  const userName = hasUserName ? searchParams.get("userName") : "unknown";

  const fontData = await font;
  return new ImageResponse(
    (
      <div
        style={{
          width: "100%",
          color: "black",
          fontSize: 70,
          display: "flex",
          padding: "40px 40px 0px 40px",
        }}>
        {title}
      </div>
    ),
    {
      width: 1200,
      height: 630,
      fonts: [
        {
          name: "NotoSansJP",
          data: fontData,
          style: "normal",
        },
      ],
    }
  );
}
