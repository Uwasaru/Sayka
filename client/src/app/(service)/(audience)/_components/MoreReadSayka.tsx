import { readSaykaTimeline } from "@/api/sayka";
import { TUser } from "@/types/User";
import React from "react";
import { SaykaTimelineList } from "./SaykaTimelineList";

type TProps = {
  lastSaykaId: string;
  loginUser?: TUser;
};

export const MoreReadSayka = async ({ lastSaykaId, loginUser }: TProps) => {
  const saykasRes = await readSaykaTimeline(lastSaykaId);

  if (saykasRes.type === "error") return;

  const saykas = saykasRes.value.data;

  if (!saykas) return;

  return <SaykaTimelineList loginUser={loginUser} saykas={saykas} />;
};
