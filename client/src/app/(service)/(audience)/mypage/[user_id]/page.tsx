import { SaykaList } from "../../_components/SaykaList";
import { Profile } from "./_components/Profile";

const Page = () => {
  return (
    <div className="grid grid-cols-5 p-5">
      <div className="hidden md:col-span-2 md:block">
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
      <div className="col-span-10 md:col-span-3">
        <SaykaList />
      </div>
    </div>
  );
};

export default Page;
