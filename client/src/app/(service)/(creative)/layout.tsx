import { redirect } from "next/navigation";

import { getLoggedInUser } from "@/api";
import { getToken } from "@/features";

import { CreativeNavigation } from "./_components/CreativeNavigation";

export default async function Layout({
  children,
}: {
  children: React.ReactNode;
}) {
  const token = getToken();
  if (!token) return redirect("/");
  const user = await (async () => {
    const user = await getLoggedInUser(token || "");
    if (user.type === "error") return undefined;
    return user.value.data;
  })();
  if (!user) return redirect("/");

  return (
    <div>
      <CreativeNavigation />
      <div className="grid grid-cols-9">
        <div className="hidden md:col-span-3 md:block" />
        <main className="col-span-10 md:col-span-3">{children}</main>
        <div className="hidden md:col-span-3 md:block" />
      </div>
    </div>
  );
}
