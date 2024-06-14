"use client";
import { BondsTable } from "@/components/BondsTable";
import { getUserBonds } from "@/requests/request";
import { useUserBondsStore } from "@/store/useBondStore";
import { Bond } from "@/types/types";
import { useQuery } from "@tanstack/react-query";
import { useEffect } from "react";

export default function DashboardPage() {
  const userId = "580b87da-e389-4290-acbf-f6191467f401";

  const [addToUserBonds, userBonds] = useUserBondsStore((state) => [
    state.addUserBonds,
    state.userBonds,
  ]);

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

  return (
    <BondsTable
      isLoading={isLoading}
      isError={isError}
      error={error}
      bonds={userBonds}
    />
  );
}
