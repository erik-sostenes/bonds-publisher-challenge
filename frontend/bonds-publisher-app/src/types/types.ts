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

export interface UserWithSession extends User {
  session: string;
  permissions: number;
  exp: number;
  iss: string;
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

export interface Banxico {
  series: Series[];
}

interface Series {
  idSerie: string;
  titulo: string;
  datos: Datos[];
}

interface Datos {
  fecha: string;
  dato: string;
}

export interface BondsRequest {
  bonds: Bond[];
  banxico: Banxico;
}
