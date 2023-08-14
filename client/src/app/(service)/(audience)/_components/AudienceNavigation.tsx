"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import React from "react";

const NAV_LINKS = [{ name: "タイムライン", href: "/timeline" }];
export const AudienceNavigation = () => {
  const pathname = usePathname();

  return (
    <div className="border-b border-b-stone-900 font-semibold">
      <div className="mx-10 flex space-x-10 font-semibold text-gray-600">
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
