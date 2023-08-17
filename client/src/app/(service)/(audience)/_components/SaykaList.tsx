import { TUser } from "@/types/User";

import { TSayka } from "@/types/Sayka";
import SaykaBase from "./SaykaBase";

type TProps = {
  loginUser?: TUser;
  saykas?: TSayka[];
};

export const SaykaList = ({ saykas, loginUser }: TProps) => {
  if (!saykas) {
    return <div>loading...</div>;
  }
  return (
    <div className="mt-5 space-y-5">
      {saykas.map((sayka) => (
        // {/* @ts-expect-error Server Component */}
        <SaykaBase key={sayka.id} sayka={sayka} loginUser={loginUser} />
      ))}
    </div>
  );
};
