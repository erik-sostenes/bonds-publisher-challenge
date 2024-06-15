"use client";
import { BondsTable } from "@/components/BondsTable";
import { Toaster } from "@/components/ui/toaster";
import { useToast } from "@/components/ui/use-toast";

import { buyBond, ForbiddenError, getBonds } from "@/requests/request";
import { useSessionUserStore } from "@/store/useUserStore";
import { BondsRequest } from "@/types/types";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useCallback } from "react";

import { useRouter } from "next/navigation";

export default function BuyBondsPage() {
  const { toast } = useToast();
  const [userWithSession] = useSessionUserStore((state) => [
    state.userWithSession,
  ]);

  const router = useRouter();

  const queryClient = useQueryClient();

  // get bonds
  const { data, isLoading, isError, error } = useQuery<BondsRequest, Error>({
    queryKey: ["bonds", userWithSession.id, userWithSession.session],
    queryFn: () =>
      getBonds({
        userId: userWithSession.id,
        token: userWithSession.session,
      }),
  });

  // bought bonds
  const { mutate, isPending } = useMutation({
    mutationFn: buyBond,
    onSuccess: async () => {
      toast({
        description: "Successful operation 🎉",
      });

      // refresh queries
      queryClient.invalidateQueries({
        queryKey: ["bonds", userWithSession.id, userWithSession.session],
      });
    },
    onError: (error) => {
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
          description: error.message,
        });
      }
    },
  });

  const onSubmit = useCallback(
    (bondId: string, buyerUserId: string, token: string) =>
      mutate({ bondId, buyerUserId, token }),
    [mutate]
  );

  return (
    <>
      <Toaster />
      <BondsTable
        isLoading={isLoading}
        isError={isError}
        error={error}
        bonds={data?.bonds}
        banxico={data?.banxico}
        isPending={isPending}
        onSubmitCallback={onSubmit}
      />
    </>
  );
}
