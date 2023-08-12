import { CreativeNavigation } from "./_components/CreativeNavigation";

export default function Layout(props: { children: React.ReactNode }) {
  return (
    <div>
      <CreativeNavigation />
      <div className="grid grid-cols-9">
        <div className="hidden md:col-span-3 md:block" />
        <main className="col-span-10 md:col-span-3">{props.children}</main>
        <div className="hidden md:col-span-3 md:block" />
      </div>
    </div>
  );
}
