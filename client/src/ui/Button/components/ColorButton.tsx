import React, { FC } from "react";

type TProps = {
  label: string;
  onClick?: () => void;
};

export const ColorButton: FC<TProps> = ({ label, onClick }) => {
  return (
    <button
      onClick={onClick}
      className="rounded bg-sc px-4 py-2 font-semibold text-white transition duration-300 hover:bg-hover-sc md:text-base text-sm">
      {label}
    </button>
  );
};
