import { useAtom } from "jotai";
import React, { FC, useState } from "react";
import { VscHeartFilled } from "react-icons/vsc";
import { VscHeart } from "react-icons/vsc";

import { modalState } from "@/store/atoms/modalAtom";

type TProps = {
  saykaId: number;
};

export const LikeButton: FC<TProps> = ({ saykaId }) => {
  // const user = await (async () => {
  //   const user = await getLoggedInUser(getToken() || '')
  //   if (user.type === 'error') return undefined
  //   return user.value.user
  // })()
  const user = null;
  const [liked, setLiked] = useState(false);

  const handleHeartClick = () => {
    setLiked(!liked);
  };

  const [_, setIsOpen] = useAtom(modalState);
  const handleHeartClickByGuest = () => {
    setIsOpen(true);
  };

  if (!user) {
    return (
      <div>
        <VscHeart size={24} onClick={handleHeartClickByGuest} />
      </div>
    );
  }

  if (!liked) {
    return <VscHeart size={24} onClick={handleHeartClick} />;
  }

  return (
    <VscHeartFilled size={24} color="#2DD4BF" onClick={handleHeartClick} />
  );
};
