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

import { useToast } from "@/components/ui/use-toast";
import { Toaster } from "@/components/ui/toaster";

import { Button } from "@/components/ui/button";
import { MaxWidthWrapper } from "@/components/MaxWidthWrapper";
import { RoleEnum, User } from "@/types/types";
import { useMutation } from "@tanstack/react-query";
import { saveUserRequest } from "@/requests/CreateUser";
import { ReloadIcon } from "@radix-ui/react-icons";

import { useCallback } from "react";

type CreateUserFormValues = z.infer<typeof formSchema>;

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
});

export default function SignUp() {
  const { toast } = useToast();

  const form = useForm<CreateUserFormValues>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: "",
      password: "",
    },
  });

  const { mutate, isPending } = useMutation({
    mutationFn: saveUserRequest,
    onSuccess: async () => {
      form.reset({
        name: "",
        password: "",
      });

      toast({
        description: "Registered successfully 🎉",
      });
    },
    onError: (error) => {
      toast({
        variant: "destructive",
        description: error.message,
      });
    },
  });

  const onSubmit = useCallback(
    (values: CreateUserFormValues) => {
      const user: User = {
        id: crypto.randomUUID(),
        role: {
          id: 1,
          type: RoleEnum.USER,
        },
        ...values,
      };
      mutate(user);
    },
    [mutate]
  );

  return (
    <MaxWidthWrapper className="flex items-center flex-col">
      <Toaster />
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
            disabled={isPending}
          >
            {!isPending && "Register"}
            {isPending && <ReloadIcon className="mr-2 h-4 w-4 animate-spin" />}
          </Button>
        </form>
      </Form>
    </MaxWidthWrapper>
  );
}
