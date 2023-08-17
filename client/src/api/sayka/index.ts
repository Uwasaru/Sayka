import { getEnv } from "@/utils/env";
import { apiClient } from "../core";
import { SaykaCreate, SaykaResponse, SaykasResponse } from "./types";

const { serverURL } = getEnv();
export const createSayka = async (sayka: SaykaCreate) =>
  await apiClient.post<SaykaResponse>(`${serverURL}/post/`, sayka);

export const readSaykaTimeline = async (last_sayka_id?: string) => {
  let endPoint;
  if (last_sayka_id) {
    endPoint = `${serverURL}/post/timeline/${last_sayka_id}`;
  } else {
    endPoint = `${serverURL}/post/timeline/init`;
  }

  return await apiClient.get<SaykasResponse>(endPoint);
};
