import { JTDDataType } from "ajv/dist/core";

const commentSchema = {
  type: "object",
  properties: {
    data: {
      type: "object",
      properties: {
        id: { type: "string" },
        sayka_id: { type: "string" },
        user_id: { type: "string" },
        content: { type: "string" },
        created_at: { type: "string" },
      },
    },
  },
} as const;

export type CommentResponse = JTDDataType<typeof commentSchema>;

export type CommentCreate = {
  sayka_id: string;
  content: string;
};
