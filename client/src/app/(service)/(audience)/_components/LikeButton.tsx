"use client";

import { fixFavorite } from "@/api";
import React, { FC } from "react";
import { VscHeartFilled } from "react-icons/vsc";
import { VscHeart } from "react-icons/vsc";

type TProps = {
  saykaId: string;
  handleHeartClick: () => void;
  handleHeartClickByGuest: () => void;
  liked: boolean;
  likeCount: number;
  token?: string;
};

export const LikeButton = ({
  saykaId,
  handleHeartClick,
  liked,
  likeCount,
  handleHeartClickByGuest,
  token,
}: TProps) => {
  if (!token) {
    return (
      <div className="flex items-center gap-x-1">
        <VscHeart size={24} onClick={handleHeartClickByGuest} />
        {likeCount}
      </div>
    );
  }

  const handleClick = () => {
    handleHeartClick();
    fixFavorite(
      {
        sayka_id: saykaId,
        is_favorite: liked,
      },
      token
    ).then(() => {
      console.log("success");
    });
  };

  if (!liked) {
    return (
      <div className="flex items-center gap-x-1">
        <VscHeart size={24} onClick={handleClick} />
        {likeCount}
      </div>
    );
  }

  return (
    <div className="flex items-center gap-x-1">
      <VscHeartFilled size={24} color="#2DD4BF" onClick={handleClick} />
      {likeCount}
    </div>
  );
};
