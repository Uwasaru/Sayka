import { FC } from "react";

import { TSayka } from "@/api/mock/type";

import { Sayka } from "./Sayka";

type TProps = {
  sayka: TSayka[];
  changeFilterTag?: (tag: string) => void;
};

export const SaykaList: FC<TProps> = ({ sayka, changeFilterTag }) => {
  return (
    <div className="mt-5 space-y-5">
      {sayka.map((sayka) => (
        <Sayka key={sayka.id} data={sayka} changeFilterTag={changeFilterTag} />
      ))}
    </div>
  );
};
