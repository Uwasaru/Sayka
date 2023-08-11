import { SaykaLogo } from "@/ui/svg/SaykaLogo";
import Link from "next/link";

export const ServiceHeader = () => {
  return (
    <header className="flex justify-between items-center h-[10vh] p-5">
      <SaykaLogo />
      <div>
        <Link
          href="/new"
          className="bg-sc hover:bg-hover-sc text-white px-4 py-2 rounded transition duration-300 font-semibold">
          投稿する
        </Link>
      </div>
    </header>
  );
};
