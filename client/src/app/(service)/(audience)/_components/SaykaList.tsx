import { FC } from "react";

import { TSayka } from "@/api/mock/type";

import { Sayka } from "./Sayka";
import { TUser } from "@/types/User";

type TProps = {
  sayka: TSayka[];
  changeFilterTag?: (tag: string) => void;
  user?: TUser;
};

export const SaykaList: FC<TProps> = ({ sayka, changeFilterTag, user }) => {
  return (
    <div className="mt-5 space-y-5">
      {sayka.map((sayka) => (
        <Sayka
          key={sayka.id}
          data={sayka}
          changeFilterTag={changeFilterTag}
          user={user}
        />
      ))}
    </div>
  );
};
