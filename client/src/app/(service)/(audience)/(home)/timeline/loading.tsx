"use client";

import { Spinner } from "@chakra-ui/react";

export default function Loading() {
  return (
    <div className="mt-10 flex flex-col items-center justify-center space-y-5">
      <div className="text-lg font-bold">Loading...</div>
      <Spinner />
    </div>
  );
}
