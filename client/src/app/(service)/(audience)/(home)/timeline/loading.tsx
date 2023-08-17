"use client";

import { Spinner } from "@chakra-ui/react";

export default function Loading() {
  return (
    <div className="flex flex-col justify-center mt-10 items-center space-y-5">
      <div className="text-lg font-bold">Loading...</div>
      <Spinner />
    </div>
  );
}
