export enum RoleEnum {
  USER = "USER",
}

export interface Role {
  id: number;
  type: RoleEnum;
}

export interface User {
  id: string;
  name: string;
  password: string;
  role: Role;
}

export interface Bond {
  id: string;
  name: string;
  quantitySale: number;
  salesPrice: number;
  isBought: boolean;
  creatorUserId: string;
  currentOwnerId: string;
}
