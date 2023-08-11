"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import React from "react";

const NAV_LINKS = [
  { name: "ホーム", href: "/" },
  { name: "タイムライン", href: "/timeline" },
];
export const ServiceNavigation = () => {
  const pathname = usePathname();

  return (
    <div className="fle font-semibold border-b-stone-900 border-b">
      <div className="space-x-10 text-gray-600 flex font-semibold mx-10">
        {NAV_LINKS.map((link) => {
          const isActive = pathname === link.href;

          return (
            <Link
              className={
                isActive ? "border-b-4 border-sc" : "hover:text-gray-900"
              }
              href={link.href}
              key={link.name}>
              {link.name}
            </Link>
          );
        })}
      </div>
    </div>
  );
};
