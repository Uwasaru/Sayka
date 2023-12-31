import { getEnv } from "@/utils/env";

import { apiClient } from "../core";

import { UserCreate, UserResponse } from "./types";

const { serverURL } = getEnv();
export const createUser = async (user: UserCreate) =>
  await apiClient.post<UserResponse>(`${serverURL}/user`, user);

export const readUser = async (user_id: string) =>
  await apiClient.get<UserResponse>(`${serverURL}/user/${user_id}`);

export const deleteUser = async (user_id: string) =>
  await apiClient.delete(`${serverURL}/user/${user_id}`);
