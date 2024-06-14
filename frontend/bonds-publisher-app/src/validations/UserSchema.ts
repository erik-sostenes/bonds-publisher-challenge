import { z } from "zod";

export type UserFormValues = z.infer<typeof userFormSchema>;

export const userFormSchema = z.object({
  name: z
    .string()
    .min(10, {
      message: "Username must be at least 10 characters",
    })
    .max(50, {
      message: "Username must not exceed 50 characters",
    }),
  password: z
    .string()
    .min(5, {
      message: "Password must be at least 5 characters.",
    })
    .max(20, {
      message: "Password must not exceed 20 characters",
    }),
});
