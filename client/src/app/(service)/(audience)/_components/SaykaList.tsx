import { TSayka } from "@/types/Sayka";

import { Sayka } from "./Sayka";

type TProps = {
  saykas: TSayka[];
  token?: string;
  noFix?: boolean;
};

export const SaykaList = ({ saykas, token, noFix }: TProps) => {
  return (
    <div className="mt-5 space-y-5 mx-3 md:mx-0">
      {saykas.map((sayka) => (
        <Sayka key={sayka.id} sayka={sayka} token={token} noFix={noFix} />
      ))}
    </div>
  );
};
