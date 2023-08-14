import { Murecho } from "next/font/google";
import "./globals.css";

import { Providers } from "./Providers";

import type { Metadata } from "next";

const murecho = Murecho({ weight: "100", subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Sayka",
  description: "成果物を投稿するアプリケーションです",
  // manifest: "/manifest.json",
  openGraph: {
    title: "Sayka",
    description: "成果物を投稿するアプリケーション",
    url: "https://sayka.vercel.app/",
    siteName: "Sayka",
    images: [
      {
        url: "https://sayka.vercel.app/ogp.png",
        width: 700,
        height: 400,
      },
    ],
    type: "website",
  },
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body style={{ fontFamily: `${murecho}` }}>
        <Providers>{children}</Providers>
      </body>
    </html>
  );
}
