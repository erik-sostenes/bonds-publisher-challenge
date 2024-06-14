import { Bond } from "@/types/types";

import { create } from "zustand";
import { persist } from "zustand/middleware";

export interface State {
  // userBonds are the bonds that belong to the user
  userBonds: Bond[];
}

export interface Action {
  addUserBonds: (bonds: Bond[]) => void;
  addNewUserBond: (bond: Bond) => void;
  deleteUserBondByName: (name: string) => void;
}

export const useUserBondsStore = create<State & Action>()(
  persist(
    (set, get) => ({
      userBonds: [],
      addUserBonds: (bonds: Bond[]) => {
        set({
          userBonds: bonds,
        });
      },
      addNewUserBond: (bond: Bond) => {
        const { userBonds } = get();

        const newUserBonds = structuredClone(userBonds);
        const bondIndex = newUserBonds.findIndex(
          (value) => value.id === bond.id
        );

        if (bondIndex === -1) {
          newUserBonds.push(bond);

          set({
            userBonds: newUserBonds,
          });
        }
      },
      deleteUserBondByName: (name: string) => {
        const { userBonds } = get();

        const newUserBonds = userBonds.filter((value) => value.name !== name);

        set({
          userBonds: newUserBonds,
        });
      },
    }),
    { name: "user-bonds" }
  )
);
