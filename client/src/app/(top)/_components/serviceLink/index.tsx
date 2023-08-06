import Image from "next/image";

export const ServiceLink = () => {
  return (
    <div className="space-y-10">
      <a
        href="/"
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
      <a
        href="/home/timeline"
        className="flex items-center rounded-md bg-blue-700 px-4 py-3 font-bold text-white hover:opacity-90">
        ログインせずに利用する
      </a>
    </div>
  );
};
