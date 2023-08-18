"use client";

import { FC } from "react";
import { TwitterShareButton, TwitterIcon } from "react-share";

import { TooltipUI } from "@/ui/Tooltip";

type ShareButtonProps = {
  saykaId: string;
  saykaTitle: string;
};

export const ShareButton: FC<ShareButtonProps> = ({ saykaId, saykaTitle }) => {
  const url = `https://sayka.vercel.app/shareSaykaInformation/${saykaId}`;
  const title = `${saykaTitle}を作成しました！\n#Sayka`;
  return (
    <TooltipUI label="Xで共有">
      <TwitterShareButton url={url} title={title}>
        <TwitterIcon size={24} round />
      </TwitterShareButton>
    </TooltipUI>
  );
};

export const ShareLongButton: FC<ShareButtonProps> = ({
  saykaId,
  saykaTitle,
}) => {
  const url = `https://sayka.vercel.app/shareSaykaInformation/${saykaId}`;
  const title = `${saykaTitle}を作成しました！\n#Sayka`;
  return (
    <TwitterShareButton url={url} title={title}>
      <div className="text-lg text-white bg-sky-300 font-bold p-3 rounded-md">
        Twitterでシェアする
      </div>
    </TwitterShareButton>
  );
};
