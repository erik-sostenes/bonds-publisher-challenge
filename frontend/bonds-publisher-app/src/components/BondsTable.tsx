"use client";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { useToast } from "@/components/ui/use-toast";

import { usePathname } from "next/navigation";

import { cn } from "@/lib/utils";
import { Bond } from "@/types/types";

import { ReloadIcon } from "@radix-ui/react-icons";
import { linkBuyBonds, linkDashboard } from "./Navbar";
import { Button } from "./ui/button";

interface Props {
  isLoading: boolean;
  isError: boolean;
  error: Error | null;
  bonds: Bond[] | undefined;
  onSubmitCallback?: (bondId: string, buyerUserId: string) => void;
  isPending?: boolean;
}

export function BondsTable({
  isLoading,
  isError,
  error,
  bonds,
  onSubmitCallback,
  isPending,
}: Props) {
  const buyerUserId = "580b87da-e389-4290-acbf-f6191467f401";

  const pathname = usePathname();
  const isBuyBond = pathname === linkDashboard + linkBuyBonds;

  // waiting to obtain the bonds
  if (isLoading) return <ReloadIcon className="mr-2 h-4 w-4 animate-spin" />;
  if (isError) return <span>Error: {error?.message}</span>;

  // waiting for the bond to be successfully purchased
  if (isPending) return <ReloadIcon className="mr-2 h-4 w-4 animate-spin" />;

  return (
    <>
      <Table>
        <TableCaption>A list of your bonds</TableCaption>
        <TableHeader>
          <TableRow>
            <TableHead>Id</TableHead>
            <TableHead>name</TableHead>
            <TableHead>Quantity Sale</TableHead>
            <TableHead className="text-right">Sales Price</TableHead>
            <TableHead>State</TableHead>
            {isBuyBond ? <TableHead></TableHead> : <></>}
          </TableRow>
        </TableHeader>
        <TableBody>
          {bonds?.map((bond) => (
            <TableRow key={bond.id}>
              <TableCell className="font-medium">{bond.id}</TableCell>
              <TableCell>{bond.name}</TableCell>
              <TableCell>{bond.quantitySale}</TableCell>
              <TableCell className="text-right">{bond.salesPrice}</TableCell>
              <TableCell>
                <span
                  className={cn(
                    " p-[2px] rounded-md",
                    bond.isBought
                      ? "bg-amber-200 text-yellow-700"
                      : "bg-emerald-300 text-teal-700"
                  )}
                >
                  {bond.isBought ? "bought" : " available"}
                </span>
              </TableCell>
              {isBuyBond ? (
                <TableCell>
                  <Button
                    variant="secondary"
                    size="icon"
                    onClick={() =>
                      onSubmitCallback && onSubmitCallback(bond.id, buyerUserId)
                    }
                  >
                    {!isPending && "Buy"}
                    {isPending && (
                      <ReloadIcon className="mr-2 h-4 w-4 animate-spin" />
                    )}
                  </Button>
                </TableCell>
              ) : (
                <></>
              )}
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </>
  );
}
