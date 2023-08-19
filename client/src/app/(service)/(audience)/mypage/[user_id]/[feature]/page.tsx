import { redirect } from "next/navigation";
import { Suspense } from "react";
import { ErrorBoundary } from "react-error-boundary";

import { readProfile } from "@/api";
import { getToken } from "@/features";

import { CommentSaykaList } from "../_components/CommentSaykaList";
import { ErrorFallback } from "../_components/ErrorFallback";
import { FavoriteSaykaList } from "../_components/FavoriteSaykaList";
import { LoadingFallback } from "../_components/LoadingFallback";
import { MypageNavigation } from "../_components/MypageNavigation";
import { Profile } from "../_components/Profile";

type TProps = {
  params: { user_id: string; feature: "favorite" | "comment" };
};

const Page = async ({ params }: TProps) => {
  const token = getToken();
  const resProfile = await readProfile(params.user_id, token);
  if (resProfile.type === "error") {
    redirect("/timeline");
  }
  const profile = resProfile.value.data;
  return (
    <div className="grid grid-cols-5 space-y-5 md:p-5">
      <div className="col-span-5 mb-5 md:col-span-2 md:mb-0 md:mr-10">
        <Profile profile={profile} />
      </div>
      <div className="col-span-5 md:col-span-3">
        <MypageNavigation userId={params.user_id} feature={params.feature} />
        <ErrorBoundary FallbackComponent={ErrorFallback}>
          <Suspense fallback={<LoadingFallback />}>
            {params.feature === "favorite" ? (
              // @ts-expect-error Server Component
              <FavoriteSaykaList userId={params.user_id} />
            ) : (
              //   @ts-expect-error Server Component
              <CommentSaykaList userId={params.user_id} />
            )}
          </Suspense>
        </ErrorBoundary>
      </div>
    </div>
  );
};

export default Page;
