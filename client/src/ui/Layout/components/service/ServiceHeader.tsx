import Link from "next/link";

import { SaykaLogo } from "@/ui/svg/SaykaLogo";

export const ServiceHeader = () => {
  return (
    <header className="flex h-[10vh] items-center justify-between p-5">
      <Link href="/timeline">
        <SaykaLogo />
      </Link>
      <div>
        <Link
          href="/new"
          className="rounded bg-sc px-4 py-2 font-semibold text-white transition duration-300 hover:bg-hover-sc">
          投稿する
        </Link>
      </div>
    </header>
  );
};
