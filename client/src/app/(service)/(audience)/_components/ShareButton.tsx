// components/ShareButton.tsx
import { FC } from "react";
import { TwitterShareButton, TwitterIcon } from "react-share";

import { TooltipUI } from "@/ui/Tooltip";

type ShareButtonProps = {
  saykaId: number;
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
