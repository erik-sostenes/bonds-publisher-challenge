import {
  Sheet,
  SheetContent,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet";
import { Button } from "./ui/button";

export function SheetCreateBond() {
  return (
    <Sheet>
      <SheetTrigger>
        <Button>Create Bond</Button>
      </SheetTrigger>
      <SheetContent>
        <SheetHeader>
          <SheetTitle>Create Bond</SheetTitle>
        </SheetHeader>
      </SheetContent>
    </Sheet>
  );
}
