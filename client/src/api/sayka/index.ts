import { getEnv } from "@/utils/env";

import { apiClient } from "../core";

import {
  DeleteSaykaResponse,
  SaykaCreate,
  SaykaResponse,
  SaykasResponse,
} from "./types";

const { serverURL } = getEnv();
export const createSayka = async (sayka: SaykaCreate, token: string) =>
  await apiClient.post<SaykaResponse>(`${serverURL}/sayka/`, sayka, token);

export const readSaykaTimeline = async (
  last_sayka_id?: string,
  token?: string
) => {
  let endPoint;
  if (last_sayka_id) {
    endPoint = `${serverURL}/sayka/timeline?id=${last_sayka_id}`;
  } else {
    endPoint = `${serverURL}/sayka/timeline/?id=init`;
  }

  return await apiClient.get<SaykasResponse>(endPoint, token);
};

export const readSayka = async (sayka_id: string) =>
  await apiClient.get<SaykaResponse>(`${serverURL}/sayka/${sayka_id}`);

export const updateSayka = async (
  sayka_id: string,
  sayka: SaykaCreate,
  token: string
) =>
  await apiClient.put<SaykaResponse>(
    `${serverURL}/sayka/${sayka_id}`,
    sayka,
    token
  );

export const deleteSayka = async (sayka_id: string, token: string) =>
  await apiClient.delete<DeleteSaykaResponse>(
    `${serverURL}/sayka/${sayka_id}`,
    undefined,
    token
  );
