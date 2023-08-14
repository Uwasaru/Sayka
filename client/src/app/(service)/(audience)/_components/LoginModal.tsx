import { useAtom } from "jotai";
import Image from "next/image";
import { FC } from "react";
import { GrClose } from "react-icons/gr";



import { modalState } from "@/store/atoms/modalAtom";
import { SaykaLogo } from "@/ui/svg/SaykaLogo";
import { getEnv } from "@/utils/env";

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
    <div className="fixed inset-0 z-50 flex items-center justify-center">
      {/* Background overlay */}
      <div
        className="absolute inset-0 bg-black opacity-50 transition-opacity duration-300 ease-in-out"
        onClick={onClose}
      />

      {/* Modal content */}
      <div className="z-10 w-4/5 overflow-hidden rounded-xl bg-white shadow-2xl md:w-2/5 ">
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
              className="mb-10 flex items-center rounded-md border-2 border-gray-800 bg-white px-4 py-3 font-bold text-stone-900 hover:bg-sc hover:opacity-90">
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
