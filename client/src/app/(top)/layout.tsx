import { redirect } from "next/navigation";

import { getLoggedInUser } from "@/api";
import { getToken } from "@/features";
import { TopHeader, TopMain } from "@/ui/Layout";

export default async function Layout({
  children,
}: {
  children: React.ReactNode;
}) {
  const token = getToken();
  const user = await getLoggedInUser(token || "");
  if (user.type === "ok") return redirect("/timeline");

  return (
    <div className="bg-white">
      <TopHeader />
      <TopMain>{children}</TopMain>
    </div>
  );
}
