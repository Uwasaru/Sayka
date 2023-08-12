import Link from "next/link";
import { SlArrowLeft } from "react-icons/sl";

import { TooltipUI } from "@/ui/Tooltip";

export const CreativeNavigation = () => {
  return (
    <div className="border-t border-t-stone-900">
      <div className="flex justify-start px-10 py-5">
        <Link href="/timeline">
          <TooltipUI label="戻る" placement="bottom">
            <SlArrowLeft size={20} />
          </TooltipUI>
        </Link>
      </div>
    </div>
  );
};
