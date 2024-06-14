"use client";
import { MaxWidthWrapper } from "@/components/MaxWidthWrapper";
import { linkDashboard, Navbar } from "@/components/Navbar";
import { SheetCreateBond } from "@/components/SheetCreateBond";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

import { usePathname } from "next/navigation";

const queryClient = new QueryClient();
export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const pathname = usePathname();

  const isDashboard = pathname === linkDashboard;

  return (
    <QueryClientProvider client={queryClient}>
      <main className="h-screen flex justify-center items-center flex-col">
        <MaxWidthWrapper className="flex items-center flex-col justify-start h-screen gap-[5rem]">
          <header className="w-full flex flex-col gap-[2rem] py-1">
            <Navbar />
            <div className="flex justify-between w-full items-center">
              <h1 className="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl">
                {isDashboard ? "My Bonds" : "Buy Bonds"}
              </h1>
              {isDashboard ? <SheetCreateBond /> : <></>}
            </div>
          </header>
          {children}
        </MaxWidthWrapper>
      </main>
    </QueryClientProvider>
  );
}
