import { redirect } from "next/navigation";
import { Suspense } from "react";
import { ErrorBoundary } from "react-error-boundary";

import { readProfile, readSayka } from "@/api";
import { getToken } from "@/features";

import { CommentModal } from "../../_components/CommentModal";

import { ErrorFallback } from "./_components/ErrorFallback";
import { LoadingFallback } from "./_components/LoadingFallback";
import { MypageNavigation } from "./_components/MypageNavigation";
import { MyselfSaykaList } from "./_components/MyselfSaykaList";
import { Profile } from "./_components/Profile";

type TProps = {
  params: { user_id: string };
  searchParams: { modal: boolean; sayka_id: string };
};

const Page = async ({ params, searchParams }: TProps) => {
  const isOpenModal = searchParams.modal;
  const saykaId = searchParams.sayka_id;
  const token = getToken();
  const resProfile = await readProfile(params.user_id, token);
  if (resProfile.type === "error") {
    redirect("/timeline");
  }
  const profile = resProfile.value.data;

  let targetSayka;
  if (saykaId) {
    const saykaRes = await readSayka(saykaId);
    if (saykaRes.type === "error") {
      throw new Error("データが取得できませんでした。");
    }
    targetSayka = saykaRes.value.data;
  }
  return (
    <>
      <div className="grid grid-cols-5 space-y-5 md:p-5">
        <div className="col-span-5 mb-5 md:col-span-2 md:mb-0 md:mr-10">
          <Profile profile={profile} />
        </div>
        <div className="col-span-5 md:col-span-3">
          <MypageNavigation userId={profile.user.id} feature="myself" />
          <ErrorBoundary FallbackComponent={ErrorFallback}>
            <Suspense fallback={<LoadingFallback />}>
              {/* @ts-expect-error Server Component */}
              <MyselfSaykaList userId={profile.user.id} />
            </Suspense>
          </ErrorBoundary>
        </div>
      </div>
      {isOpenModal && targetSayka && (
        <CommentModal sayka={targetSayka} token={token} />
      )}
    </>
  );
};

export default Page;

const saykaError = () => {
  return <div>sayka error</div>;
};
