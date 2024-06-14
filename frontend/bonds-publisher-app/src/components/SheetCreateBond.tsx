import {
  Sheet,
  SheetContent,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet";

export function SheetCreateBond() {
  return (
    <Sheet>
      <SheetTrigger>Create Bond</SheetTrigger>
      <SheetContent>
        <SheetHeader>
          <SheetTitle>Create Bond</SheetTitle>
        </SheetHeader>
      </SheetContent>
    </Sheet>
  );
}
