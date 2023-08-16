import { getEnv } from "@/utils/env";
import { apiClient } from "../core";
import { SaykaCreate, SaykaResponse, SaykaTimelineRequest } from "./types";

const { serverURL } = getEnv();
export const createSayka = async (sayka: SaykaCreate) =>
  await apiClient.post<SaykaResponse>(`${serverURL}/post/`, sayka);
