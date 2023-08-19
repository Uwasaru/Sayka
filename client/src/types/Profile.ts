import { TUser } from "./User";

export type TProfile = {
  user: TUser;
  sayka_count: number;
  favorited_count: number;
  comment_count: number;
  is_me: boolean;
};
