import { FC } from "react";

type TProps = {
  tag: {
    id: number;
    name: string;
  };
};
export const TagInItem: FC<TProps> = ({ tag }) => {
  return (
    <div className="rounded-lg bg-gray-400 px-2 py-1 text-sm"># {tag.name}</div>
  );
};
