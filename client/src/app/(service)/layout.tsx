import { ServiceHeader } from "@/ui/Layout";

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <div className="mb-16 bg-white">
      <ServiceHeader />
      <main>{children}</main>
    </div>
  );
}
