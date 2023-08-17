"use client";

import { FC, useState } from "react";

import { TUser } from "@/types/User";

import { SaykaList } from "./SaykaList";
import { TSayka } from "@/types/Sayka";
import { MoreReadSayka } from "./MoreReadSayka";
import { ColorButton } from "@/ui/Button";

type TProps = {
  loginUser?: TUser;
  saykas: TSayka[];
};

export const SaykaTimelineList: FC<TProps> = ({ loginUser, saykas }) => {
  const lastSaykaId = saykas[saykas.length - 1].id;
  const [showMore, setShowMore] = useState(false);

  const handleClick = () => {
    setShowMore(true);
  };

  return (
    <div>
      <SaykaList saykas={saykas} loginUser={loginUser} />
      <div className="flex justify-center mt-10">
        <ColorButton onClick={handleClick}>もっと見る</ColorButton>
      </div>
      {showMore && (
        // {/* @ts-expect-error Server Component */}
        <MoreReadSayka lastSaykaId={lastSaykaId} loginUser={loginUser} />
      )}
    </div>
  );
};
