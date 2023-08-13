// components/ShareButton.tsx
import { getEnv } from "@/utils/env";
import { FC } from "react";
import { IoMdShareAlt } from "react-icons/io";

type ShareButtonProps = {
  saykaId: number;
};

const { clientURL } = getEnv();

export const ShareButton: FC<ShareButtonProps> = ({ saykaId }) => {
  return (
    <button onClick={() => handleShare(saykaId)}>
      <IoMdShareAlt size={24} />
    </button>
  );
};

const handleShare = async (saykaId: number) => {
  const postURL = `${clientURL}/share/${saykaId}`; // 例: 'https://yourdomain.com/post/1'

  try {
    await navigator.clipboard.writeText(postURL);
    alert("URL copied to clipboard!");
  } catch (err) {
    alert("Failed to copy URL: ");
  }
};

export { handleShare };
