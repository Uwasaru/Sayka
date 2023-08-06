import { TopHeader, TopMain } from "@/ui/Layout";
import { ServiceHeader } from "@/ui/Layout/components/ServiceHeader";

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <div className="bg-white">
      <ServiceHeader />
      {children}
    </div>
  );
}
