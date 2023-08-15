import { getLoggedInUser } from "@/api";
import { getToken } from "@/features";
import { ServiceHeader } from "@/ui/Layout";

export default async function Layout({
  children,
}: {
  children: React.ReactNode;
}) {
  const user = await (async () => {
    const user = await getLoggedInUser(getToken() || "");
    if (user.type === "error") return undefined;
    return user.value.data;
  })();

  return (
    <div className="mb-16 bg-white">
      <ServiceHeader user={user} />
      <main>{children}</main>
    </div>
  );
}
