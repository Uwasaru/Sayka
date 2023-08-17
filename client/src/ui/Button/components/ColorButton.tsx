import React, { FC } from "react";

type TProps = {
  onClick?: () => void;
  children?: React.ReactNode;
};

export const ColorButton: FC<TProps> = ({ children, onClick }) => {
  return (
    <button
      onClick={onClick}
      className="rounded bg-sc px-4 py-2 text-sm font-semibold text-white transition duration-300 hover:bg-hover-sc md:text-base">
      {children}
    </button>
  );
};
