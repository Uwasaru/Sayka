import { readSayka } from "@/api";
import { ColorButton } from "@/ui/Button";
import { redirect } from "next/navigation";
import { Sayka } from "../../_components/Sayka";

type TProps = {
  params: { sayka_id: string };
};

const Page = async ({ params }: TProps) => {
  const saykaRes = await readSayka(params.sayka_id);
  if (saykaRes.type === "error")
    throw new Error("データが取得できませんでした。");
  const sayka = saykaRes.value.data;
  if (!sayka) throw new Error("投稿がありません");

  return (
    <div className="mt-40 space-y-5">
      <div className="flex justify-center text-lg font-bold">
        投稿が完了しました。
      </div>
      <Sayka sayka={sayka} />
    </div>
  );
};

export default Page;
