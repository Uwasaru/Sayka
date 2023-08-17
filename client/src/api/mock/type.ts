import { TUser } from "@/types/User";

export type TTag = string;

export type TSayka = {
  id: number;
  title: string;
  description: string;
  like_count: number;
  comment_count: number;
  github_url?: string;
  slide_url?: string;
  figma_url?: string;
  article_url?: string;
  application_url?: string;
  user: TUser;
  tags?: TTag[];
};

export type TComment = {
  id: number;
  contents: string;
  created_at: string;
  user: TUser;
};
