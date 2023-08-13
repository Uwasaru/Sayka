import { SaykaLogo } from "@/ui/svg/SaykaLogo";
import Image from "next/image";
import { FC } from "react";
import { GrClose } from "react-icons/gr";

import { getEnv } from "@/utils/env";
import { useAtom } from "jotai";
import { modalState } from "@/store/atoms/modalAtom";

const { serverURL, clientURL } = getEnv();

type TProps = {
  isVisible: boolean;
  onClose: () => void;
};

export const LoginModal: FC = () => {
  const authUrl = `${serverURL}/auth/login?redirect_url=${clientURL}/timeline`;
  const [isOpen, setIsOpen] = useAtom(modalState);
  const onClose = () => {
    setIsOpen(false);
  };

  if (!isOpen) return null;
  return (
    <div className="fixed inset-0 z-50 flex justify-center items-center">
      {/* Background overlay */}
      <div
        className="absolute inset-0 bg-black opacity-50 transition-opacity duration-300 ease-in-out"
        onClick={onClose}
      />

      {/* Modal content */}
      <div className="z-10 w-2/5 bg-white rounded-xl shadow-2xl overflow-hidden">
        {/* Header */}
        <div className="p-5">
          <div className="flex justify-end">
            <GrClose
              size={33}
              onClick={onClose}
              className="rounded-full p-2 transition-colors duration-300 hover:bg-teal-400"
            />
          </div>
        </div>
        {/* Body */}
        <div className="space-y-10">
          <div className="flex justify-center">
            <SaykaLogo />
          </div>
          <div className="flex justify-center">
            <a
              href={authUrl}
              className="flex items-center rounded-md bg-white px-4 py-3 font-bold text-stone-900 hover:opacity-90 border-2 border-gray-800 hover:bg-sc mb-10">
              <Image
                src="/github-logo.svg"
                alt="github"
                width={30}
                height={30}
                className="mr-2"
              />
              GitHubでログイン
            </a>
          </div>
        </div>
      </div>
    </div>
  );
};
