import { Murecho } from "next/font/google";
import "./globals.css";

import { Providers } from "./Providers";

import type { Metadata } from "next";

const murecho = Murecho({ weight: "100", subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Sayka",
  description: "成果物を投稿するアプリケーション",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <link
        rel="stylesheet"
        href="https://fonts.googleapis.com/css2?family=Press+Start+2P&display=swap"
      />
      <body style={{ fontFamily: `${murecho}` }}>
        <Providers>{children}</Providers>
      </body>
    </html>
  );
}
