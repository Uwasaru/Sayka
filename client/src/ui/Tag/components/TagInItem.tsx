import { FC } from "react";

type TProps = {
  tag: {
    id: number;
    name: string;
  };
  changeFilterTag?: (tag: string) => void;
};

export const TagInItem: FC<TProps> = ({ tag, changeFilterTag }) => {
  const handleClick = () => {
    if (!changeFilterTag) return;
    changeFilterTag(tag.name);
  };
  return (
    <div
      className="inline-block whitespace-nowrap rounded-2xl border border-gray-300 bg-gradient-to-r from-gray-200 to-gray-300 px-3 py-1 text-xs font-semibold shadow-md transition-all duration-300 ease-in-out hover:from-gray-100 hover:to-gray-200 hover:shadow-lg md:text-base"
      onClick={handleClick}>
      # {tag.name}
    </div>
  );
};
