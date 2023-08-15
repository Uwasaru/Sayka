// import { Auth } from "@/features";
import { Auth } from "@/features";
import { CreativeNavigation } from "./_components/CreativeNavigation";

export default async function Layout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <>
      {/* @ts-expect-error Server Component */}
      <Auth>
        <div>
          <CreativeNavigation />
          <div className="grid grid-cols-9">
            <div className="hidden md:col-span-3 md:block" />
            <main className="col-span-10 md:col-span-3">{children}</main>
            <div className="hidden md:col-span-3 md:block" />
          </div>
        </div>
      </Auth>
    </>
  );
}
