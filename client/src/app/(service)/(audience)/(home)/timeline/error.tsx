"use client";

export default function Error({
  error,
  reset,
}: {
  error: Error;
  reset: () => void;
}) {
  return (
    <div className="mt-10 flex flex-col items-center justify-center space-y-5">
      <div className="text-lg font-bold">{error.message}</div>
    </div>
  );
}
