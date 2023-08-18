import { getToken } from "@/features";
import { getEnv } from "@/utils/env";

import { apiClient } from "../core";

import { CommentCreate, CommentResponse } from "./types";

const { serverURL } = getEnv();

export const createComment = async (comment: CommentCreate, token: string) => {
  await apiClient.post<CommentResponse>(`${serverURL}/comment`, comment, token);
};
