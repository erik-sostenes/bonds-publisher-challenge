import { Banxico, Bond, BondsRequest } from "@/types/types";
import { User } from "@/types/types";

export const getUserBonds = async ({
  userId,
  token,
}: {
  userId: string;
  token: string;
}): Promise<BondsRequest> => {
  const response = await fetch(
    `http://localhost:8080/api/v1/bonds/user?current_owner_id=${userId}&limit=25&page=1`,
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
    }
  );

  if (response.status === 403) {
    throw new ForbiddenError("Your session has expired. Please log in again.");
  }

  const body = await response.json();

  if (!response.ok) {
    throw new Error(body.message || "Failed to obtain user bonds");
  }
  const { data, metadata } = body;
  const bonds = data.map((bond: any) => toCamelCase(bond)) as Bond[];
  const banxico = metadata as Banxico;

  return { bonds, banxico } as BondsRequest;
};

export const getBonds = async ({
  userId,
  token,
}: {
  userId: string;
  token: string;
}): Promise<BondsRequest> => {
  const response = await fetch(
    `http://localhost:8080/api/v1/bonds/all?current_owner_id=${userId}&limit=25&page=1`,
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
    }
  );

  if (response.status === 403) {
    throw new ForbiddenError("Your session has expired. Please log in again.");
  }

  const body = await response.json();

  if (!response.ok) {
    throw new Error(body.message || "Failed to obtain user bonds");
  }

  const { data, metadata } = body;
  const bonds = data.map((bond: any) => toCamelCase(bond)) as Bond[];
  const banxico = metadata as Banxico;

  return { bonds, banxico } as BondsRequest;
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

export const saveBond = async ({
  bond,
  token,
}: {
  bond: Bond;
  token: string;
}) => {
  const response = await fetch(`http://localhost:8080/api/v1/bonds/create`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
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

  if (response.status === 403) {
    throw new ForbiddenError("Your session has expired. Please log in again.");
  }
  if (!response.ok) {
    const body = await response.json();
    throw new Error(body.message || "Failed to register bond");
  }
};

export const buyBond = async ({
  bondId,
  buyerUserId,
  token,
}: {
  bondId: string;
  buyerUserId: string;
  token: string;
}) => {
  const paramBondId = encodeURIComponent(bondId);
  const paramBuyerUserId = encodeURIComponent(buyerUserId);

  const response = await fetch(
    `http://localhost:8080/api/v1/bonds/buy/${paramBondId}/${paramBuyerUserId}`,
    {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
    }
  );

  if (response.status === 403) {
    throw new ForbiddenError("Your session has expired. Please log in again.");
  }

  if (!response.ok) {
    const body = await response.json();
    throw new Error(body.message || "Failed to buy bond");
  }
};

export const userAuthentication = async (user: User) => {
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
    throw new Error(body.message || "Failed to authenticate user");
  }

  return body.token as string;
};

export const saveUser = async (user: User) => {
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

// customize error
export class ForbiddenError extends Error {
  constructor(message: string) {
    super(message);
    this.name = "ForbiddenError";
  }
}
