import { Bond } from "@/types/types";
import { User } from "@/types/types";

export const getUserBonds = async (userId: string): Promise<Bond[]> => {
  const response = await fetch(
    `http://localhost:8080/api/v1/bonds/user?current_owner_id=${userId}&limit=25&page=1`,
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    }
  );

  const body = await response.json();

  if (!response.ok) {
    throw new Error(body.message || "Failed to obtain user bonds");
  }

  return body.map((bond: any) => toCamelCase(bond));
};

const toCamelCase = (bond: any): Bond => {
  return {
    id: bond.id,
    name: bond.name,
    quantitySale: bond.quantity_sale,
    salesPrice: bond.sales_price,
    isBought: bond.is_bought,
    creatorUserId: bond.creator_user_id,
    currentOwnerId: bond.current_owner_id,
  };
};

export const saveBond = async (bond: Bond) => {
  const response = await fetch(`http://localhost:8080/api/v1/bonds/create`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      id: bond.id,
      name: bond.name,
      quantity_sale: bond.quantitySale,
      sales_price: bond.salesPrice,
      creator_user_id: bond.creatorUserId,
      current_owner_id: bond.currentOwnerId,
    }),
  });

  if (!response.ok) {
    const body = await response.json();
    throw new Error(body.message || "Failed to register bond");
  }
};

export const userAuthenticationRequest = async (user: User) => {
  const username = encodeURIComponent(user.name);
  const userPassword = encodeURIComponent(user.password);

  const response = await fetch(
    `http://localhost:8080/api/v1/login?username=${username}&user_password=${userPassword}`,
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    }
  );

  const body = await response.json();

  if (!response.ok) {
    throw new Error(body.message || "Failed to register user");
  }

  return body as string;
};

export const saveUserRequest = async (user: User) => {
  const response = await fetch(`http://localhost:8080/api/v1/register`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      id: user.id,
      name: user.name,
      password: user.password,
      role: {
        id: user.role.id,
        type: user.role.type,
      },
    }),
  });

  if (!response.ok) {
    const body = await response.json();
    throw new Error(body.message || "Failed to register user");
  }
};
