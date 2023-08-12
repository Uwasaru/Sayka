"use client";
import { Tooltip } from "@chakra-ui/react";
import { FC } from "react";

type TProps = {
  label: string;
  children: React.ReactNode;
  placement?: "top" | "right" | "bottom" | "left";
};

export const TooltipUI: FC<TProps> = ({
  label,
  children,
  placement = "top",
}) => {
  return (
    <Tooltip label={label} placement={placement} hasArrow>
      {children}
    </Tooltip>
  );
};
