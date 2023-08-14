// components/ShareButton.tsx
import { getEnv } from "@/utils/env";
import { FC } from "react";
import { IoMdShareAlt } from "react-icons/io";
import { TwitterShareButton, TwitterIcon } from "react-share";

type ShareButtonProps = {
  saykaId: number;
};

const { clientURL } = getEnv();

// export const ShareButton: FC<ShareButtonProps> = ({ saykaId }) => {
//   return (
//     <button onClick={() => handleShare(saykaId)}>
//       <IoMdShareAlt size={24} />
//     </button>
//   );
// };

// const handleShare = async (saykaId: number) => {
//   const postURL = `${clientURL}/share/${saykaId}`; // 例: 'https://yourdomain.com/post/1'

//   try {
//     await navigator.clipboard.writeText(postURL);
//     alert("URL copied to clipboard!");
//   } catch (err) {
//     alert("Failed to copy URL: ");
//   }
// };

export const ShareButton: FC<ShareButtonProps> = ({ saykaId }) => {
  const url = `https://sayka.vercel.app/share/${saykaId}`;
  const title = "成果物共有アプリ『Sayka』を作成しました！\n#Sayka\n";
  return (
    <TwitterShareButton url={url} title={title}>
      <TwitterIcon size={32} round />
    </TwitterShareButton>
  );
};
