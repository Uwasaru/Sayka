"use client";

import Image from "next/image";
import Link from "next/link";
import { FC, useMemo, useState } from "react";
import { AiFillGithub } from "react-icons/ai";
import { BsThreeDotsVertical } from "react-icons/bs";
import { MdOutlineArticle } from "react-icons/md";
import { PiFigmaLogoDuotone } from "react-icons/pi";
import { TfiLayoutSlider } from "react-icons/tfi";
import { TfiWorld } from "react-icons/tfi";

import { TUser } from "@/types/User";
import { TagInItem } from "@/ui/Tag";
import { TooltipUI } from "@/ui/Tooltip";

import { CommentButton } from "./CommentButton";
import { FixModal } from "./FixModal";
import { LikeButton } from "./LikeButton";
import { ShareButton } from "./ShareButton";
import { TSayka } from "@/types/Sayka";

type TProps = {
  sayka: TSayka;
  loginUser?: TUser;
  changeFilterTag?: (tag: string) => void;
};

export const Sayka: FC<TProps> = ({ sayka, changeFilterTag, loginUser }) => {
  return (
    <div className="space-y-5 rounded-md border-4 border-gray-200 p-5 shadow-lg transition-transform hover:-translate-y-1 hover:shadow-xl">
      <SaykaHeader sayka={sayka} loginUser={loginUser} />
      <SaykaBody sayka={sayka} changeFilterTag={changeFilterTag} />
      <SaykaFooter sayka={sayka} loginUser={loginUser} />
    </div>
  );
};

const SaykaHeader: FC<TProps> = ({ sayka, loginUser }) => {
  const [isModalVisible, setModalVisible] = useState(false);

  const isMySayka = useMemo(() => {
    if (!loginUser) return false;
    return loginUser.id === sayka.user_id;
  }, [loginUser, sayka]);

  return (
    <div className="relative flex items-center justify-between">
      <FixModal
        isVisible={isModalVisible}
        onClose={() => setModalVisible(false)}
        id={sayka.id}
      />
      <div className="flex items-center">
        <Image
          src={sayka.user.img}
          alt="user icon"
          width={30}
          height={30}
          className="mr-2 rounded-full"
        />
        <Link
          href={`/mypage/${sayka.user.id}`}
          className="border-b-teal-400 hover:border-b-2">
          @{sayka.user.id}
        </Link>
      </div>
      {isMySayka && (
        <BsThreeDotsVertical
          size={24}
          onClick={() => setModalVisible(!isModalVisible)}
        />
      )}
    </div>
  );
};
const SaykaBody: FC<TProps> = ({ sayka, changeFilterTag }) => {
  return (
    <div className="flex flex-col space-y-3">
      <div className="text-xl font-extrabold md:text-3xl">{sayka.title}</div>
      <div className="text-xs md:text-base">{sayka.description}</div>
      <div className="flex flex-wrap">
        {sayka.tags?.map((tag, i) => (
          <div className="mr-2 py-1" key={i}>
            <TagInItem tag={tag} changeFilterTag={changeFilterTag} />
          </div>
        ))}
      </div>
    </div>
  );
};

const SaykaFooter: FC<TProps> = ({ sayka, loginUser }) => {
  return (
    <div className="flex flex-col-reverse md:flex-row">
      <div className="flex items-center justify-start gap-5 md:w-[50%] ">
        <CommentButton saykaId={sayka.id} commentCount={sayka.comments} />
        <LikeButton
          saykaId={sayka.id}
          likeCount={sayka.favorites}
          loginUser={loginUser}
        />
        {/* <ShareButton saykaId={sayka.id} saykaTitle={sayka.title} /> */}
      </div>
      <div className="flex items-center justify-end gap-5 pb-5 md:w-[50%] md:pb-0">
        {sayka.github_url && (
          <TooltipUI label="ソースコードへ">
            <Link href={sayka.github_url} className="flex items-center gap-1">
              <AiFillGithub size={27} />
            </Link>
          </TooltipUI>
        )}
        {sayka.figma_url && (
          <TooltipUI label="Figmaへ">
            <Link href={sayka.figma_url} className="flex items-center gap-1">
              <PiFigmaLogoDuotone size={24} />
            </Link>
          </TooltipUI>
        )}
        {sayka.slide_url && (
          <TooltipUI label="プレゼンテーションへ">
            <Link href={sayka.slide_url} className="flex items-center gap-1">
              <TfiLayoutSlider size={24} />
            </Link>
          </TooltipUI>
        )}
        {sayka.article_url && (
          <TooltipUI label="記事へ">
            <Link href={sayka.article_url} className="flex items-center gap-1">
              <MdOutlineArticle size={24} />
            </Link>
          </TooltipUI>
        )}
        {sayka.app_url && (
          <TooltipUI label="アプリケーションへ">
            <Link href={sayka.app_url} className="flex items-center gap-1">
              <TfiWorld size={24} />
            </Link>
          </TooltipUI>
        )}
      </div>
    </div>
  );
};
