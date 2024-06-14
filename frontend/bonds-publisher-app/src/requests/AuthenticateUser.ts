import { User } from "@/types/types";

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
