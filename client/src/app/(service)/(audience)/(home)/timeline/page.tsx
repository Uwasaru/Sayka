import { getLoggedInUser, readUser } from "@/api";
import { readSayka, readSaykaTimeline } from "@/api/sayka";
import { getToken } from "@/features";

import { CommentModal } from "../../_components/CommentModal";
import { SaykaListWithFilter } from "../../_components/SaykaListWithFilter";
import { SaykaTimelineList } from "../../_components/SaykaTimelineList";

type TProps = {
  searchParams: { modal: boolean; sayka_id: string };
};

const Page = async ({ searchParams }: TProps) => {
  const isOpenModal = searchParams.modal;
  const saykaId = searchParams.sayka_id;

  const loginUser = await (async () => {
    const loginUser = await getLoggedInUser(getToken() || "");
    if (loginUser.type === "error") return undefined;
    return loginUser.value.data;
  })();

  const saykasRes = await readSaykaTimeline();
  if (saykasRes.type === "error") {
    throw new Error("データが取得できませんでした。");
  }
  const saykas = saykasRes.value.data;

  if (!saykas) {
    throw new Error("投稿がありません");
  }

  let targetSayka;
  let targetSaykaUser;

  if (saykaId) {
    const saykaRes = await readSayka(saykaId);
    if (saykaRes.type === "error") {
      throw new Error("データが取得できませんでした。");
    }
    targetSayka = saykaRes.value.data;
    const userRes = await readUser(targetSayka.user_id);
    if (userRes.type === "error") {
      throw new Error("データが取得できませんでした。");
    }
    targetSaykaUser = userRes.value.data;
  }

  return (
    <div className="mt-5 md:mx-16">
      <SaykaTimelineList loginUser={loginUser} saykas={saykas} />
      {isOpenModal && targetSayka && targetSaykaUser && (
        <CommentModal
          sayka={targetSayka}
          loginUser={loginUser}
          user={targetSaykaUser}
        />
      )}
    </div>
  );
};

export default Page;
