import { readSayka } from "@/api";

import { Sayka } from "../../_components/Sayka";
import { Metadata } from "next";
import { getEnv } from "@/utils/env";

type TProps = {
  params: { sayka_id: string };
};

const { clientURL } = getEnv();

export async function generateMetadata({ params }: TProps): Promise<Metadata> {
  const saykaRes = await readSayka(params.sayka_id);
  if (saykaRes.type === "error")
    throw new Error("データが取得できませんでした。");
  const sayka = saykaRes.value.data;
  if (!sayka) throw new Error("投稿がありません");

  return {
    openGraph: {
      title: sayka.title,
      description: sayka.description,
      url: `${clientURL}/shareSaykaInformation/${sayka.id}}`,
      siteName: "Sayka",
      images: [
        {
          url: `${clientURL}/og?title=${sayka.title}&userName=${sayka.user.id}`,
        },
      ],
      type: "website",
    },
  };
}

const Page = async ({ params }: TProps) => {
  const saykaRes = await readSayka(params.sayka_id);
  if (saykaRes.type === "error")
    throw new Error("データが取得できませんでした。");
  const sayka = saykaRes.value.data;
  if (!sayka) throw new Error("投稿がありません");
  return (
    <div className="mt-40">
      <Sayka sayka={sayka} />
    </div>
  );
};

export default Page;
