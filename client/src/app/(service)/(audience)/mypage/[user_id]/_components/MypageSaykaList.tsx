"use client";

import { FC, useState } from "react";

import { mock_saykas, mock_saykaNull } from "@/api";

import { SaykaList } from "../../../_components/SaykaList";

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

type TProps = {
  userId: number;
};
export const MypageSaykaList: FC<TProps> = ({ userId }) => {
  const [nowList, setList] = useState<list_kind>("MYSELF");
  return (
    <div>
      <div className="mx-5 flex justify-center border-b border-b-stone-900 font-semibold">
        <div className="mx-10 flex items-center justify-center space-x-10 text-sm font-semibold text-gray-600">
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
        {nowList === "MYSELF" && <MyselfSaykaList userId={userId} />}
        {nowList === "FAVORITE" && <FavoriteSaykaList userId={userId} />}
        {nowList === "COMMENT" && <CommentSaykaList userId={userId} />}
      </div>
    </div>
  );
};

const MyselfSaykaList: FC<TProps> = ({ userId }) => {
  return <SaykaList sayka={mock_saykas} />;
};

const FavoriteSaykaList: FC<TProps> = ({ userId }) => {
  return <SaykaList sayka={mock_saykas} />;
};

const CommentSaykaList: FC<TProps> = ({ userId }) => {
  return <SaykaList sayka={mock_saykaNull} />;
};