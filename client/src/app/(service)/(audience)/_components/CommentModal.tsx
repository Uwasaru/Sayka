"use client";

import Image from "next/image";
import Link from "next/link";
import { useRouter } from "next/navigation";
import React, { FC } from "react";
import { GrClose } from "react-icons/gr";
import { VscSend } from "react-icons/vsc";
import ReactMarkdown from "react-markdown";

import { CodeBlock } from "@/ui/Text/components/CodeBlock";

import { SaykaMockData } from "./SaykaList";
import { useAtom } from "jotai";
import { modalState } from "@/store/atoms/modalAtom";

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
      " > console.log(data) \n\nこれは一人で開発されましたか？よければ一緒にやりたいです！ \n\n```console.log('Hello World')```ga`aa`",
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
  // const user = await (async () => {
  //   const user = await getLoggedInUser(getToken() || '')
  //   if (user.type === 'error') return undefined
  //   return user.value.user
  // })()
  const user = null;
  const [_, setIsOpen] = useAtom(modalState);

  const router = useRouter();

  // SaykaMockDataのidが{id}と一致するものを取得
  const sayka = SaykaMockData.find((sayka) => sayka.id === id);

  if (!sayka) return null;

  const handleClose = () => {
    router.back();
  };

  // commentを取得
  const comments = saykaComment;

  const commentByGuest = () => {
    setIsOpen(true);
  };

  // commentを取得
  return (
    <div className="fixed inset-0 z-10 flex justify-end">
      <div
        className="absolute inset-0 bg-black opacity-50 transition-opacity duration-300 ease-in-out"
        onClick={handleClose}
      />

      <div className="z-10 flex h-full md:w-2/5 w-full flex-col overflow-hidden rounded-l-xl bg-white shadow-2xl">
        {/* sayka */}
        <div className="border-b border-gray-200 p-8">
          <div className="flex justify-end">
            <GrClose
              size={33}
              onClick={handleClose}
              className="rounded-full p-2 transition-colors duration-300 hover:bg-teal-400"
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

        <div className="grow flex-col overflow-y-scroll bg-gray-50">
          {comments.map((comment) => (
            <div
              key={comment.id}
              className="flex flex-col gap-2 border-t border-gray-300 px-8 py-4">
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
                <div className="text-gray-700">
                  <ReactMarkdown className="m-auto" components={CodeBlock}>
                    {comment.contents}
                  </ReactMarkdown>
                </div>
                <div className="text-sm font-medium text-gray-500">
                  {comment.created_at}
                </div>
              </div>
            </div>
          ))}
        </div>

        {user ? (
          <div className="flex items-center gap-5 border-t border-gray-200 bg-white p-5">
            <textarea
              className="focus:shadow-outline w-full rounded-md border px-5 py-2 transition-shadow focus:outline-none"
              placeholder="Aa"></textarea>
            <button className="rounded-full bg-teal-400 p-2 transition-colors duration-300 hover:bg-teal-500">
              <VscSend size={30} fill="white" />
            </button>
          </div>
        ) : (
          <div className="flex justify-end items-center gap-5 border-t border-gray-200 bg-white p-5">
            <button
              onClick={commentByGuest}
              className="rounded bg-sc px-4 py-2 font-semibold text-white transition duration-300 hover:bg-hover-sc">
              コメントする
            </button>
          </div>
        )}
      </div>
    </div>
  );
};
