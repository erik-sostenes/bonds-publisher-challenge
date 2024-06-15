"use client";
import { BondsTable } from "@/components/BondsTable";
import { getUserBonds } from "@/requests/request";
import { useSessionUserStore } from "@/store/useUserStore";
import { Bond, BondsRequest } from "@/types/types";
import { useQuery } from "@tanstack/react-query";

export default function DashboardPage() {
  const [userWithSession] = useSessionUserStore((state) => [
    state.userWithSession,
  ]);

  const { data, isLoading, isError, error } = useQuery<BondsRequest, Error>({
    queryKey: ["user-bonds"],
    queryFn: () =>
      getUserBonds({
        userId: userWithSession.id,
        token: userWithSession.session,
      }),
  });

  return (
    <BondsTable
      isLoading={isLoading}
      isError={isError}
      error={error}
      bonds={data?.bonds}
      banxico={data?.banxico}
    />
  );
}
