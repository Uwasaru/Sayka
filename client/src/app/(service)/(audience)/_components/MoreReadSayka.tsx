import React from "react";

import { readSaykaTimeline } from "@/api/sayka";
import { TUser } from "@/types/User";

import { SaykaTimelineList } from "./SaykaTimelineList";

type TProps = {
  lastSaykaId: string;
};

export const MoreReadSayka = async ({ lastSaykaId }: TProps) => {
  const saykasRes = await readSaykaTimeline(lastSaykaId);

  if (saykasRes.type === "error") return;

  const saykas = saykasRes.value.data;

  if (!saykas) return;

  return <SaykaTimelineList saykas={saykas} />;
};
