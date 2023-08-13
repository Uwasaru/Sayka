import Image from "next/image";
import Link from "next/link";

import { SaykaLogo } from "@/ui/svg/SaykaLogo";

export const ServiceHeader = () => {
  const user_id = 1;
  const user = {
    id: 1,
    name: "39TO",
    icon: "/icon.jpg",
  };
  return (
    <header className="flex h-[10vh] items-center justify-between p-5">
      <Link href="/timeline">
        <SaykaLogo />
      </Link>
      <div className="flex items-center space-x-5">
        <Link href={`/mypage/${user_id}`}>
          <Image
            src={user.icon}
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
