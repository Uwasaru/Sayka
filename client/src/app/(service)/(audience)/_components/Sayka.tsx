"use client";

import Image from "next/image";
import Link from "next/link";
import { FC, useState } from "react";
import { AiFillGithub } from "react-icons/ai";
import { BsThreeDotsVertical } from "react-icons/bs";
import { TfiLayoutSlider } from "react-icons/tfi";
import { TfiWorld } from "react-icons/tfi";

import { TagInItem } from "@/ui/Tag";
import { TooltipUI } from "@/ui/Tooltip";

import { CommentModalButton } from "./CommentModalButton";
import { FixModal } from "./FixModal";
import { LikeButton } from "./LikeButton";
import { ShareButton } from "./ShareButton";

export type TSayka = {
  id: number;
  title: string;
  description: string;
  like_count: number;
  comment_count: number;
  github_url?: string;
  slide_url?: string;
  application_url?: string;
  user: {
    id: number;
    name: string;
    icon: string;
  };
  tags?: {
    id: number;
    name: string;
  }[];
};

type TProps = {
  data: TSayka;
};

type TSaykaProps = {
  data: TSayka;
  changeFilterTag?: (tag: string) => void;
};

export const Sayka: FC<TSaykaProps> = ({ data, changeFilterTag }) => {
  return (
    <div className="space-y-5 rounded-md border border-gray-900 p-5 shadow-lg transition-transform hover:-translate-y-1 hover:shadow-xl">
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
          className="rounded-full mr-1"
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
      <div>{data.description}</div>
      <div className="flex space-x-2">
        {data.tags?.map((tag) => (
          <TagInItem key={tag.id} tag={tag} changeFilterTag={changeFilterTag} />
        ))}
      </div>
    </div>
  );
};

const SaykaFooter: FC<TProps> = ({ data }) => {
  return (
    <div className="flex md:flex-row flex-col-reverse">
      <div className="flex md:w-[50%] items-center justify-start gap-5 ">
        <div className="flex items-center gap-1">
          <CommentModalButton id={data.id} />
          {data.comment_count}
        </div>
        <div className="flex items-center gap-x-1">
          <LikeButton saykaId={data.id} />
          {data.like_count}
        </div>
        <ShareButton saykaId={data.id} />
      </div>
      <div className="flex md:w-[50%] items-center justify-end gap-5 pb-5 md:pb-0">
        {data.github_url && (
          <TooltipUI label="ソースコードへ">
            <Link href={data.github_url} className="flex items-center gap-1">
              <AiFillGithub size={27} />
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
