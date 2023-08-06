import { TopHeader, TopMain } from "@/ui/Layout";

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <div className="bg-white">
      <TopHeader />
      <TopMain>{children}</TopMain>
    </div>
  );
}
