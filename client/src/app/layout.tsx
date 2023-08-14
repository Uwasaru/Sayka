import { Murecho } from "next/font/google";
import "./globals.css";

import { Providers } from "./Providers";

import type { Metadata } from "next";

const murecho = Murecho({ weight: "100", subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Sayka",
  description: "成果物を投稿するアプリケーションですby39TO",
  openGraph: {
    title: "Saykaです",
    description: "成果物を投稿するアプリケーション",
    url: "https://sayka.vercel.app/",
    siteName: "Sayka",
    images: [
      {
        url: "https://sayka.vercel.app/_next/image?url=%2Fogp.png&w=64&q=75",
        width: 800,
        height: 600,
      },
    ],
    locale: "en_US",
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
