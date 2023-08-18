"use client";

import { useAtom } from "jotai";
import Image from "next/image";
import Link from "next/link";
import { useRouter } from "next/navigation";
import React, { FC, useEffect, useState } from "react";
import { BsSend } from "react-icons/bs";
import { GrClose } from "react-icons/gr";
import ReactMarkdown from "react-markdown";

import { modalState } from "@/store/atoms/modalAtom";
import { TSayka } from "@/types/Sayka";
import { CodeBlock } from "@/ui/Text/components/CodeBlock";
import { TComment } from "@/types/Comment";
import { createComment, readCommentBySayka } from "@/api";
import { formatDateToJapanese } from "@/utils/formatDateToJapanese";

type TProps = {
  sayka: TSayka;
  token?: string;
};

export const CommentModal: FC<TProps> = ({ sayka, token }) => {
  const [_, setIsOpen] = useAtom(modalState);
  const [message, setMessage] = useState("");

  const [comments, setComments] = useState<TComment[] | undefined>();
  const [error, setError] = useState<string | undefined>();

  const fetchComments = async () => {
    const res = await readCommentBySayka(sayka.id);
    if (res.type === "error") {
      setError("コメントの取得に失敗しました。");
      return;
    }
    setComments(res.value.data);
  };

  useEffect(() => {
    fetchComments();
  }, []);

  const router = useRouter();

  const handleClose = () => {
    router.back();
  };

  const commentByGuest = () => {
    setIsOpen(true);
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setMessage(e.target.value);
  };

  const handleSend = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (!token) return;
    createComment(
      {
        sayka_id: sayka.id,
        content: message,
      },
      token
    ).then(() => {
      setMessage("");
      fetchComments();
    });
  };

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
                src={sayka.user.img}
                alt="user icon"
                width={30}
                height={30}
                className="rounded-full"
              />
              <Link
                href={`/mypage/${sayka.user.id}`}
                className="border-b-teal-400 hover:border-b-2">
                @{sayka.user.id}
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
          <>
            {error ? (
              <div className="flex h-full items-center justify-center">
                <div className="text-red-500">{error}</div>
              </div>
            ) : !comments ? (
              <div className="flex h-full items-center justify-center">
                <div className="text-gray-500">コメントはありません。</div>
              </div>
            ) : (
              comments.map((c) => (
                <div
                  key={c.id}
                  className="flex flex-col gap-2 border-t border-gray-300 px-8 py-4">
                  <div className="flex items-center space-x-3">
                    <Image
                      src={c.user.img}
                      alt="user icon"
                      width={30}
                      height={30}
                      className="rounded-full"
                    />
                    <Link
                      href={`/mypage/${c.user_id}`}
                      className="border-b-teal-400 hover:border-b-2">
                      @{c.user_id}
                    </Link>
                  </div>
                  <div className="flex flex-col space-y-3">
                    <div className="text-gray-700">
                      <ReactMarkdown className="m-auto" components={CodeBlock}>
                        {c.content}
                      </ReactMarkdown>
                    </div>
                    <div className="text-sm font-medium text-gray-500">
                      {formatDateToJapanese(c.created_at)}
                    </div>
                  </div>
                </div>
              ))
            )}
          </>
        </div>

        {token ? (
          <form
            onSubmit={handleSend}
            className="flex items-center gap-5 border-t border-gray-200 bg-white p-5">
            <textarea
              value={message}
              onChange={handleInputChange}
              className=" w-full rounded-xl border px-5 py-2 text-sm transition-shadow focus:outline-none"
              placeholder="Aa"></textarea>
            <button className="rounded-full p-2 hover:bg-teal-400 ">
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
