"use client";

import { FC, useEffect, useMemo, useState } from "react";

import { readSaykaTimeline } from "@/api";
import { TSayka } from "@/types/Sayka";
import { TUser } from "@/types/User";
import { ColorButton } from "@/ui/Button";

import { MoreReadSayka } from "./MoreReadSayka";
import { SaykaList } from "./SaykaList";

type TProps = {
  saykas: TSayka[];
  token?: string;
};

// TODO: 削除した時に、UIから消えるように
export const SaykaTimelineList: FC<TProps> = ({ saykas, token }) => {
  const [showMore, setShowMore] = useState(true);
  const [saykaList, setSaykaList] = useState<TSayka[]>(saykas);

  const lastSaykaId = useMemo(() => {
    return saykaList ? saykaList[saykaList.length - 1].id : undefined;
  }, [saykaList]);

  const handleClick = async () => {
    const saykasRes = await readSaykaTimeline(lastSaykaId, token);
    if (saykasRes.type === "error") {
      throw new Error("データが取得できませんでした。");
    }
    const data = saykasRes.value.data;
    if (!data) setShowMore(false);
    setSaykaList([...saykaList, ...data]);
  };

  return (
    <div>
      <SaykaList saykas={saykaList} token={token} noFix={true} />
      {showMore && (
        <div className="mt-10 flex justify-center">
          <ColorButton onClick={handleClick}>もっと見る</ColorButton>
        </div>
      )}
    </div>
  );
};
