"use client";

import { Sayka } from "../../_components/Sayka";

type TProps = {
  params: { sayka_id: string };
};

const Page = ({ params }: TProps) => {
  const sayka = null;
  if (!sayka) {
    return <div>not found</div>;
  }
  return (
    <div className="mt-40">
      <Sayka sayka={sayka} />
    </div>
  );
};

export default Page;
