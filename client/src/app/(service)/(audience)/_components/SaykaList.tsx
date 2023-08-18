import { TUser } from "@/types/User";

import { TSayka } from "@/types/Sayka";
import { Sayka } from "./Sayka";

type TProps = {
  loginUser?: TUser;
  saykas: TSayka[];
};

export const SaykaList = ({ saykas, loginUser }: TProps) => {
  return (
    <div className="mt-5 space-y-5">
      {saykas.map((sayka) => (
        <Sayka key={sayka.id} sayka={sayka} loginUser={loginUser} />
      ))}
    </div>
  );
};
