export type TSayka = {
  id: string;
  user_id: string;
  title: string;
  description: string;
  github_url?: string;
  slide_url?: string;
  figma_url?: string;
  article_url?: string;
  app_url?: string;
  tags?: string[];
  favorites: number;
  comments: number;
  is_favorite: boolean;
  created_at: string;
};
