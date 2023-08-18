import { TSayka } from "@/types/Sayka";
import { TUser } from "@/types/User";

import { Sayka } from "./Sayka";

type TProps = {
  saykas: TSayka[];
  token?: string;
};

export const SaykaList = ({ saykas, token }: TProps) => {
  return (
    <div className="mt-5 space-y-5">
      {saykas.map((sayka) => (
        <Sayka key={sayka.id} sayka={sayka} token={token} />
      ))}
    </div>
  );
};
