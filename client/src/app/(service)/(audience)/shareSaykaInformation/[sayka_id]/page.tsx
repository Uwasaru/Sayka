"use client";

import { mock_saykas } from "@/api";

import { Sayka } from "../../_components/Sayka";

type TProps = {
  params: { sayka_id: string };
};

const Page = ({ params }: TProps) => {
  const sayka = mock_saykas.find(
    (sayka) => sayka.id === Number(params.sayka_id)
  );

  if (!sayka) {
    return <div>not found</div>;
  }
  return (
    <div className="mt-40">
      <Sayka data={sayka} />
    </div>
  );
};

export default Page;
