import Image from "next/image";
import Link from "next/link";

import { getEnv } from "@/utils/env";

const { serverURL, clientURL } = getEnv();
export const ServiceLink = () => {
  const authUrl = `${serverURL}/auth/login?redirect_url=${clientURL}/timeline`;
  return (
    <div className="space-y-10">
      <a
        href={authUrl}
        className="flex items-center rounded-md bg-white px-4 py-3 font-bold text-stone-900 hover:opacity-90">
        <Image
          src="/github-logo.svg"
          alt="github"
          width={30}
          height={30}
          className="mr-2"
        />
        GitHubでログイン
      </a>
      <Link
        href="/timeline"
        className="flex items-center rounded-md bg-blue-700 px-4 py-3 font-bold text-white hover:opacity-90">
        ログインせずに利用する
      </Link>
    </div>
  );
};
