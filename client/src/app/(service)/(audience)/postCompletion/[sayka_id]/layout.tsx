import { Metadata } from "next";

// export const metadata: Metadata = {
//   title: "Sayka",
//   description: "成果物を投稿するアプリケーション",
//   openGraph: {
//     title: "Sayka",
//     description: "成果物を投稿するアプリケーション",
//     images: "http://localhost:3000/_next/image?url=%2Ficon.jpg&w=64&q=75",
//   },
// };

export default function Layout(props: { children: React.ReactNode }) {
  return (
    <main className="grid grid-cols-10 p-5">
      <div className="hidden md:col-span-2 md:block" />
      <div className="col-span-10 md:col-span-6">{props.children}</div>
      <div className="hidden md:col-span-2 md:block" />
    </main>
  );
}
