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
