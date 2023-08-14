import { SaykaList } from "../../_components/SaykaList";
import { MypageSaykaList } from "./_components/MypageSaykaList";

import { Profile } from "./_components/Profile";

const Page = () => {
  return (
    <div className="grid grid-cols-5 md:p-5 space-y-5">
      <div className="md:col-span-2 col-span-5 md:mr-10 mb-5 md:mb-0">
        <Profile
          user={{
            id: 1,
            name: "39TO",
            icon: "/icon.jpg",
            like_count: 247,
            sayka_count: 3,
          }}
        />
      </div>
      <div className="col-span-5 md:col-span-3">
        <MypageSaykaList />
      </div>
    </div>
  );
};

export default Page;
