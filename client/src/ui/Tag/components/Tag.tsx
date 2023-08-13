import { FC } from "react";

type TProps = {
  text: string;
};

export const Tag: FC<TProps> = ({ text }) => {
  return (
    <div className="inline-block rounded-2xl border border-gray-300 bg-gradient-to-r from-gray-400 to-gray-500 px-3 py-1 text-sm font-semibold shadow-md">
      # {text}
    </div>
  );
};
