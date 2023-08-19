import { getEnv } from "@/utils/env";

import { apiClient } from "../core";

import { AuthUserResponse } from "./types";

const { serverURL } = getEnv();
export const getLoggedInUser = async (token: string) =>
  await apiClient.get<AuthUserResponse>(
    `${process.env.NEXT_PUBLIC_SERVER_URL}/auth/user`,
    token
  );
