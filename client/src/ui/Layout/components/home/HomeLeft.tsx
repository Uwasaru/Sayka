export const HomeLeft = ({ children }: { children: React.ReactNode }) => {
  return (
    <main className="flex h-[90vh] w-[60vh] flex-col justify-center p-10">
      {children}
    </main>
  );
};
