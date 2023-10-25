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
            <svg
              width="300"
              height="80"
              viewBox="0 0 480 129"
              fill="none"
              xmlns="http://www.w3.org/2000/svg">
              <circle cx="56" cy="73" r="56" fill="#141313" />
              <path
                d="M27.4443 102.054C10.9521 99.7737 8.049 79.8086 47.6811 32C40.4782 46.2514 30.8748 79.6499 36.9196 82.1202C38.1128 82.4049 46.7459 83.4491 55.9405 74.5262C56.9933 72.4378 43.9364 104.335 27.4443 102.054Z"
                fill="white"
              />
              <path
                d="M84.6518 46.8729C100.99 49.3464 103.739 69.3279 64.1319 116.626C71.3663 102.472 81.1051 69.2168 75.1273 66.6773C73.946 66.3787 65.3926 65.2334 56.2178 74.0397C55.1603 76.1137 68.3137 44.3994 84.6518 46.8729Z"
                fill="white"
              />
              <path
                d="M164.304 103.96C159.248 103.96 154.672 102.936 150.576 100.888C146.544 98.84 143.248 95.896 140.688 92.056L148.656 84.376C150.768 86.872 153.328 88.92 156.336 90.52C159.344 92.12 162.16 92.92 164.784 92.92C167.728 92.92 170.096 92.248 171.888 90.904C173.68 89.56 174.576 87.8 174.576 85.624C174.576 84.024 174.288 82.712 173.712 81.688C173.136 80.664 172.048 79.64 170.448 78.616C168.848 77.528 166.224 76.088 162.576 74.296C157.776 71.928 154 69.72 151.248 67.672C148.496 65.624 146.448 63.416 145.104 61.048C143.824 58.616 143.184 55.8 143.184 52.6C143.184 49.08 144.112 45.976 145.968 43.288C147.824 40.6 150.448 38.52 153.84 37.048C157.232 35.576 161.136 34.84 165.552 34.84C170.416 34.84 174.704 35.768 178.416 37.624C182.192 39.48 184.944 42.04 186.672 45.304L177.936 51.832C174.736 47.864 170.8 45.88 166.128 45.88C163.184 45.88 160.912 46.456 159.312 47.608C157.712 48.76 156.912 50.392 156.912 52.504C156.912 53.784 157.296 54.936 158.064 55.96C158.896 56.984 160.368 58.104 162.48 59.32C164.592 60.536 167.76 62.136 171.984 64.12C176.08 66.04 179.28 67.96 181.584 69.88C183.952 71.8 185.648 73.944 186.672 76.312C187.696 78.68 188.208 81.496 188.208 84.76C188.208 88.6 187.216 91.992 185.232 94.936C183.312 97.816 180.528 100.056 176.88 101.656C173.296 103.192 169.104 103.96 164.304 103.96ZM207.291 103.96C204.155 103.96 201.339 103.32 198.843 102.04C196.347 100.696 194.395 98.872 192.987 96.568C191.579 94.2 190.875 91.576 190.875 88.696C190.875 83.768 192.827 79.992 196.731 77.368C200.635 74.744 206.203 73.432 213.435 73.432H220.923V71.416C220.923 68.856 220.155 66.904 218.619 65.56C217.147 64.152 215.035 63.448 212.283 63.448C209.403 63.448 206.651 63.96 204.027 64.984C201.403 66.008 199.259 67.416 197.595 69.208L192.219 62.104C194.843 59.608 197.947 57.656 201.531 56.248C205.115 54.776 208.827 54.04 212.667 54.04C219.067 54.04 224.091 55.544 227.739 58.552C231.387 61.56 233.211 65.72 233.211 71.032V92.92C233.211 94.392 233.307 96.12 233.499 98.104C233.755 100.088 234.043 101.72 234.363 103H223.515C223.131 102.168 222.843 101.368 222.651 100.6C222.459 99.768 222.299 98.648 222.171 97.24H221.979C218.331 101.72 213.435 103.96 207.291 103.96ZM210.267 94.744C213.339 94.744 215.867 93.976 217.851 92.44C219.899 90.84 220.923 88.824 220.923 86.392V81.592H213.051C209.787 81.592 207.259 82.2 205.467 83.416C203.739 84.632 202.875 86.392 202.875 88.696C202.875 90.616 203.515 92.12 204.795 93.208C206.139 94.232 207.963 94.744 210.267 94.744ZM283.159 55L264.824 105.208C262.839 110.584 260.087 114.616 256.567 117.304C253.111 120.056 248.951 121.432 244.087 121.432C241.527 121.432 238.903 121.08 236.215 120.376V110.68C236.983 110.936 237.943 111.16 239.095 111.352C240.247 111.544 241.271 111.64 242.167 111.64C244.983 111.64 247.383 110.648 249.367 108.664C251.415 106.744 253.239 103.672 254.839 99.448L236.983 55H250.039L260.599 85.72L270.103 55H283.159ZM318.572 103L304.364 80.632L298.604 86.296V103H286.316V31.768H298.604V71.608L315.5 55H330.572L312.812 72.376L332.396 103H318.572ZM345.276 103.96C342.14 103.96 339.324 103.32 336.828 102.04C334.332 100.696 332.38 98.872 330.972 96.568C329.564 94.2 328.86 91.576 328.86 88.696C328.86 83.768 330.812 79.992 334.716 77.368C338.62 74.744 344.188 73.432 351.42 73.432H358.908V71.416C358.908 68.856 358.14 66.904 356.604 65.56C355.132 64.152 353.02 63.448 350.268 63.448C347.388 63.448 344.636 63.96 342.012 64.984C339.388 66.008 337.244 67.416 335.58 69.208L330.204 62.104C332.828 59.608 335.932 57.656 339.516 56.248C343.1 54.776 346.812 54.04 350.652 54.04C357.052 54.04 362.076 55.544 365.724 58.552C369.372 61.56 371.196 65.72 371.196 71.032V92.92C371.196 94.392 371.292 96.12 371.484 98.104C371.74 100.088 372.028 101.72 372.348 103H361.5C361.116 102.168 360.828 101.368 360.636 100.6C360.444 99.768 360.284 98.648 360.156 97.24H359.964C356.316 101.72 351.42 103.96 345.276 103.96ZM348.252 94.744C351.324 94.744 353.852 93.976 355.836 92.44C357.884 90.84 358.908 88.824 358.908 86.392V81.592H351.036C347.772 81.592 345.244 82.2 343.452 83.416C341.724 84.632 340.86 86.392 340.86 88.696C340.86 90.616 341.5 92.12 342.78 93.208C344.124 94.232 345.948 94.744 348.252 94.744Z"
                fill="black"
              />
            </svg>
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
