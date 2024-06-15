import { UserWithSession } from "@/types/types";
import { create } from "zustand";
import { persist } from "zustand/middleware";

export interface State {
  userWithSession: UserWithSession;
}

export interface Action {
  saveSessionUser: (userWithSession: UserWithSession) => void;
}

export const useSessionUserStore = create<State & Action>()(
  persist(
    (set, get) => ({
      userWithSession: {} as UserWithSession,
      saveSessionUser: (userWithSession: UserWithSession) => {
        set({
          userWithSession: userWithSession,
        });
      },
    }),
    { name: "session-user" }
  )
);
