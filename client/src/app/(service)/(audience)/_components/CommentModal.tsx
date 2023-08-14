"use client";

import { useAtom } from "jotai";
import Image from "next/image";
import Link from "next/link";
import { useRouter } from "next/navigation";
import React, { FC, useEffect, useState } from "react";
import { GrClose } from "react-icons/gr";
import ReactMarkdown from "react-markdown";

import { mock_saykaComments_saykaId1, mock_saykas } from "@/api";
import { modalState } from "@/store/atoms/modalAtom";
import { CodeBlock } from "@/ui/Text/components/CodeBlock";
import { TComment } from "@/api/mock/type";
import { BsSend } from "react-icons/bs";

type TProps = {
  id: number;
};

export const CommentModal: FC<TProps> = ({ id }) => {
  // const user = await (async () => {
  //   const user = await getLoggedInUser(getToken() || '')
  //   if (user.type === 'error') return undefined
  //   return user.value.user
  // })()
  const user = null;
  const [_, setIsOpen] = useAtom(modalState);
  const [commentList, setCommentList] = useState<TComment[]>([]);
  const [comment, setComment] = useState("");

  // commentを取得
  useEffect(() => {
    const comments = mock_saykaComments_saykaId1;
    setCommentList(comments);
  }, []);

  const router = useRouter();

  // SaykaMockDataのidが{id}と一致するものを取得
  const sayka = mock_saykas.find((sayka) => sayka.id === id);

  if (!sayka) return null;

  const handleClose = () => {
    router.back();
  };

  const commentByGuest = () => {
    setIsOpen(true);
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setComment(e.target.value);
  };

  const handleSend = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    // TODO: リクエスト
    console.log(comment);
    setComment("");
  };

  // commentを取得
  return (
    <div className="fixed inset-0 z-10 flex justify-end">
      <div
        className="absolute inset-0 bg-black opacity-50 transition-opacity duration-300 ease-in-out"
        onClick={handleClose}
      />

      <div className="z-10 flex h-full w-full flex-col overflow-hidden rounded-l-xl bg-white shadow-2xl md:w-2/5">
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
          {commentList.map((comment) => (
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

        {!user ? (
          <form
            onSubmit={handleSend}
            className="flex items-center gap-5 border-t border-gray-200 bg-white p-5">
            <textarea
              value={comment}
              onChange={handleInputChange}
              className="focus:shadow-outline w-full rounded-xl border px-5 py-2 text-sm transition-shadow focus:outline-none"
              placeholder="Aa"></textarea>
            <button className="rounded-full hover:bg-teal-400 p-2 ">
              <BsSend size={24} />
            </button>
          </form>
        ) : (
          <div className="flex items-center justify-end gap-5 border-t border-gray-200 bg-white p-5">
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
