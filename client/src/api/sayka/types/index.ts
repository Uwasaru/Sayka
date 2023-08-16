import { TSayka } from "@/types/Sayka";
import { JTDDataType } from "ajv/dist/core";

const saykaResponseBaseSchema = {
  type: "object",
  properties: {
    user_id: { type: "string" },
    id: { type: "string" },
    title: { type: "string" },
    description: { type: "string" },
    tags: {
      elements: {
        type: "string",
      },
    },
    github_url: { type: "string" },
    slide_url: { type: "string" },
    figma_url: { type: "string" },
    article_url: { type: "string" },
    app_url: { type: "string" },
    favorites: { type: "int8" },
    comments: { type: "int8" },
    is_favorite: { type: "boolean" },
    created_at: { type: "string" },
  },
} as const;

const saykaResponseSchema = {
  type: "object",
  properties: {
    data: saykaResponseBaseSchema,
  },
} as const;

export type SaykaResponse = JTDDataType<typeof saykaResponseSchema>;

export type SaykaBase = TSayka;

export type SaykaCreate = Omit<
  SaykaBase,
  "id" | "created_at" | "favorites" | "comments" | "is_favorite"
>;
