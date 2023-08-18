import { getEnv } from "@/utils/env";

import { apiClient } from "../core";

import { FavoriteCreate, FavoriteResponse } from "./types";

const { serverURL } = getEnv();

export const fixFavorite = async (favorite: FavoriteCreate, token: string) => {
  await apiClient.post<FavoriteResponse>(
    `${serverURL}/favorite`,
    favorite,
    token
  );
};
