"use client";
import { Input } from "@/components/ui/input";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";

import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Button, buttonVariants } from "@/components/ui/button";
import { cn } from "@/lib/utils";
import { MaxWidthWrapper } from "@/components/MaxWidthWrapper";

enum RoleEnum {
  USER = "USER",
}

interface Role {
  id: number;
  role: RoleEnum;
}

interface User {
  id: number;
  name: string;
  password: string;
  role: Role;
}

const roleSchema = z.object({
  id: z.number(),
  type: z.nativeEnum(RoleEnum),
});

type FormValues = z.infer<typeof formSchema>;

const formSchema = z.object({
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
  role: roleSchema,
});

export default function Page() {
  const form = useForm<FormValues>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: "",
      password: "",
      role: {
        id: 1,
        type: RoleEnum.USER,
      },
    },
  });

  function onSubmit(values: FormValues) {
    const user = { ...values, id: crypto.randomUUID() };
    console.log(user);
  }

  return (
    <MaxWidthWrapper className="flex items-center flex-col">
      <h1 className="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl">
        Sign UP
      </h1>
      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(onSubmit)}
          className="w-2/3 space-y-6"
        >
          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem>
                <FormLabel>User name</FormLabel>
                <FormControl>
                  <Input placeholder="user name" {...field} className="h-10" />
                </FormControl>
                <FormDescription>
                  Enter a unique username for identification
                </FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="password"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Password</FormLabel>
                <FormControl>
                  <Input type="password" {...field} className="h-10" />
                </FormControl>
                <FormDescription>Create a secure password</FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
          <Button
            variant="default"
            size="lg"
            className=" mt-5 h-12 text-base font-bold cursor-pointer w-full"
          >
            Register
          </Button>
        </form>
      </Form>
    </MaxWidthWrapper>
  );
}
