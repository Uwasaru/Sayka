import { InquiryForm } from "./InquiryForm";
import { Newses } from "./Newses";
import { RightTopAd } from "./RightTopAd";

export const RightBox = () => {
  return (
    <div className="space-y-10">
      <RightTopAd />
      <Newses />
      <InquiryForm />
    </div>
  );
};
