import { Murecho } from "next/font/google";
import "./globals.css";

import { Providers } from "./Providers";

import type { Metadata } from "next";

const murecho = Murecho({ weight: "100", subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Sayka",
  description: "成果物を投稿するアプリケーションです",
  manifest: "https://sayka.site/manifest.json",
  openGraph: {
    title: "Sayka",
    description: "成果物を投稿するアプリケーション",
    url: "https://sayka.site/",
    siteName: "Sayka",
    images: [
      {
        url: "https://sayka.site/top-ogp.png",
      },
    ],
    type: "website",
  },
  twitter: {
    card: "summary_large_image",
  },
  icons: "https://sayka.site/logo_solo.png",
};

export default async function RootLayout({
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
