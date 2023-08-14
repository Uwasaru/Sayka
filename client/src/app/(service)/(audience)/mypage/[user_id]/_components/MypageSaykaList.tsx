"use client";

import { useState } from "react";
import { Sayka } from "../../../_components/Sayka";
import {
  SaykaList,
  SaykaMockData,
  SaykaMockNullData,
} from "../../../_components/SaykaList";

type list_kind = "MYSELF" | "FAVORITE" | "COMMENT";

type list_type = {
  name: list_kind;
  label: string;
};

const LISTS: list_type[] = [
  {
    name: "MYSELF",
    label: "投稿",
  },
  {
    name: "FAVORITE",
    label: "いいねした投稿",
  },
  {
    name: "COMMENT",
    label: "コメントした投稿",
  },
];

export const MypageSaykaList = () => {
  const [nowList, setList] = useState<list_kind>("MYSELF");
  return (
    <div>
      <div className="border-b border-b-stone-900 font-semibold flex justify-center mx-5">
        <div className="mx-10 flex space-x-10 font-semibold text-gray-600 items-center justify-center text-sm">
          {LISTS.map((list) => {
            const isActive = nowList === list.name;

            return (
              <button
                className={
                  isActive ? "border-b-4 border-sc" : "hover:text-gray-900"
                }
                key={list.name}
                onClick={() => setList(list.name)}>
                {list.label}
              </button>
            );
          })}
        </div>
      </div>
      <div className="mx-10 mt-10">
        {nowList === "MYSELF" && <MyselfSaykaList />}
        {nowList === "FAVORITE" && <FavoriteSaykaList />}
        {nowList === "COMMENT" && <CommentSaykaList />}
      </div>
    </div>
  );
};

const MyselfSaykaList = () => {
  return (
    <div className="mt-5 space-y-5">
      {SaykaMockData.map((sayka) => (
        <Sayka key={sayka.id} data={sayka} />
      ))}
    </div>
  );
};

const FavoriteSaykaList = () => {
  return (
    <div className="mt-5 space-y-5">
      {SaykaMockData.map((sayka) => (
        <Sayka key={sayka.id} data={sayka} />
      ))}
    </div>
  );
};

const CommentSaykaList = () => {
  return (
    <div className="mt-5 space-y-5">
      {SaykaMockNullData.map((sayka) => (
        <Sayka key={sayka.id} data={sayka} />
      ))}
    </div>
  );
};
