"use client";

import { FC } from "react";
import { TwitterShareButton, TwitterIcon } from "react-share";

import { TooltipUI } from "@/ui/Tooltip";
import { getEnv } from "@/utils/env";

type ShareButtonProps = {
  saykaId: string;
  saykaTitle: string;
};

const { clientURL } = getEnv();

export const ShareButton: FC<ShareButtonProps> = ({ saykaId, saykaTitle }) => {
  const url = `${clientURL}//shareSaykaInformation/${saykaId}`;
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
  const url = `${clientURL}/shareSaykaInformation/${saykaId}`;
  const title = `${saykaTitle}を作成しました！\n#Sayka`;
  return (
    <TwitterShareButton url={url} title={title}>
      <div className="rounded-md bg-sky-300 p-3 text-lg font-bold text-white">
        Twitterでシェアする
      </div>
    </TwitterShareButton>
  );
};
