export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <main className="h-screen flex justify-center items-center flex-col">
      {children}
    </main>
  );
}
