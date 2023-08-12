import { FC } from "react";

type TProps = {
  text: string;
  onClick?: () => void;
};

export const Tag: FC<TProps> = ({ text, onClick }) => {
  return (
    <div className="inline-block rounded-lg border border-gray-300 bg-gradient-to-r from-gray-400 to-gray-500 px-3 py-1 text-sm font-semibold shadow-md">
      # {text}
    </div>
  );
};
