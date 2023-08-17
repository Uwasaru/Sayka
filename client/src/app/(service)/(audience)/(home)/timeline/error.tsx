"use client";

import { ColorButton } from "@/ui/Button";
import { useEffect } from "react";

export default function Error({
  error,
  reset,
}: {
  error: Error;
  reset: () => void;
}) {
  return (
    <div className="flex flex-col justify-center mt-10 items-center space-y-5">
      <div className="text-lg font-bold">{error.message}</div>
    </div>
  );
}
