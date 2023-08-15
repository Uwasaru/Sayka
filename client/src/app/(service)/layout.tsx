import { getLoggedInUser } from "@/api";
import { getToken } from "@/features";
import { ServiceHeader } from "@/ui/Layout";

export default async function Layout({
  children,
}: {
  children: React.ReactNode;
}) {
  console.log("token", getToken());
  const user = await (async () => {
    const user = await getLoggedInUser(getToken() || "");
    console.log("user", user);
    if (user.type === "error") return undefined;
    console.log("user", user.value.user);
    return user.value.user;
  })();

  return (
    <div className="mb-16 bg-white">
      <ServiceHeader user={user} />
      <main>{children}</main>
    </div>
  );
}
