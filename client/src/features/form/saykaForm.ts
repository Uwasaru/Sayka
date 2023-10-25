import { z } from "zod";

import { TSayka } from "@/types/Sayka";

export type SaykaFormData = {
  title: TSayka["title"];
  description: TSayka["description"];
  github_url?: TSayka["github_url"];
  figma_url?: TSayka["figma_url"];
  slide_url?: TSayka["slide_url"];
  article_url?: TSayka["article_url"];
  app_url?: TSayka["app_url"];
};

export const saykaSchema = z.object({
  title: z
    .string()
    .nonempty({ message: "必須項目です。" })
    .max(50, "50字以内で入力してください。"),
  description: z
    .string()
    .nonempty({ message: "必須項目です。" })
    .max(100, "100字以内で入力してください。"),
  github_url: z.union([
    z.literal(""),
    z.string().regex(/^https:\/\//, "URLを正しい形で入力してください。"),
  ]),
  figma_url: z.union([
    z.literal(""),
    z.string().regex(/^https:\/\//, "URLを正しい形で入力してください。"),
  ]),
  slide_url: z.union([
    z.literal(""),
    z.string().regex(/^https:\/\//, "URLを正しい形で入力してください。"),
  ]),
  article_url: z.union([
    z.literal(""),
    z.string().regex(/^https:\/\//, "URLを正しい形で入力してください。"),
  ]),
  app_url: z.union([
    z.literal(""),
    z.string().regex(/^https:\/\//, "URLを正しい形で入力してください。"),
  ]),
});
