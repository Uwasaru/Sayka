"use client";
import { SlArrowLeft } from "react-icons/sl";

import { TooltipUI } from "@/ui/Tooltip";
import { useRouter } from "next/navigation";

export const CreativeNavigation = () => {
  const router = useRouter();
  const handleBack = () => {
    router.back();
  };
  return (
    <div className="border-t border-t-stone-900">
      <div className="flex justify-start px-10 py-5">
        <TooltipUI label="戻る" placement="bottom">
          <SlArrowLeft size={20} onClick={handleBack} />
        </TooltipUI>
      </div>
    </div>
  );
};
