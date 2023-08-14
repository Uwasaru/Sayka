// components/ShareButton.tsx
import { getEnv } from "@/utils/env";
import { FC } from "react";
import { IoMdShareAlt } from "react-icons/io";
import { TwitterShareButton, TwitterIcon } from "react-share";

type ShareButtonProps = {
  saykaId: number;
};

const { clientURL } = getEnv();

export const ShareButton: FC<ShareButtonProps> = ({ saykaId }) => {
  const url = `https://sayka.vercel.app/shareSaykaInformation/${saykaId}`;
  const title = "成果物共有アプリ『Sayka』を作成しました！\n#Sayka";
  return (
    <TwitterShareButton url={url} title={title}>
      <TwitterIcon size={32} round />
    </TwitterShareButton>
  );
};
