import { useAtom } from "jotai";
import React, { FC, useState } from "react";
import { VscHeartFilled } from "react-icons/vsc";
import { VscHeart } from "react-icons/vsc";

import { modalState } from "@/store/atoms/modalAtom";
import { TUser } from "@/types/User";

type TProps = {
  saykaId: number;
  likeCount: number;
  user?: TUser;
};

export const LikeButton: FC<TProps> = ({ saykaId, likeCount, user }) => {
  const [liked, setLiked] = useState(false);
  const [likeCountState, setLikeCountState] = useState(likeCount);

  const handleHeartClick = () => {
    setLiked(!liked);
    if (liked) {
      setLikeCountState(likeCountState - 1);
    } else {
      setLikeCountState(likeCountState + 1);
    }
    // リクエスト
    // TODO: すぐ送らずに遅延させる
  };

  const [_, setIsOpen] = useAtom(modalState);
  const handleHeartClickByGuest = () => {
    setIsOpen(true);
  };

  if (!user) {
    return (
      <div className="flex items-center gap-x-1">
        <VscHeart size={24} onClick={handleHeartClickByGuest} />
        {likeCountState}
      </div>
    );
  }

  if (!liked) {
    return (
      <div className="flex items-center gap-x-1">
        <VscHeart size={24} onClick={handleHeartClick} />
        {likeCountState}
      </div>
    );
  }

  return (
    <div className="flex items-center gap-x-1">
      <VscHeartFilled size={24} color="#2DD4BF" onClick={handleHeartClick} />
      {likeCountState}
    </div>
  );
};
