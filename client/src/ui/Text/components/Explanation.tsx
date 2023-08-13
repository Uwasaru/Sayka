import React, { FC } from "react";

type TProps = {
  text: string;
};

export const Explanation: FC<TProps> = ({ text }) => {
  return (
    <div className="py-2">
      <div className="pb-3 text-sm text-gray-800">{text}</div>
    </div>
  );
};
