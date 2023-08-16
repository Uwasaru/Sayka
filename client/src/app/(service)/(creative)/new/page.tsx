import { getLoggedInUser } from "@/api";
import { getToken } from "@/features";
import { ContentTitle } from "@/ui/Text";
import { redirect } from "next/navigation";

import { SaykaForm } from "../_components/SaykaForm";

const Page = async () => {
  const userRes = await getLoggedInUser(getToken() || "");
  if (userRes.type === "error") return redirect("/timeline");
  const user = userRes.value.data;
  return (
    <div>
      <ContentTitle text="成果の投稿" />
      <div className="px-2 md:px-0">
        <SaykaForm user={user} />
      </div>
    </div>
  );
};

export default Page;
