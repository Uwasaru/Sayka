"use client";

import { usePathname, useRouter } from "next/navigation";
import { FC } from "react";
import { LiaCommentAlt } from "react-icons/lia";

type TProps = {
  saykaId: string;
  commentCount: number;
};

export const CommentButton: FC<TProps> = ({ saykaId, commentCount }) => {
  const pathname = usePathname();
  const router = useRouter();

  const newQuery = { modal: "true", sayka_id: String(saykaId) };

  const queryString = new URLSearchParams(newQuery).toString();

  const handleClick = () => {
    router.push(`${pathname}?${queryString}`, { scroll: false });
  };

  return (
    <div className="flex items-center gap-x-1">
      <button
        onClick={handleClick}
        className="transition-transform hover:-translate-y-1">
        <LiaCommentAlt size={24} />
      </button>
      {commentCount}
    </div>
  );
};
