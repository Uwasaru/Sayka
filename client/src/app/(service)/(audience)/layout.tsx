import { AudienceNavigation } from "./_components/AudienceNavigation";

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <div>
      <AudienceNavigation />
      {children}
    </div>
  );
}
