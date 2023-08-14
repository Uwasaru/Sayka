import { useRouter } from "next/router";

export const useRedirect = () => {
  const router = useRouter();

  const redirectToTimeline = () => {
    router.push("/timeline");
  };

  return {
    redirectToTimeline,
  };
};
