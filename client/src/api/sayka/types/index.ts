import { TSayka } from "@/types/Sayka";
import { JTDDataType } from "ajv/dist/core";

const saykaResponseBaseSchema = {
  type: "object",
  properties: {
    id: { type: "string" },
    user_id: { type: "string" },
    title: { type: "string" },
    github_url: { type: "string" },
    app_url: { type: "string" },
    slide_url: { type: "string" },
    figma_url: { type: "string" },
    description: { type: "string" },
    article_url: { type: "string" },
    favorites: { type: "int32" },
    comments: { type: "int32" },
    is_favorite: { type: "boolean" },
    created_at: { type: "string" },
    tags: {
      elements: {
        type: "string",
      },
    },
    user: {
      type: "object",
      properties: {
        id: { type: "string" },
        name: { type: "string" },
        img: { type: "string" },
      },
    },
  },
  required: [
    "id",
    "user_id",
    "title",
    "github_url",
    "app_url",
    "slide_url",
    "figma_url",
    "description",
    "article_url",
    "favorites",
    "comments",
    "is_favorite",
    "created_at",
    "user",
  ],
  additionalProperties: false,
} as const;

const saykaResponseSchema = {
  type: "object",
  properties: {
    data: saykaResponseBaseSchema,
  },
} as const;

export type SaykaResponse = JTDDataType<typeof saykaResponseSchema>;

const saykasResponseSchema = {
  type: "object",
  properties: {
    data: {
      elements: saykaResponseBaseSchema,
    },
  },
} as const;

export type SaykasResponse = JTDDataType<typeof saykasResponseSchema>;

export type SaykaBase = TSayka;

export type SaykaCreate = Omit<
  SaykaBase,
  "id" | "created_at" | "favorites" | "comments" | "is_favorite" | "user"
>;
