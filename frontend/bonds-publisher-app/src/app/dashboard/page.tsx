import { MaxWidthWrapper } from "@/components/MaxWidthWrapper";
import { Navbar } from "@/components/Navbar";
import { SheetCreateBond } from "@/components/SheetCreateBond";
import { BondsTable } from "@/components/BondsTable";

export default function DashboardPage() {
  return (
    <MaxWidthWrapper className="flex items-center flex-col justify-start h-screen gap-[5rem]">
      <header className="w-full flex flex-col gap-[2rem] py-1">
        <Navbar />
        <div className="flex justify-between w-full items-center">
          <h1 className="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl">
            My Bonds
          </h1>
          <SheetCreateBond />
        </div>
      </header>

      <BondsTable />
    </MaxWidthWrapper>
  );
}
