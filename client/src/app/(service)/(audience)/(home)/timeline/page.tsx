import { getLoggedInUser } from "@/api";
import { getToken } from "@/features";

import { CommentModal } from "../../_components/CommentModal";
import { SaykaListWithFilter } from "../../_components/SaykaListWithFilter";

type TProps = {
  searchParams: { modal: boolean; sayka_id: string };
};

const Page = async ({ searchParams }: TProps) => {
  const isOpenModal = searchParams.modal;
  const saykaId = searchParams.sayka_id;

  const user = await (async () => {
    const user = await getLoggedInUser(getToken() || "");
    if (user.type === "error") return undefined;
    return user.value.user;
  })();

  return (
    <div className="mt-5 md:mx-16">
      <SaykaListWithFilter user={user} />
      {isOpenModal && saykaId && (
        <CommentModal id={Number(saykaId)} user={user} />
      )}
    </div>
  );
};

export default Page;
