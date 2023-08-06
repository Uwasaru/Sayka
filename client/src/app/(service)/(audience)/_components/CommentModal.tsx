"use client";

import Image from "next/image";
import Link from "next/link";
import { useRouter } from "next/navigation";
import React, { FC } from "react";

import { SaykaMockData } from "./SaykaList";

import { GrClose } from "react-icons/gr";
import { VscSend } from "react-icons/vsc";

type TProps = {
  id: number;
};

const saykaComment = [
  {
    id: 1,
    contents:
      "面白いですね！！！これからも頑張ってください！面白いですね！！！これからも頑張ってください！面白いですね！！！これからも頑張ってください！面白いですね！！！これからも頑張ってください！面白いですね！！！これからも頑張ってください！面白いですね！！！これからも頑張ってください！面白いですね！！！これからも頑張ってください！面白いですね！！！これからも頑張ってください！",
    created_at: "2021-10-10",
    user: {
      id: 5,
      name: "teyang",
      icon: "/icon2.jpg",
    },
  },
  {
    id: 2,
    contents:
      "これは一人で開発されましたか？よければ一緒にやりたいです！これは一人で開発されましたか？よければ一緒にやりたいです！これは一人で開発されましたか？よければ一緒にやりたいです！これは一人で開発されましたか？よければ一緒にやりたいです！これは一人で開発されましたか？よければ一緒にやりたいです！これは一人で開発されましたか？よければ一緒にやりたいです！これは一人で開発されましたか？よければ一緒にやりたいです！これは一人で開発されましたか？よければ一緒にやりたいです！",
    created_at: "2021-10-10",
    user: {
      id: 5,
      name: "teyang",
      icon: "/icon2.jpg",
    },
  },
  {
    id: 3,
    contents:
      "これは一人で開発されましたか？よければ一緒にやりたいです！これは一人で開発されましたか？よければ一緒にやりたいです！これは一人で開発されましたか？よければ一緒にやりたいです！これは一人で開発されましたか？よければ一緒にやりたいです！これは一人で開発されましたか？よければ一緒にやりたいです！これは一人で開発されましたか？よければ一緒にやりたいです！これは一人で開発されましたか？よければ一緒にやりたいです！これは一人で開発されましたか？よければ一緒にやりたいです！",
    created_at: "2021-10-10",
    user: {
      id: 5,
      name: "teyang",
      icon: "/icon2.jpg",
    },
  },
];

export const CommentModal: FC<TProps> = ({ id }) => {
  const router = useRouter();

  // SaykaMockDataのidが{id}と一致するものを取得
  const sayka = SaykaMockData.find((sayka) => sayka.id === id);

  if (!sayka) return null;

  const handleClose = () => {
    router.back();
  };

  // commentを取得
  const comments = saykaComment;

  // commentを取得
  return (
    <div className="fixed inset-0 z-50 flex justify-end">
      <div
        className="absolute inset-0 bg-black opacity-50"
        onClick={handleClose}
      />

      <div className="z-10 flex h-full w-2/5 flex-col bg-white shadow-xl">
        {/* sayka */}
        <div className="p-8">
          <div className="flex justify-end">
            <GrClose
              size={33}
              onClick={handleClose}
              className="hover:bg-teal-400 rounded-full p-2"
            />
          </div>
          <div className="space-y-5">
            <div className="flex items-center space-x-3">
              <Image
                src={sayka.user.icon}
                alt="user icon"
                width={30}
                height={30}
                className="rounded-full"
              />
              <Link
                href={`/mypage/${sayka.user.id}`}
                className="border-b-teal-400 hover:border-b-2">
                @{sayka.user.name}
              </Link>
            </div>
            <div className="flex flex-col space-y-3">
              <div className="text-xl font-extrabold md:text-3xl">
                {sayka.title}
              </div>
              <div>{sayka.description}</div>
            </div>
          </div>
        </div>
        <div className="border-2 border-teal-400" />
        {/* comment */}
        <div className="flex flex-col flex-grow overflow-y-scroll">
          {comments.map((comment) => (
            <div
              key={comment.id}
              className="flex flex-col px-8 py-4 border-t-gray-300 gap-2 border-t-2">
              <div className="flex items-center space-x-3">
                <Image
                  src={comment.user.icon}
                  alt="user icon"
                  width={30}
                  height={30}
                  className="rounded-full"
                />
                <Link
                  href={`/mypage/${comment.user.id}`}
                  className="border-b-teal-400 hover:border-b-2">
                  @{comment.user.name}
                </Link>
              </div>
              <div className="flex flex-col space-y-3">
                <div className="">{comment.contents}</div>
                <div className="text-sm font-semibold">
                  {comment.created_at}
                </div>
              </div>
            </div>
          ))}
        </div>
        <div className="flex px-5 items-center gap-5">
          <textarea
            className="w-full px-5 py-2 border mb-3 rounded-full"
            placeholder="Write your comment here..."></textarea>
          <button className="rounded-lg">
            <VscSend size={30} fill="#2DD4BF" />
          </button>
        </div>
      </div>
    </div>
  );
};
