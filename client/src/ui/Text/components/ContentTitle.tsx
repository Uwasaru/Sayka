import React, { FC } from "react";

type TProps = {
  text: string;
};

export const ContentTitle: FC<TProps> = ({ text }) => {
  return (
    <div className="flex items-center justify-center py-7">
      <div className="text-xl font-bold">{text}</div>
    </div>
  );
};
