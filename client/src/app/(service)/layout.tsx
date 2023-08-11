import { ServiceHeader } from "@/ui/Layout";
import { ServiceNavigation } from "./_components/ServiceNavigation";

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <div className="bg-white">
      <ServiceHeader />
      <ServiceNavigation />
      <main>{children}</main>
    </div>
  );
}
