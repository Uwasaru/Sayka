import Image from "next/image";
import Link from "next/link";
import { FC } from "react";
import { AiFillGithub } from "react-icons/ai";
import { TfiLayoutSlider } from "react-icons/tfi";
import { TfiWorld } from "react-icons/tfi";
import { VscHeartFilled } from "react-icons/vsc";

import { TagInItem } from "@/ui/Tag";
import { TooltipUI } from "@/ui/Tooltip";

import { CommentModalButton } from "./CommentModalButton";

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

type TSaykaHeaderProps = {
  user: TSayka["user"];
};

export const Sayka: FC<TProps> = ({ data }) => {
  return (
    <div className="space-y-5 rounded-md border border-gray-900 p-5">
      <SaykaHeader user={data.user} />
      <SaykaBody data={data} />
      <SaykaFooter data={data} />
    </div>
  );
};

const SaykaHeader: FC<TSaykaHeaderProps> = ({ user }) => {
  return (
    <div className="flex items-center space-x-3">
      <Image
        src={user.icon}
        alt="user icon"
        width={30}
        height={30}
        className="rounded-full"
      />
      <Link
        href={`/mypage/${user.id}`}
        className="border-b-teal-400 hover:border-b-2">
        @{user.name}
      </Link>
    </div>
  );
};

const SaykaBody: FC<TProps> = ({ data }) => {
  return (
    <div className="flex flex-col space-y-3">
      <div className="text-xl font-extrabold md:text-3xl">{data.title}</div>
      <div>{data.description}</div>
      <div className="flex space-x-2">
        {data.tags?.map((tag) => (
          <TagInItem key={tag.id} tag={tag} />
        ))}
      </div>
    </div>
  );
};

const SaykaFooter: FC<TProps> = ({ data }) => {
  return (
    <div className="flex">
      <div className="flex w-[50%] items-center justify-start gap-5">
        <div className="flex items-center gap-1">
          <CommentModalButton id={data.id} />
          {data.comment_count}
        </div>
        <div className="flex items-center gap-x-1">
          <VscHeartFilled size={24} color="#2DD4BF" />
          {data.like_count}
        </div>
      </div>
      <div className="flex w-[50%] items-center justify-end gap-5">
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
