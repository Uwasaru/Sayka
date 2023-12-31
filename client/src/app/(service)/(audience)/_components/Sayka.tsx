"use client";

import { useAtom } from "jotai";
import Image from "next/image";
import Link from "next/link";
import { FC, useMemo, useState } from "react";
import { AiFillGithub } from "react-icons/ai";
import { BsThreeDotsVertical } from "react-icons/bs";
import { MdOutlineArticle } from "react-icons/md";
import { PiFigmaLogoDuotone } from "react-icons/pi";
import { TfiLayoutSlider } from "react-icons/tfi";
import { TfiWorld } from "react-icons/tfi";

import { modalState } from "@/store/atoms/modalAtom";
import { TSayka } from "@/types/Sayka";
import { TUser } from "@/types/User";
import { TagInItem } from "@/ui/Tag";
import { TooltipUI } from "@/ui/Tooltip";

import { CommentButton } from "./CommentButton";
import { FixModal } from "./FixModal";
import { LikeButton } from "./LikeButton";
import { ShareButton } from "./ShareButton";

type TProps = {
  sayka: TSayka;
  changeFilterTag?: (tag: string) => void;
  token?: string;
  noFix?: boolean; // timelineでfixした時に挙動がおかしいので、その対策
};

export const Sayka: FC<TProps> = ({ sayka, changeFilterTag, token, noFix }) => {
  return (
    <div className="space-y-5 rounded-md border-4 border-gray-200 p-5 shadow-lg transition-transform hover:-translate-y-1 hover:shadow-xl">
      <SaykaHeader sayka={sayka} token={token} noFix={noFix} />
      <SaykaBody sayka={sayka} changeFilterTag={changeFilterTag} />
      <SaykaFooter sayka={sayka} token={token} />
    </div>
  );
};

const SaykaHeader: FC<TProps> = ({ sayka, token, noFix }) => {
  const [isModalVisible, setModalVisible] = useState(false);

  return (
    <div className="relative flex items-center justify-between">
      <FixModal
        isVisible={isModalVisible}
        onClose={() => setModalVisible(false)}
        id={sayka.id}
        token={token}
      />
      <div className="flex items-center">
        <TooltipUI label="マイページへ">
          <Link href={`/mypage/${sayka.user.id}`}>
            <Image
              src={sayka.user.img}
              alt="user icon"
              width={30}
              height={30}
              className="mr-2 rounded-full"
            />
          </Link>
        </TooltipUI>
        <Link
          href={`/mypage/${sayka.user.id}`}
          className="border-b-teal-400 hover:border-b-2">
          @{sayka.user.id}
        </Link>
      </div>
      {sayka.is_me && !noFix && (
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

const SaykaFooter: FC<TProps> = ({ sayka, token }) => {
  const [liked, setLiked] = useState(sayka.is_favorite);
  const [likeCountState, setLikeCountState] = useState(sayka.favorite_count);
  const [_, setIsOpen] = useAtom(modalState);

  const handleHeartClickByGuest = () => {
    setIsOpen(true);
  };

  const handleHeartClick = () => {
    setLiked(!liked);
    if (liked) {
      setLikeCountState(likeCountState - 1);
    } else {
      setLikeCountState(likeCountState + 1);
    }
  };

  return (
    <div className="flex flex-col-reverse md:flex-row">
      <div className="flex items-center justify-start gap-5 md:w-[50%] ">
        <CommentButton saykaId={sayka.id} commentCount={sayka.comment_count} />
        <LikeButton
          saykaId={sayka.id}
          liked={liked}
          likeCount={likeCountState}
          handleHeartClick={handleHeartClick}
          handleHeartClickByGuest={handleHeartClickByGuest}
          token={token}
        />
        <ShareButton saykaId={sayka.id} saykaTitle={sayka.title} />
      </div>
      <div className="flex items-center justify-end gap-5 pb-5 md:w-[50%] md:pb-0">
        {sayka.github_url && (
          <TooltipUI label="ソースコードへ">
            <a
              href={sayka.github_url}
              className="flex items-center gap-1"
              target="_blank"
              rel="noopener noreferrer">
              <AiFillGithub size={27} />
            </a>
          </TooltipUI>
        )}
        {sayka.figma_url && (
          <TooltipUI label="Figmaへ">
            <a
              href={sayka.figma_url}
              className="flex items-center gap-1"
              target="_blank"
              rel="noopener noreferrer">
              <PiFigmaLogoDuotone size={24} />
            </a>
          </TooltipUI>
        )}
        {sayka.slide_url && (
          <TooltipUI label="プレゼンテーションへ">
            <a
              href={sayka.slide_url}
              className="flex items-center gap-1"
              target="_blank"
              rel="noopener noreferrer">
              <TfiLayoutSlider size={24} />
            </a>
          </TooltipUI>
        )}
        {sayka.article_url && (
          <TooltipUI label="記事へ">
            <a
              href={sayka.article_url}
              className="flex items-center gap-1"
              target="_blank"
              rel="noopener noreferrer">
              <MdOutlineArticle size={24} />
            </a>
          </TooltipUI>
        )}
        {sayka.app_url && (
          <TooltipUI label="アプリケーションへ">
            <a
              href={sayka.app_url}
              className="flex items-center gap-1"
              target="_blank"
              rel="noopener noreferrer">
              <TfiWorld size={24} />
            </a>
          </TooltipUI>
        )}
      </div>
    </div>
  );
};
