import { FC } from "react";

import { TUser } from "@/types/User";

import { Sayka } from "./Sayka";
import { TSayka } from "@/types/Sayka";
import { readUser } from "@/api";
import { redirect } from "next/navigation";

type TProps = {
  loginUser?: TUser;
  sayka: TSayka;
  changeFilterTag?: (tag: string) => void;
};

const SaykaBase = async ({ sayka, changeFilterTag, loginUser }: TProps) => {
  const readUserRes = await readUser(sayka.user_id);
  if (readUserRes.type === "error") {
    redirect("/");
  }
  return (
    <Sayka
      sayka={sayka}
      changeFilterTag={changeFilterTag}
      loginUser={loginUser}
      user={readUserRes.value.data}
    />
  );
};

export default SaykaBase;
