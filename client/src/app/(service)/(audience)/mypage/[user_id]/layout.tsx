export default function Layout(props: { children: React.ReactNode }) {
  return (
    <div className="grid grid-cols-9 md:p-5 p-1">
      <div className="hidden md:col-span-1 md:block" />
      <main className="col-span-10 md:col-span-7">{props.children}</main>
      <div className="hidden md:col-span-1 md:block" />
    </div>
  );
}
