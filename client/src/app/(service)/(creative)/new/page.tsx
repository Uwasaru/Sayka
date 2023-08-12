import { ContentTitle } from "@/ui/Text";

import { SaykaForm } from "../_components/SaykaForm";

const Page = () => {
  return (
    <div>
      <ContentTitle text="成果の投稿" />
      <div>
        <SaykaForm />
      </div>
    </div>
  );
};

export default Page;
