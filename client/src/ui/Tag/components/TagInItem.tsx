import { FC } from "react";

type TProps = {
  tag: {
    id: number;
    name: string;
  };
};

export const TagInItem: FC<TProps> = ({ tag }) => {
  return (
    <div className="inline-block rounded-lg border border-gray-300 bg-gradient-to-r from-gray-400 to-gray-500 px-3 py-1 text-sm font-semibold shadow-md transition-all duration-300 ease-in-out hover:from-gray-300 hover:to-gray-400 hover:shadow-lg">
      # {tag.name}
    </div>
  );
};
