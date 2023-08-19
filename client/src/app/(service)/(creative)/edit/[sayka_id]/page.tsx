import { redirect } from "next/navigation";

import { readSayka } from "@/api";
import { getToken } from "@/features";
import { ContentTitle } from "@/ui/Text";

import { SaykaEditForm } from "../../_components/SaykaEditForm";

type TProps = {
  params: { sayka_id: string };
};

const Page = async ({ params }: TProps) => {
  const token = getToken();
  if (!token) return redirect("/timeline");
  const saykaRes = await readSayka(params.sayka_id);
  if (saykaRes.type === "error")
    throw new Error("データが取得できませんでした。");
  const sayka = saykaRes.value.data;
  if (!sayka) throw new Error("投稿がありません");
  return (
    <div>
      <ContentTitle text="成果の編集" />
      <div className="px-2 md:px-0">
        <SaykaEditForm sayka={sayka} token={token} />
      </div>
    </div>
  );
};

export default Page;
