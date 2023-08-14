"use client";

import { mock_users } from "@/api";
import { useRouter } from "next/navigation";
import { MypageSaykaList } from "./_components/MypageSaykaList";
import { Profile } from "./_components/Profile";

type TProps = {
  params: { user_id: string };
};

const Page = ({ params }: TProps) => {
  const router = useRouter();
  const user = mock_users.find((user) => user.id === Number(params.user_id));
  if (!user) {
    router.push("/timeline");
    return null;
  }

  return (
    <div className="grid grid-cols-5 space-y-5 md:p-5">
      <div className="col-span-5 mb-5 md:col-span-2 md:mb-0 md:mr-10">
        <Profile user={user} />
      </div>
      <div className="col-span-5 md:col-span-3">
        <MypageSaykaList userId={user.id} />
      </div>
    </div>
  );
};

export default Page;
