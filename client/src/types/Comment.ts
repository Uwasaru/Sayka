export type TComment = {
  id: string;
  user_id: string;
  content: string;
  created_at: string;
  user: {
    id: string;
    name: string;
    img: string;
  };
};
