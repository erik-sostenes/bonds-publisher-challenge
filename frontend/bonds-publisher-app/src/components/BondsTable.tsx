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

import { usePathname, useRouter } from "next/navigation";

import { cn } from "@/lib/utils";
import { Banxico, Bond } from "@/types/types";

import { ReloadIcon } from "@radix-ui/react-icons";
import { linkBuyBonds, linkDashboard } from "./Navbar";
import { useSessionUserStore } from "@/store/useUserStore";
import { Button } from "./ui/button";
import { ForbiddenError } from "@/requests/request";
import { useToast } from "./ui/use-toast";
import { useCallback, useMemo, useState } from "react";
import { formatterForCurrency } from "@/lib/helpers";
import { Currency, locations } from "@/lib/currencies";

interface Props {
  isLoading: boolean;
  isError: boolean;
  error: Error | null;
  bonds: Bond[] | undefined;
  banxico: Banxico | undefined;
  onSubmitCallback?: (
    bondId: string,
    buyerUserId: string,
    token: string
  ) => void;
  isPending?: boolean;
}

export function BondsTable({
  isLoading,
  isError,
  bonds,
  banxico,
  isPending,
  error,
  onSubmitCallback,
}: Props) {
  const [userWithSession] = useSessionUserStore((state) => [
    state.userWithSession,
  ]);
  const { toast } = useToast();
  const router = useRouter();

  const pathname = usePathname();
  const isBuyBond = pathname === linkDashboard + linkBuyBonds;
  const [currencyFormat, setCurrencyFormat] = useState<Currency>(Currency.MXN);

  const formatterCurrency = useCallback(formatterForCurrency, []);

  const bondsFormatterCurrency = useMemo(() => {
    const strDolar = banxico?.series[0].datos[0].dato;
    const dolar = strDolar ? parseFloat(strDolar) : 0;

    return bonds?.map((bond) => {
      return {
        ...bond,
        salesPrice: formatterCurrency(currencyFormat).format(
          currencyFormat === Currency.MXN
            ? bond.salesPrice
            : bond.salesPrice / dolar
        ),
      };
    });
  }, [bonds, currencyFormat, formatterCurrency, banxico]);

  // waiting to obtain the bonds
  if (isLoading) return <ReloadIcon className="mr-2 h-4 w-4 animate-spin" />;

  if (isError) {
    if (error instanceof ForbiddenError) {
      // Redirect to login or some other page
      router.push("/");
      setTimeout(() => {
        toast({
          variant: "destructive",
          description: "your session has expired",
        });
      }, 1000);
    } else {
      return <p>error occurred</p>;
    }
  }

  // waiting for the bond to be successfully purchased
  if (isPending) return <ReloadIcon className="mr-2 h-4 w-4 animate-spin" />;

  return (
    <>
      <div className="w-full flex justify-start gap-3">
        <Button
          variant={currencyFormat === Currency.MXN ? "default" : "outline"}
          onClick={() => setCurrencyFormat(Currency.MXN)}
        >
          {Currency.MXN}
        </Button>
        <Button
          variant={currencyFormat === Currency.USD ? "default" : "outline"}
          onClick={() => setCurrencyFormat(Currency.USD)}
        >
          {Currency.USD}
        </Button>
      </div>

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
          {bondsFormatterCurrency?.map((bond) => (
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
                      onSubmitCallback &&
                      onSubmitCallback(
                        bond.id,
                        userWithSession.id,
                        userWithSession.session
                      )
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
