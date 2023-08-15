import { ContentTitle } from "@/ui/Text";

import { SaykaEditForm } from "../../_components/SaykaEditForm";

type TProps = {
  params: { sayka_id: string };
};

const Page = ({ params }: TProps) => {
  return (
    <div>
      <ContentTitle text="成果の編集" />
      <div className="px-2 md:px-0">
        <SaykaEditForm saykaId={Number(params.sayka_id)} />
      </div>
    </div>
  );
};

export default Page;
