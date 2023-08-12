import { HomeNavigation } from "./_components/HomeNavigation";

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <div>
      <HomeNavigation />
      {children}
    </div>
  );
}
