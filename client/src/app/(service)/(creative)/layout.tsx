import { RightBox } from "../(audience)/(home)/_components/RightBox";

export default function Layout(props: { children: React.ReactNode }) {
  return (
    <main className="grid grid-cols-10 p-5">
      <div className="hidden md:col-span-2 md:block" />
      <div className="col-span-10 md:col-span-6">{props.children}</div>
      <div className="hidden md:col-span-2 md:block">
        <RightBox />
      </div>
    </main>
  );
}
