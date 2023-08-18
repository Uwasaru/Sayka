"use client";

import { FC, useState } from "react";

import { TSayka } from "@/types/Sayka";
import { TUser } from "@/types/User";
import { ColorButton } from "@/ui/Button";

import { MoreReadSayka } from "./MoreReadSayka";
import { SaykaList } from "./SaykaList";


type TProps = {
  saykas: TSayka[];
  token?: string;
};

export const SaykaTimelineList: FC<TProps> = ({ saykas, token }) => {
  const [showMore, setShowMore] = useState(false);

  const lastSaykaId = saykas ? saykas[saykas.length - 1].id : undefined;

  const handleClick = () => {
    setShowMore(true);
  };

  return (
    <div>
      <SaykaList saykas={saykas} token={token} />
      {!showMore && lastSaykaId && (
        <div className="mt-10 flex justify-center">
          <ColorButton onClick={handleClick}>もっと見る</ColorButton>
        </div>
      )}
      {showMore && (
        // @ts-expect-error Server Component
        <MoreReadSayka lastSaykaId={lastSaykaId} />
      )}
    </div>
  );
};
