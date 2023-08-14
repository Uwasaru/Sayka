// components/ShareButton.tsx
import { getEnv } from "@/utils/env";
import { FC } from "react";
import { IoMdShareAlt } from "react-icons/io";
import { TwitterShareButton, TwitterIcon } from "react-share";

type ShareButtonProps = {
  saykaId: number;
  saykaTitle: string;
};

const { clientURL } = getEnv();

export const ShareButton: FC<ShareButtonProps> = ({ saykaId, saykaTitle }) => {
  const url = `https://sayka.vercel.app/shareSaykaInformation/${saykaId}`;
  const title = `${saykaTitle}を作成しました！\n#Sayka`;
  return (
    <TwitterShareButton url={url} title={title}>
      <TwitterIcon size={32} round />
    </TwitterShareButton>
  );
};
