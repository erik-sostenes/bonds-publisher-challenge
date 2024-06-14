"use client";
import { BondsTable } from "@/components/BondsTable";
import { Toaster } from "@/components/ui/toaster";
import { useToast } from "@/components/ui/use-toast";

import { buyBond, getBonds } from "@/requests/request";
import { Bond } from "@/types/types";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useCallback } from "react";

export default function BuyBondsPage() {
  const { toast } = useToast();
  const userId = "580b87da-e389-4290-acbf-f6191467f401";

  const queryClient = useQueryClient();

  // get bonds
  const {
    data: bonds,
    isLoading,
    isError,
    error,
  } = useQuery<Bond[], Error>({
    queryKey: ["bonds", userId],
    queryFn: () => getBonds(userId),
  });

  // bought bonds
  const { mutate, isPending } = useMutation({
    mutationFn: buyBond,
    onSuccess: async () => {
      toast({
        description: "Successful operation 🎉",
      });

      // refresh queries
      queryClient.invalidateQueries({ queryKey: ["bonds", userId] });
    },
    onError: (error) => {
      toast({
        variant: "destructive",
        description: error.message,
      });
    },
  });

  const onSubmit = useCallback(
    (bondId: string, buyerUserId: string) => mutate({ bondId, buyerUserId }),
    [mutate]
  );

  return (
    <>
      <Toaster />
      <BondsTable
        isLoading={isLoading}
        isError={isError}
        error={error}
        bonds={bonds}
        isPending={isPending}
        onSubmitCallback={onSubmit}
      />
    </>
  );
}
