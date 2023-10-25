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

  // const fontData = await font;
  return new ImageResponse(
    (
      <div
        style={{
          fontSize: 80,
          backgroundImage: `url(${`data:image/svg+xml,${encodeURIComponent(
            '<svg id="visual" viewBox="0 0 900 600" width="900" height="600" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1"><defs><filter id="blur1" x="-10%" y="-10%" width="120%" height="120%"><feFlood flood-opacity="0" result="BackgroundImageFix"></feFlood><feBlend mode="normal" in="SourceGraphic" in2="BackgroundImageFix" result="shape"></feBlend><feGaussianBlur stdDeviation="161" result="effect1_foregroundBlur"></feGaussianBlur></filter></defs><rect width="900" height="600" fill="#bcdaff"></rect><g filter="url(#blur1)"><circle cx="140" cy="72" fill="#46ffd2" r="357"></circle><circle cx="424" cy="163" fill="#bcdaff" r="357"></circle><circle cx="18" cy="468" fill="#46ffd2" r="357"></circle><circle cx="597" cy="68" fill="#46ffd2" r="357"></circle><circle cx="782" cy="590" fill="#bcdaff" r="357"></circle><circle cx="420" cy="487" fill="#46ffd2" r="357"></circle></g></svg>'
          )}`})`,
          backgroundRepeat: "no-repeat",
          backgroundPosition: "center",
          backgroundSize: "100% 100%",
          width: "100%",
          height: "100%",
          display: "flex",
          position: "relative",
        }}>
        <div
          style={{
            backgroundColor: "white",
            borderRadius: 20,
            display: "flex",
            width: "95%",
            height: "90%",
            margin: "auto",
            boxShadow: "10px 10px 10px rgba(0, 0, 0, 0.25)",
          }}>
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
          <div
            style={{
              display: "flex",
              position: "absolute",
              width: "100%",
              bottom: 40,
              padding: "0px 40px 0px 50px",
              justifyContent: "space-between",
            }}>
            <div
              style={{
                display: "flex",
                alignItems: "center",
              }}>
              <img
                src={`https://github.com/${userName}.png`}
                alt=""
                width="60"
                height="60"
                style={{ borderRadius: 60, marginRight: 20 }}
              />
              <span
                style={{
                  color: "black",
                  fontSize: 60,
                }}>
                {userName}
              </span>
            </div>
          </div>
        </div>
      </div>
    ),
    {
      width: 1200,
      height: 630,
      // fonts: [
      //   {
      //     name: "NotoSansJP",
      //     data: fontData,
      //     style: "normal",
      //   },
      // ],
    }
  );
}
