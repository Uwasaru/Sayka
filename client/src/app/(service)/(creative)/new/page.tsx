import { redirect } from "next/navigation";

import { getLoggedInUser } from "@/api";
import { getToken } from "@/features";
import { ContentTitle } from "@/ui/Text";

import { SaykaForm } from "../_components/SaykaForm";

const Page = async () => {
  const token = getToken();
  if (!token) return redirect("/timeline");
  const userRes = await getLoggedInUser(token || "");
  if (userRes.type === "error") return redirect("/timeline");
  const user = userRes.value.data;
  return (
    <div>
      <ContentTitle text="成果の投稿" />
      <div className="px-2 md:px-0">
        <SaykaForm user={user} token={token} />
      </div>
    </div>
  );
};

export default Page;
