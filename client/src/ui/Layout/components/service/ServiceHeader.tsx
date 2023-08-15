import { useAtom } from "jotai";
import Image from "next/image";
import Link from "next/link";
import { FC, useState } from "react";
import { LiaUserCircleSolid } from "react-icons/lia";

import { LoginModal } from "@/app/(service)/(audience)/_components/LoginModal";
import { modalState } from "@/store/atoms/modalAtom";
import { TUser } from "@/types/User";
import { ColorButton } from "@/ui/Button";
import { SaykaLogo } from "@/ui/svg/SaykaLogo";

type TProps = {
  user?: TUser;
};

export const ServiceHeader: FC<TProps> = ({ user }) => {
  const [isOpen, setIsOpen] = useAtom(modalState);
  const handleOpen = () => {
    setIsOpen(true);
  };

  if (!user) {
    return (
      <>
        <header className="flex h-[10vh] items-center justify-between p-5">
          <Link href="/timeline">
            <SaykaLogo />
          </Link>
          <div className="flex items-center space-x-5 ">
            <LiaUserCircleSolid size={34} className="hidden md:block" />
            <ColorButton label="ログイン" onClick={handleOpen} />
          </div>
        </header>
        {isOpen && <LoginModal />}
      </>
    );
  }
  return (
    <header className="flex h-[10vh] items-center justify-between p-5">
      <Link href="/timeline">
        <SaykaLogo />
      </Link>
      <div className="flex items-center space-x-5">
        <Link href={`/mypage/${user.id}`}>
          <Image
            src={user.img}
            alt="user icon"
            width={35}
            height={35}
            className="rounded-full"
          />
        </Link>
        <Link
          href="/new"
          className="rounded bg-sc px-4 py-2 font-semibold text-white transition duration-300 hover:bg-hover-sc">
          投稿する
        </Link>
      </div>
    </header>
  );
};
