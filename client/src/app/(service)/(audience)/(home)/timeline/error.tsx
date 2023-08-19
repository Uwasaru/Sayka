"use client";

const errorMessage = ["データが取得できませんでした", "投稿がありません"];

export default function Error({
  error,
  reset,
}: {
  error: Error;
  reset: () => void;
}) {
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
