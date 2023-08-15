"use client";

import { FC, useState } from "react";

import { mock_saykas, mock_saykaNull } from "@/api";
import { TUser } from "@/types/User";

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
  userId: TUser["id"];
};
export const MypageSaykaList: FC<TProps> = ({ userId }) => {
  const [nowList, setList] = useState<list_kind>("MYSELF");
  return (
    <div>
      <div className="mx-5 flex justify-center border-b border-b-stone-900 font-semibold">
        <div className="flex items-center justify-center space-x-5 text-sm font-semibold text-gray-600 ">
          {LISTS.map((list) => {
            const isActive = nowList === list.name;

            return (
              <button
                className={
                  isActive
                    ? "w-24 border-b-4 border-sc text-xs md:w-36 lg:w-32 lg:text-sm "
                    : "w-24 text-xs hover:text-gray-900 md:w-36 lg:w-32 lg:text-sm "
                }
                key={list.name}
                onClick={() => setList(list.name)}>
                {list.label}
              </button>
            );
          })}
        </div>
      </div>
      <div className="mt-10">
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
