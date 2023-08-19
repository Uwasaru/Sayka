import { redirect } from "next/navigation";

import { getToken } from "@/features";

import { MypageSaykaList } from "./_components/MypageSaykaList";
import { Profile } from "./_components/Profile";
import { MypageNavigation } from "./_components/MypageNavigation";
import { readProfile } from "@/api";
import { MyselfSaykaList } from "./_components/MyselfSaykaList";

type TProps = {
  params: { user_id: string };
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
        <MypageNavigation userId={profile.user.id} feature="myself" />
        {/* @ts-expect-error Server Component */}
        <MyselfSaykaList userId={profile.user.id} />
      </div>
    </div>
  );
};

export default Page;

const saykaError = () => {
  return <div>sayka error</div>;
};
