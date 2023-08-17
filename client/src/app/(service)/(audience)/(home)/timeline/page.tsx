import { getLoggedInUser, readUser } from "@/api";
import { readSaykaTimeline } from "@/api/sayka";
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

  const saykaRes = await readSaykaTimeline();
  if (saykaRes.type === "error") {
    throw new Error("データが取得できませんでした。");
  }
  const saykas = saykaRes.value.data;

  return (
    <div className="mt-5 md:mx-16">
      <SaykaTimelineList loginUser={loginUser} saykas={saykas} />
      {isOpenModal && saykaId && (
        <CommentModal id={Number(saykaId)} loginUser={loginUser} />
      )}
    </div>
  );
};

export default Page;
