import React, { FC } from "react";

type TProps = {
  text: string;
};

export const ContentSubTitle: FC<TProps> = ({ text }) => {
  return (
    <div className="py-5">
      <div className="border-b-2 border-gray-500 pb-3 text-xl font-bold text-gray-500">
        {text}
      </div>
    </div>
  );
};
