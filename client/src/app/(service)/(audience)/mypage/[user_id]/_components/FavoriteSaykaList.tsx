import { readFavoriteSaykas } from "@/api";
import { getToken } from "@/features";
import { TUser } from "@/types/User";

import { SaykaList } from "../../../_components/SaykaList";

type TProps = {
  userId: TUser["id"];
};

export const FavoriteSaykaList = async ({ userId }: TProps) => {
  const token = getToken();
  const saykasRes = await readFavoriteSaykas(userId, token);
  if (saykasRes.type === "error")
    throw new Error("データが取得できませんでした。");
  const saykas = saykasRes.value.data;
  if (!saykas) throw new Error("投稿がありません");

  return <SaykaList saykas={saykas} token={token} />;
};
