"use client";

import { FC, useEffect, useMemo, useState } from "react";

import { TSayka } from "@/types/Sayka";
import { TUser } from "@/types/User";
import { ColorButton } from "@/ui/Button";

import { MoreReadSayka } from "./MoreReadSayka";
import { SaykaList } from "./SaykaList";
import { readSaykaTimeline } from "@/api";

type TProps = {
  saykas: TSayka[];
  token?: string;
};

export const SaykaTimelineList: FC<TProps> = ({ saykas, token }) => {
  const [showMore, setShowMore] = useState(true);
  const [saykaList, setSaykaList] = useState<TSayka[]>(saykas);

  const lastSaykaId = useMemo(() => {
    return saykaList ? saykaList[saykaList.length - 1].id : undefined;
  }, [saykaList]);

  const handleClick = async () => {
    const saykasRes = await readSaykaTimeline(lastSaykaId, token);
    console.log(saykasRes);
    if (saykasRes.type === "error") {
      throw new Error("データが取得できませんでした。");
    }
    const data = saykasRes.value.data;
    if (!data) setShowMore(false);
    console.log(data);
    setSaykaList([...saykaList, ...data]);
  };

  return (
    <div>
      <SaykaList saykas={saykaList} token={token} />
      {showMore && (
        <div className="mt-10 flex justify-center">
          <ColorButton onClick={handleClick}>もっと見る</ColorButton>
        </div>
      )}
    </div>
  );
};
