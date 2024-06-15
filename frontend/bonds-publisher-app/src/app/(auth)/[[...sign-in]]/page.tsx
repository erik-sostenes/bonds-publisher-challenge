"use client";
import { Input } from "@/components/ui/input";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";

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
import { Role, User, UserWithSession } from "@/types/types";
import { useMutation } from "@tanstack/react-query";
import { ReloadIcon } from "@radix-ui/react-icons";

import { useCallback } from "react";
import Link from "next/link";
import { userFormSchema, UserFormValues } from "@/validations/UserSchema";
import { userAuthentication } from "@/requests/request";

import { jwtDecode } from "jwt-decode";
import { useSessionUserStore } from "@/store/useUserStore";

import { useRouter } from "next/navigation";
import { linkDashboard } from "@/components/Navbar";

interface DecodedToken {
  Permissions: number;
  Role: Role;
  UserID: string;
  UserName: string;
  exp: number;
  iss: string;
}

export default function SignIn() {
  const router = useRouter();
  const { toast } = useToast();

  const [saveSessionUser] = useSessionUserStore((state) => [
    state.saveSessionUser,
  ]);

  const form = useForm<UserFormValues>({
    resolver: zodResolver(userFormSchema),
    defaultValues: {
      name: "",
      password: "",
    },
  });

  const { mutate, isPending } = useMutation({
    mutationFn: userAuthentication,
    onSuccess: async (token) => {
      form.reset({
        name: "",
        password: "",
      });

      const { UserID, UserName, Role, Permissions, exp, iss } = jwtDecode(
        token
      ) as DecodedToken;

      saveSessionUser({
        id: UserID,
        name: UserName,
        role: {
          id: Role.id,
          type: Role.type,
        },
        permissions: Permissions,
        exp: exp,
        iss: iss,
        session: token,
      } as UserWithSession);

      router.push(linkDashboard);
    },
    onError: (error) => {
      toast({
        variant: "destructive",
        description: error.message,
      });
    },
  });

  const onSubmit = useCallback(
    (values: UserFormValues) => {
      const user = {
        ...values,
      };
      mutate(user as User);
    },
    [mutate]
  );

  return (
    <MaxWidthWrapper className="flex items-center flex-col">
      <Toaster />
      <h1 className="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl">
        Sign IN
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
            className="mt-5 h-12 text-base font-bold cursor-pointer w-full"
            disabled={isPending}
          >
            {!isPending && "Start"}
            {isPending && <ReloadIcon className="mr-2 h-4 w-4 animate-spin" />}
          </Button>
        </form>
      </Form>
      <Link href="/sign-up">Register</Link>
    </MaxWidthWrapper>
  );
}
