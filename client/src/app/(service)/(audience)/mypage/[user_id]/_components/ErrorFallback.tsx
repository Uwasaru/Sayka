"use client";

import { FallbackProps } from "react-error-boundary";

export function ErrorFallback({ error }: FallbackProps) {
  return (
    <div className="mt-10 flex flex-col items-center justify-center space-y-5">
      <div className="text-lg font-bold">{error.message}</div>
    </div>
  );
}
