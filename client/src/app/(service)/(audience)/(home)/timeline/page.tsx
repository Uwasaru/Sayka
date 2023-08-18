import { getLoggedInUser, readUser } from "@/api";
import { readSayka, readSaykaTimeline } from "@/api/sayka";
import { getToken } from "@/features";

import { CommentModal } from "../../_components/CommentModal";
import { SaykaTimelineList } from "../../_components/SaykaTimelineList";

type TProps = {
  searchParams: { modal: boolean; sayka_id: string };
};

const Page = async ({ searchParams }: TProps) => {
  const isOpenModal = searchParams.modal;
  const saykaId = searchParams.sayka_id;
  const token = getToken() || "";

  const saykasRes = await readSaykaTimeline(undefined, token);
  if (saykasRes.type === "error") {
    throw new Error("データが取得できませんでした。");
  }
  const saykas = saykasRes.value.data;

  if (!saykas) {
    throw new Error("投稿がありません");
  }

  let targetSayka;

  if (saykaId) {
    const saykaRes = await readSayka(saykaId);
    if (saykaRes.type === "error") {
      throw new Error("データが取得できませんでした。");
    }
    targetSayka = saykaRes.value.data;
  }

  return (
    <div className="mt-5 md:mx-16">
      <SaykaTimelineList saykas={saykas} token={token} />
      {isOpenModal && targetSayka && (
        <CommentModal sayka={targetSayka} token={token} />
      )}
    </div>
  );
};

export default Page;
