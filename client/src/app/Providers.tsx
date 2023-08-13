"use client";

import { ChakraProvider } from "@chakra-ui/react";
import { Provider as JotaiProvider } from "jotai";

export function Providers({ children }: { children: React.ReactNode }) {
  return (
    <ChakraProvider>
      <JotaiProvider>{children}</JotaiProvider>
    </ChakraProvider>
  );
}
