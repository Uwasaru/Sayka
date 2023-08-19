"use client";

import { FallbackProps } from "react-error-boundary";

const errorMessage = ["データが取得できませんでした", "投稿がありません"];

export function ErrorFallback({ error }: FallbackProps) {
  // error messageはerrorMessageType以外は全て"データが取得できませんでした"にする
  let err = error.message;
  if (!errorMessage.includes(err)) {
    err = errorMessage[0];
  }

  return (
    <div className="mt-10 flex flex-col items-center justify-center space-y-5">
      <div className="text-lg font-bold">{err}</div>
    </div>
  );
}
