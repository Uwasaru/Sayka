import { readCommentSayka, readFavoriteSayka, readSaykaByUser } from "@/api";
import { getToken } from "@/features";
import { TUser } from "@/types/User";
import { SaykaList } from "../../../_components/SaykaList";

type TProps = {
  userId: TUser["id"];
};

export const CommentSaykaList = async ({ userId }: TProps) => {
  const token = getToken();
  const saykasRes = await readCommentSayka(userId, token);
  if (saykasRes.type === "error") return null;
  const saykas = saykasRes.value.data;
  if (!saykas) return null;

  return <SaykaList saykas={saykas} />;
};
