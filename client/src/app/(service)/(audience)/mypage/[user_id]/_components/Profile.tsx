import Image from "next/image";
import React, { FC } from "react";
import { AiFillGithub } from "react-icons/ai";

import { TooltipUI } from "@/ui/Tooltip";
import { TUser } from "@/types/User";

type TProps = {
  user: TUser;
};
export const Profile: FC<TProps> = ({ user }) => {
  return (
    <div className="rounded-lg border border-gray-800 p-4">
      <div className="flex flex-col items-center justify-center space-y-5 p-5">
        <Image
          src={user.Img}
          alt="user icon"
          width={100}
          height={100}
          className="rounded-full"
        />
        <div className="text-lg font-semibold">@{user.Name}</div>
        <div>B3 / Student / Software Engineer</div>
        <div className="flex items-center justify-center">
          <TooltipUI label="GitHubへ">
            <a href={`https://github.com/${user.Name}`}>
              <AiFillGithub size={27} />
            </a>
          </TooltipUI>
        </div>
        <div>
          <div className="mb-5 border-b px-5">120 contributions</div>
          <div className="flex items-center justify-center">
            <div className="flex flex-col items-center text-sm">
              <div>3</div>
              <div>投稿数</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};
