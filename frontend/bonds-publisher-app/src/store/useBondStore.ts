import { Bond } from "@/types/types";

import { create } from "zustand";
import { persist } from "zustand/middleware";

export interface State {
  // userBonds are the bonds that belong to the user
  userBonds: Bond[];
}

export interface Action {
  addToUserBonds: (bonds: Bond[]) => void;
}

export const useUserBondsStore = create<State & Action>()(
  persist(
    (set, get) => ({
      userBonds: [],
      addToUserBonds: (bonds: Bond[]) => {
        set({
          userBonds: bonds,
        });
      },
    }),
    { name: "user-bonds" }
  )
);
