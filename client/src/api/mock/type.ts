export type TUser = {
  id: number;
  name: string;
  icon: string;
  like_count: number;
  sayka_count: number;
};

export type TTag = {
  id: number;
  name: string;
};

export type TSayka = {
  id: number;
  title: string;
  description: string;
  like_count: number;
  comment_count: number;
  github_url?: string;
  slide_url?: string;
  application_url?: string;
  user: TUser;
  tags?: TTag[];
};
