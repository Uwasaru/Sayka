import { redirect } from "next/navigation";

import { getLoggedInUser, mock_users, readUser } from "@/api";
import { getToken } from "@/features";

import { MypageSaykaList } from "./_components/MypageSaykaList";
import { Profile } from "./_components/Profile";

type TProps = {
  params: { user_id: string };
};

const Page = async ({ params }: TProps) => {
  const readUserRes = await readUser(params.user_id);
  console.log("readUserRes", readUserRes);
  if (readUserRes.type === "error") {
    redirect("/timeline");
  }
  console.log("readUser", readUserRes.value.data);
  const loggedUserRes = await getLoggedInUser(getToken() || "");
  const loggedInUser = loggedUserRes.type === "ok" && loggedUserRes.value.data;

  const isMe = loggedInUser && loggedInUser.id === readUserRes.value.data.id;
  // const isMe = true;

  return (
    <div className="grid grid-cols-5 space-y-5 md:p-5">
      <div className="col-span-5 mb-5 md:col-span-2 md:mb-0 md:mr-10">
        <Profile user={readUserRes.value.data} isMe={isMe} />
      </div>
      <div className="col-span-5 md:col-span-3">
        <MypageSaykaList userId={readUserRes.value.data.id} />
      </div>
    </div>
  );
};

export default Page;
