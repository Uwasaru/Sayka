"use client";
import { Tooltip } from "@chakra-ui/react";
import { FC } from "react";

type TProps = {
  label: string;
  children: React.ReactNode;
};

export const TooltipUI: FC<TProps> = ({ label, children }) => {
  return (
    <Tooltip label={label} placement="top" hasArrow>
      {children}
    </Tooltip>
  );
};
