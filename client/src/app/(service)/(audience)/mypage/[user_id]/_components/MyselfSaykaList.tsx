import { readSaykaByUser } from "@/api";
import { TUser } from "@/types/User";
import { SaykaList } from "../../../_components/SaykaList";

type TProps = {
  userId: TUser["id"];
};

export const MyselfSaykaList = async ({ userId }: TProps) => {
  const saykasRes = await readSaykaByUser(userId);
  if (saykasRes.type === "error") return null;
  const saykas = saykasRes.value.data;
  if (!saykas) return null;

  return <SaykaList saykas={saykas} />;
};
