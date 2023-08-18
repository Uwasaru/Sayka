import { readSayka } from "@/api";
import { Sayka } from "../../_components/Sayka";
import { ShareLongButton } from "../../_components/ShareButton";

type TProps = {
  params: { sayka_id: string };
};

const Page = async ({ params }: TProps) => {
  const saykaRes = await readSayka(params.sayka_id);
  if (saykaRes.type === "error")
    throw new Error("データが取得できませんでした。");
  const sayka = saykaRes.value.data;
  if (!sayka) throw new Error("投稿がありません");

  console.log(sayka);

  return (
    <div className="mt-32 space-y-5">
      <div className="flex justify-center text-lg font-bold">
        投稿が完了しました!
      </div>
      <Sayka sayka={sayka} />
      <div className="flex justify-end">
        <ShareLongButton saykaId={params.sayka_id} saykaTitle={sayka.title} />
      </div>
    </div>
  );
};

export default Page;
