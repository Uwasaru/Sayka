import Link from "next/link";
import { FC } from "react";

import { TUser } from "@/types/User";

type list_kind = "myself" | "favorite" | "comment";

type list_type = {
  name: list_kind;
  label: string;
  path: string;
};

const LISTS: list_type[] = [
  {
    name: "myself",
    label: "投稿",
    path: "",
  },
  {
    name: "favorite",
    label: "いいねした投稿",
    path: "/favorite",
  },
  {
    name: "comment",
    label: "コメントした投稿",
    path: "/comment",
  },
];

type TProps = {
  userId: TUser["id"];
  feature: list_kind;
};
export const MypageNavigation: FC<TProps> = ({ userId, feature }) => {
  return (
    <div className="mx-0 md:mx-5 flex justify-center border-b border-b-stone-900 font-semibold">
      <div className="flex items-center justify-center space-x-2 md:space-x-5 text-center text-sm font-semibold text-gray-600">
        {LISTS.map((list) => {
          const isActive = feature === list.name;

          return (
            <Link
              href={`/mypage/${userId}${list.path}`}
              className={
                isActive
                  ? "w-24 border-b-4 border-sc text-xs md:w-36 lg:w-28 lg:text-sm "
                  : "w-24 text-xs hover:text-gray-900 md:w-36 lg:w-28 lg:text-sm "
              }
              key={list.name}>
              {list.label}
            </Link>
          );
        })}
      </div>
    </div>
  );
};
