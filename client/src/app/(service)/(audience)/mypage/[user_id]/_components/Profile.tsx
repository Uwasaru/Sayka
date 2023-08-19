import Image from "next/image";
import React, { FC } from "react";
import { AiFillGithub } from "react-icons/ai";

import { TUser } from "@/types/User";
import { ColorButton } from "@/ui/Button";
import { TooltipUI } from "@/ui/Tooltip";
import { TProfile } from "@/types/Profile";

type TProps = {
  profile: TProfile;
};
export const Profile: FC<TProps> = ({ profile }) => {
  return (
    <div className="rounded-lg border border-gray-800 p-4">
      <div className="flex flex-col items-center justify-center space-y-5 p-5">
        <Image
          src={profile.user.img}
          alt="user icon"
          width={100}
          height={100}
          className="rounded-full"
        />
        <div className="text-lg font-semibold">@{profile.user.id}</div>
        <div>{profile.user.name}</div>
        <div className="flex items-center justify-center">
          <TooltipUI label="GitHubへ">
            <a href={`https://github.com/${profile.user.id}`}>
              <AiFillGithub size={27} />
            </a>
          </TooltipUI>
        </div>
        <div>
          <div className="mb-5 border-b px-5">
            {profile.favorited_count} contributions
          </div>
          <div className="flex items-center justify-center">
            <div className="flex flex-col items-center text-sm">
              <div>{profile.sayka_count}</div>
              <div>投稿数</div>
            </div>
          </div>
        </div>
        {/* {isMe && <ColorButton>編集する</ColorButton>} */}
      </div>
    </div>
  );
};
