"use client";
import { usePathname } from "next/navigation";
import { buttonVariants } from "./ui/button";
import Link from "next/link";
import { cn } from "@/lib/utils";

const items = [
  { label: "Dashboard", link: "/dashboard" },
  { label: "Buy", link: "/buy-bonds" },
];

export function Navbar() {
  return (
    <>
      <div className="border-separate border-b bg-background">
        <nav className="container flex items-center justify-between px-8">
          <div className="flex h-full">
            {items.map((item) => (
              <NavbarItem
                key={item.label}
                link={item.link}
                label={item.label}
              />
            ))}
          </div>
        </nav>
      </div>
    </>
  );
}

function NavbarItem({ link, label }: { link: string; label: string }) {
  const pathname = usePathname();
  const isActive = pathname === link;

  return (
    <div className="relative flex items-center">
      <Link
        href={link}
        className={cn(
          buttonVariants({ variant: "ghost" }),
          "w-full justify-start text-lg text-muted-foreground hover:text-foreground",
          isActive && "text-foreground"
        )}
      >
        {label}
      </Link>
    </div>
  );
}
