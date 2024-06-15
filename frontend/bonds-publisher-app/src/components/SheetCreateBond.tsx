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
import { Button, buttonVariants } from "@/components/ui/button";
import { Input } from "@/components/ui/input";

import { zodResolver } from "@hookform/resolvers/zod";
import { ReloadIcon } from "@radix-ui/react-icons";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useCallback, useRef } from "react";
import { useForm } from "react-hook-form";
import { useRouter } from "next/navigation";

import { bondFormSchema, BondFormValues } from "@/validations/BondSchema";

import { Bond } from "@/types/types";
import { ForbiddenError, saveBond } from "@/requests/request";
import { useSessionUserStore } from "@/store/useUserStore";

export function SheetCreateBond() {
  const [userWithSession] = useSessionUserStore((state) => [
    state.userWithSession,
  ]);

  return (
    <Sheet>
      <SheetTrigger className={buttonVariants({ variant: "secondary" })}>
        Create Bond
      </SheetTrigger>
      <SheetContent className="w-full sm:w-[550px]">
        <SheetHeader>
          <SheetTitle>Create Bond</SheetTitle>
        </SheetHeader>
        <CreateBond
          userId={userWithSession.id}
          token={userWithSession.session}
        ></CreateBond>
      </SheetContent>
    </Sheet>
  );
}

export function CreateBond({
  userId,
  token,
}: {
  userId: string;
  token: string;
}) {
  const currentNewUBond = useRef<Bond>();
  const queryClient = useQueryClient();
  const router = useRouter();
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
      form.reset({
        name: "",
        quantitySale: 0,
        salesPrice: 0,
      });

      toast({
        description: "Registered successfully 🎉",
      });
      // refresh queries
      queryClient.invalidateQueries({
        queryKey: ["user-bonds"],
      });
    },
    onError: (error) => {
      currentNewUBond.current = {} as Bond;

      if (error instanceof ForbiddenError) {
        // Redirect to login or some other page
        setTimeout(() => {
          router.push("/");
        }, 4000);
        toast({
          variant: "destructive",
          description: "your session has expired",
        });
      } else {
        toast({
          variant: "destructive",
          description: error?.message,
        });
      }
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

      mutate({ bond: bond as Bond, token: token });
    },
    [mutate, currentNewUBond, userId, token]
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
