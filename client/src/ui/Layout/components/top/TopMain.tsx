export const TopMain = ({ children }: { children: React.ReactNode }) => {
  return (
    <main>
      <div className="flex h-[90vh] flex-col items-center justify-center space-y-10 bg-zinc-900">
        {children}
      </div>
    </main>
  );
};
