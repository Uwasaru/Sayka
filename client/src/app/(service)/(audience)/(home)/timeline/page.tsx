"use client";

import { useSearchParams } from "next/navigation";

import { CommentModal } from "../../_components/CommentModal";
import { SaykaList } from "../../_components/SaykaList";

const Page = () => {
  const searchParams = useSearchParams();
  const isOpenModal = searchParams.get("modal");
  const saykaId = searchParams.get("sayka_id");

  console.log(isOpenModal);
  console.log(saykaId);

  return (
    <div className="mt-5 md:mx-16">
      <SaykaList />
      {isOpenModal && saykaId && <CommentModal id={Number(saykaId)} />}
    </div>
  );
};

export default Page;
