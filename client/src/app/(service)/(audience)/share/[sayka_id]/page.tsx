"use client";

import { Sayka } from "../../_components/Sayka";
import { SaykaMockData } from "../../_components/SaykaList";

type TProps = {
  params: { sayka_id: string };
};

const Page = ({ params }: TProps) => {
  const sayka = SaykaMockData.find(
    (sayka) => sayka.id === Number(params.sayka_id)
  );

  if (!sayka) {
    return <div>404</div>;
  }
  return (
    <div className="mt-5">
      <Sayka data={sayka} />
    </div>
  );
};

export default Page;
