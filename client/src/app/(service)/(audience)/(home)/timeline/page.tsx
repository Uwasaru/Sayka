"use client";

import { useSearchParams } from "next/navigation";

import { CommentModal } from "../../_components/CommentModal";
import { SaykaListWithFilter } from "../../_components/SaykaListWithFilter";

const Page = () => {
  const searchParams = useSearchParams();
  const isOpenModal = searchParams.get("modal");
  const saykaId = searchParams.get("sayka_id");

  return (
    <div className="mt-5 md:mx-16">
      <SaykaListWithFilter />
      {isOpenModal && saykaId && <CommentModal id={Number(saykaId)} />}
    </div>
  );
};

export default Page;
