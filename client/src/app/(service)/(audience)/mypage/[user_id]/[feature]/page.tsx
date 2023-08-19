import { redirect } from "next/navigation";

import { getToken } from "@/features";

import { readProfile } from "@/api";
import { Profile } from "../_components/Profile";
import { MypageNavigation } from "../_components/MypageNavigation";
import { FavoriteSaykaList } from "../_components/FavoriteSaykaList";
import { CommentSaykaList } from "../_components/CommentSaykaList";

type TProps = {
  params: { user_id: string; feature: "favorite" | "comment" };
};

const Page = async ({ params }: TProps) => {
  const token = getToken();
  const resProfile = await readProfile(params.user_id, token);
  if (resProfile.type === "error") {
    redirect("/timeline");
  }
  const profile = resProfile.value.data;
  return (
    <div className="grid grid-cols-5 space-y-5 md:p-5">
      <div className="col-span-5 mb-5 md:col-span-2 md:mb-0 md:mr-10">
        <Profile profile={profile} />
      </div>
      <div className="col-span-5 md:col-span-3">
        <MypageNavigation userId={profile.user.id} feature={params.feature} />
        {params.feature === "favorite" ? (
          // @ts-expect-error Server Component
          <FavoriteSaykaList userId={profile.user.id} />
        ) : (
          // @ts-expect-error Server Component
          <CommentSaykaList userId={profile.user.id} />
        )}
      </div>
    </div>
  );
};

export default Page;
