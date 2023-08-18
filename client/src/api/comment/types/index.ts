import { JTDDataType } from "ajv/dist/core";

const commentBaseSchema = {
  type: "object",
  properties: {
    id: { type: "string" },
    sayka_id: { type: "string" },
    user_id: { type: "string" },
    content: { type: "string" },
    created_at: { type: "string" },
  },
} as const;

const commentResponseSchema = {
  type: "object",
  properties: {
    data: commentBaseSchema,
  },
} as const;

const commentsResponseSchema = {
  type: "object",
  properties: {
    data: {
      elements: commentBaseSchema,
    },
  },
} as const;

export type CommentResponse = JTDDataType<typeof commentResponseSchema>;

export type CommentsResponse = JTDDataType<typeof commentsResponseSchema>;

export type CommentCreate = {
  sayka_id: string;
  content: string;
};
