"use client";

import Link from "next/link";
import { usePathname, useRouter } from "next/navigation";
import { FC } from "react";
import { LiaCommentAlt } from "react-icons/lia";

type TProps = {
  id: number;
};

export const CommentModalButton: FC<TProps> = ({ id }) => {
  const pathname = usePathname();
  const router = useRouter();

  const newQuery = { modal: "true", sayka_id: String(id) };

  const queryString = new URLSearchParams(newQuery).toString();

  const handleClick = () => {
    router.push(`${pathname}?${queryString}`, { scroll: false });
  };

  return (
    <div>
      <button onClick={handleClick}>
        <LiaCommentAlt size={24} />
      </button>
    </div>
  );
};
