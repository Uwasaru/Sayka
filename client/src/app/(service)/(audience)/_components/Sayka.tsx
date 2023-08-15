"use client";

import Image from "next/image";
import Link from "next/link";
import { FC, useState } from "react";
import { AiFillGithub } from "react-icons/ai";
import { BsThreeDotsVertical } from "react-icons/bs";
import { TfiLayoutSlider } from "react-icons/tfi";
import { TfiWorld } from "react-icons/tfi";
import { MdOutlineArticle } from "react-icons/md";
import { PiFigmaLogoDuotone } from "react-icons/pi";

import { TSayka } from "@/api/mock/type";
import { TagInItem } from "@/ui/Tag";
import { TooltipUI } from "@/ui/Tooltip";

import { CommentButton } from "./CommentButton";
import { FixModal } from "./FixModal";
import { LikeButton } from "./LikeButton";
import { ShareButton } from "./ShareButton";

type TProps = {
  data: TSayka;
};

type TSaykaProps = {
  data: TSayka;
  changeFilterTag?: (tag: string) => void;
};

export const Sayka: FC<TSaykaProps> = ({ data, changeFilterTag }) => {
  return (
    <div className="space-y-5 rounded-md border-4 border-gray-200 p-5 shadow-lg transition-transform hover:-translate-y-1 hover:shadow-xl">
      <SaykaHeader data={data} />
      <SaykaBody data={data} changeFilterTag={changeFilterTag} />
      <SaykaFooter data={data} />
    </div>
  );
};

const SaykaHeader: FC<TProps> = ({ data }) => {
  const [isModalVisible, setModalVisible] = useState(false);

  return (
    <div className="relative flex items-center justify-between">
      <FixModal
        isVisible={isModalVisible}
        onClose={() => setModalVisible(false)}
        id={data.id}
      />
      <div className="flex items-center">
        <Image
          src={data.user.icon}
          alt="user icon"
          width={30}
          height={30}
          className="mr-2 rounded-full"
        />
        <Link
          href={`/mypage/${data.user.id}`}
          className="border-b-teal-400 hover:border-b-2">
          @{data.user.name}
        </Link>
      </div>
      <div>
        <BsThreeDotsVertical
          size={24}
          onClick={() => setModalVisible(!isModalVisible)}
        />
      </div>
    </div>
  );
};

const SaykaBody: FC<TSaykaProps> = ({ data, changeFilterTag }) => {
  return (
    <div className="flex flex-col space-y-3">
      <div className="text-xl font-extrabold md:text-3xl">{data.title}</div>
      <div className="md:text-base text-xs">{data.description}</div>
      <div className="flex flex-wrap">
        {data.tags?.map((tag) => (
          <div className="py-1 mr-2" key={tag.id}>
            <TagInItem tag={tag} changeFilterTag={changeFilterTag} />
          </div>
        ))}
      </div>
    </div>
  );
};

const SaykaFooter: FC<TProps> = ({ data }) => {
  return (
    <div className="flex flex-col-reverse md:flex-row">
      <div className="flex items-center justify-start gap-5 md:w-[50%] ">
        <CommentButton saykaId={data.id} commentCount={data.comment_count} />
        <LikeButton saykaId={data.id} likeCount={data.like_count} />
        {/* <ShareButton saykaId={data.id} saykaTitle={data.title} /> */}
      </div>
      <div className="flex items-center justify-end gap-5 pb-5 md:w-[50%] md:pb-0">
        {data.github_url && (
          <TooltipUI label="ソースコードへ">
            <Link href={data.github_url} className="flex items-center gap-1">
              <AiFillGithub size={27} />
            </Link>
          </TooltipUI>
        )}
        {data.figma_url && (
          <TooltipUI label="Figmaへ">
            <Link href={data.figma_url} className="flex items-center gap-1">
              <PiFigmaLogoDuotone size={24} />
            </Link>
          </TooltipUI>
        )}
        {data.slide_url && (
          <TooltipUI label="プレゼンテーションへ">
            <Link href={data.slide_url} className="flex items-center gap-1">
              <TfiLayoutSlider size={24} />
            </Link>
          </TooltipUI>
        )}
        {data.article_url && (
          <TooltipUI label="記事へ">
            <Link href={data.article_url} className="flex items-center gap-1">
              <MdOutlineArticle size={24} />
            </Link>
          </TooltipUI>
        )}
        {data.application_url && (
          <TooltipUI label="アプリケーションへ">
            <Link
              href={data.application_url}
              className="flex items-center gap-1">
              <TfiWorld size={24} />
            </Link>
          </TooltipUI>
        )}
      </div>
    </div>
  );
};
