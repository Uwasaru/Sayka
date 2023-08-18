"use client";

import { FC, useState } from "react";

import { TSayka } from "@/types/Sayka";
import { TUser } from "@/types/User";

import { SaykaList } from "./SaykaList";
import { TagFilter } from "./TagFilter";

type TProps = {
  user?: TUser;
  saykas?: TSayka[];
};
export const SaykaListWithFilter: FC<TProps> = ({ user, saykas }) => {
  // const [tagState, setTagState] = useState<string>("ALL");
  const [lastSaykaId, setLastSaykaId] = useState<string | undefined>(undefined);

  // const changeFilterTag = (tag: string) => {
  //   setTagState(tag);
  // };

  // const handleSubmit = (tag: string) => {
  //   console.log("Submitted:", tag);
  // };

  // const clickTag = (tag: string) => {
  //   changeFilterTag(tag);
  //   if (tagState !== tag) {
  //     handleSubmit(tag);
  //   }
  // };

  return (
    <div>
      {/* <TagFilter
        tag={tagState}
        changeFilterTag={changeFilterTag}
        onSubmit={handleSubmit}
      /> */}
      {/* <SaykaList saykas={saykas} user={user} /> */}
      <button>もっと見る</button>
    </div>
  );
};
