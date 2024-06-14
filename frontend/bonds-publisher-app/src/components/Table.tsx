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
const bonds = [
  {
    id: "1b4e28ba-2fa1-11d2-883f-0016d3cca427",
    name: "Government Bond",
    quantitySale: 500,
    salesPrice: 25000.0,
    isBought: false,
    creatorUserId: "92b92977-7c7d-4e53-8a38-5250f7f1d83b",
    currentOwnerId: "507f1f77bcf86cd799439011",
  },
  {
    id: "1b4e28ba-2fa1-11d2-883f-0016d3cca428",
    name: "Corporate Bond",
    quantitySale: 1000,
    salesPrice: 50000.0,
    isBought: true,
    creatorUserId: "c0fdf9be-27ee-4bb5-b80a-70063d8afc24",
    currentOwnerId: "507f1f77bcf86cd799439012",
  },
  {
    id: "1b4e28ba-2fa1-11d2-883f-0016d3cca429",
    name: "Municipal Bond",
    quantitySale: 750,
    salesPrice: 37500.0,
    isBought: false,
    creatorUserId: "7e8dcb98-c629-42ad-a41f-2b795cbd7983",
    currentOwnerId: "507f1f77bcf86cd799439013",
  },
  {
    id: "1b4e28ba-2fa1-11d2-883f-0016d3cca430",
    name: "Treasury Bond",
    quantitySale: 200,
    salesPrice: 10000.0,
    isBought: true,
    creatorUserId: "9c8f7c59-41a4-4d5e-a6e2-7b88e01a824d",
    currentOwnerId: "507f1f77bcf86cd799439014",
  },
  {
    id: "1b4e28ba-2fa1-11d2-883f-0016d3cca431",
    name: "Savings Bond",
    quantitySale: 300,
    salesPrice: 15000.0,
    isBought: false,
    creatorUserId: "d1b8c1c1-711f-4d49-8b0e-66c4efb9ecb5",
    currentOwnerId: "507f1f77bcf86cd799439015",
  },
  {
    id: "1b4e28ba-2fa1-11d2-883f-0016d3cca432",
    name: "Investment Bond",
    quantitySale: 400,
    salesPrice: 20000.0,
    isBought: true,
    creatorUserId: "b6c4a2e2-830c-4030-92e8-5cbd1c8f1fcb",
    currentOwnerId: "507f1f77bcf86cd799439016",
  },
  {
    id: "1b4e28ba-2fa1-11d2-883f-0016d3cca433",
    name: "High Yield Bond",
    quantitySale: 600,
    salesPrice: 30000.0,
    isBought: false,
    creatorUserId: "3d06c5d7-7915-48b8-badf-603e4b44c6bc",
    currentOwnerId: "507f1f77bcf86cd799439017",
  },
  {
    id: "1b4e28ba-2fa1-11d2-883f-0016d3cca434",
    name: "Zero Coupon Bond",
    quantitySale: 900,
    salesPrice: 45000.0,
    isBought: true,
    creatorUserId: "2f77d9a8-558f-42e6-843b-0ab8bcdaf678",
    currentOwnerId: "507f1f77bcf86cd799439018",
  },
  {
    id: "1b4e28ba-2fa1-11d2-883f-0016d3cca435",
    name: "Convertible Bond",
    quantitySale: 800,
    salesPrice: 40000.0,
    isBought: false,
    creatorUserId: "58a0c0e9-8f89-4d23-8c2a-b6b9f44a1326",
    currentOwnerId: "507f1f77bcf86cd799439019",
  },
  {
    id: "1b4e28ba-2fa1-11d2-883f-0016d3cca436",
    name: "War Bond",
    quantitySale: 100,
    salesPrice: 5000.0,
    isBought: true,
    creatorUserId: "d90b5b1c-6d5b-40b7-875e-53b847c5a1d1",
    currentOwnerId: "507f1f77bcf86cd799439020",
  },
];

export function BondsTable() {
  return (
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
        {bonds.map((bond) => (
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
  );
}
