import { getEnv } from "@/utils/env";

import { apiClient } from "../core";

import { CommentCreate, CommentResponse, CommentsResponse } from "./types";

const { serverURL } = getEnv();

export const createComment = async (comment: CommentCreate, token: string) => {
  await apiClient.post<CommentResponse>(
    `${serverURL}/comment/`,
    comment,
    token
  );
};

export const readCommentBySayka = async (sayka_id: string) =>
  await apiClient.get<CommentsResponse>(
    `${serverURL}/comment/sayka/${sayka_id}`
  );
