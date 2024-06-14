import { z } from "zod";

export type BondFormValues = z.infer<typeof bondFormSchema>;

export const bondFormSchema = z.object({
  name: z
    .string()
    .min(3, {
      message: "Bond name must be at least 3 characters long",
    })
    .max(40, {
      message: "Bond name must not exceed 40 characters",
    }),
  quantitySale: z
    .number()
    .int()
    .min(1, {
      message:
        "Quantity for sale must be an integer greater than or equal to 1",
    })
    .max(10000, {
      message: "Quantity for sale must not exceed 10,000",
    }),
  salesPrice: z
    .number()
    .min(0, {
      message: "Sales price must be greater than or equal to 0",
    })
    .max(100000000, {
      message: "Sales price must not exceed 100,000,000",
    }),
});
