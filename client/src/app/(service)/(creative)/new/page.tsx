import { ContentTitle } from "@/ui/Text";

import { SaykaForm } from "../_components/SaykaForm";

const Page = () => {
  return (
    <div>
      <ContentTitle text="成果の投稿" />
      <div className="md:px-0 px-2">
        <SaykaForm />
      </div>
    </div>
  );
};

export default Page;
