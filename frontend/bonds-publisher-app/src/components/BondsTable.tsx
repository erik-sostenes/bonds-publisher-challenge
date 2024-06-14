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

import { cn } from "@/lib/utils";
import { getUserBonds } from "@/requests/request";
import { useUserBondsStore } from "@/store/useBondStore";
import { Bond } from "@/types/types";
import { ReloadIcon } from "@radix-ui/react-icons";
import { useQuery } from "@tanstack/react-query";
import { useEffect } from "react";

export function BondsTable() {
  const [addToUserBonds, userBonds] = useUserBondsStore((state) => [
    state.addUserBonds,
    state.userBonds,
  ]);

  const userId = "580b87da-e389-4290-acbf-f6191467f401";

  const {
    data: bonds,
    isLoading,
    isError,
    error,
  } = useQuery<Bond[], Error>({
    queryKey: ["user-bonds", userId],
    queryFn: () => getUserBonds(userId),
  });

  useEffect(() => {
    addToUserBonds(bonds || []);
  }, [addToUserBonds, bonds]);

  if (isLoading) return <ReloadIcon className="mr-2 h-4 w-4 animate-spin" />;
  if (isError) return <span>Error: {error.message}</span>;

  return (
    <>
      <Table>
        <TableCaption>A list of your bonds</TableCaption>
        <TableHeader>
          <TableRow>
            <TableHead className="w-[100px]">Id</TableHead>
            <TableHead>name</TableHead>
            <TableHead>Quantity Sale</TableHead>
            <TableHead className="text-right">Sales Price</TableHead>
            <TableHead>State</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {userBonds?.map((bond) => (
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
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </>
  );
}
