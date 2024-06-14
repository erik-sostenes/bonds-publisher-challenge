"use client";
import {
  Sheet,
  SheetContent,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet";

import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";

import { useToast } from "@/components/ui/use-toast";
import { Toaster } from "@/components/ui/toaster";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";

import { zodResolver } from "@hookform/resolvers/zod";
import { ReloadIcon } from "@radix-ui/react-icons";
import { useMutation } from "@tanstack/react-query";
import { useCallback, useRef } from "react";
import { useForm } from "react-hook-form";

import { bondFormSchema, BondFormValues } from "@/validations/BondSchema";

import { Bond } from "@/types/types";
import { saveBond } from "@/requests/request";
import { useUserBondsStore } from "@/store/useBondStore";

export function SheetCreateBond() {
  const userId = "580b87da-e389-4290-acbf-f6191467f401";

  return (
    <Sheet>
      <SheetTrigger>Create Bond</SheetTrigger>
      <SheetContent className="w-full sm:w-[550px]">
        <SheetHeader>
          <SheetTitle>Create Bond</SheetTitle>
        </SheetHeader>
        <CreateBond userId={userId}></CreateBond>
      </SheetContent>
    </Sheet>
  );
}

export function CreateBond({ userId }: { userId: string }) {
  const currentNewUBond = useRef<Bond>();
  const [addNewUserBond] = useUserBondsStore((state) => [state.addNewUserBond]);

  const { toast } = useToast();

  const form = useForm<BondFormValues>({
    resolver: zodResolver(bondFormSchema),
    defaultValues: {
      name: "",
      quantitySale: 0,
      salesPrice: 0,
    },
  });

  const { mutate, isPending } = useMutation({
    mutationFn: saveBond,
    onSuccess: async () => {
      addNewUserBond(currentNewUBond.current || ({} as Bond));
      form.reset({
        name: "",
        quantitySale: 0,
        salesPrice: 0,
      });

      toast({
        description: "Registered successfully 🎉",
      });
    },
    onError: (error) => {
      currentNewUBond.current = {} as Bond;
      toast({
        variant: "destructive",
        description: error.message,
      });
    },
  });

  const onSubmit = useCallback(
    (values: BondFormValues) => {
      const bond = {
        id: crypto.randomUUID(),
        isBought: false,
        creatorUserId: userId,
        currentOwnerId: userId,
        ...values,
      };

      currentNewUBond.current = bond as Bond;

      mutate(bond as Bond);
    },
    [mutate, currentNewUBond, userId]
  );

  return (
    <>
      <Toaster />
      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(onSubmit)}
          className="w-full space-y-6"
        >
          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Bond name</FormLabel>
                <FormControl>
                  <Input placeholder="bond name" {...field} className="h-10" />
                </FormControl>
                <FormDescription>
                  Enter a unique bond name for identification
                </FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="quantitySale"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Quantity Sale</FormLabel>
                <FormControl>
                  <Input
                    type="number"
                    {...field}
                    className="h-10"
                    value={field.value || ""}
                    onChange={(e) =>
                      field.onChange(parseInt(e.target.value, 10))
                    }
                  />
                </FormControl>
                <FormDescription>Establishes a sales quantity</FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="salesPrice"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Sales Price</FormLabel>
                <FormControl>
                  <Input
                    className="h-10"
                    {...field}
                    type="number"
                    value={field.value || ""}
                    onChange={(e) =>
                      field.onChange(parseInt(e.target.value, 10))
                    }
                  />
                </FormControl>
                <FormDescription>Establishes a selling price</FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
          <Button
            variant="default"
            size="lg"
            className="mt-5 h-12 text-base font-bold cursor-pointer w-full"
            disabled={isPending}
          >
            {!isPending && "Created"}
            {isPending && <ReloadIcon className="mr-2 h-4 w-4 animate-spin" />}
          </Button>
        </form>
      </Form>
    </>
  );
}
