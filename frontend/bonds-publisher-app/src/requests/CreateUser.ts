import { User } from "@/types/types";

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
