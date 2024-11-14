"use client";
import { useEffect } from "react";
import { ulid } from "ulidx";
import { SignInButton, SignedIn, SignedOut, UserButton } from "@clerk/nextjs";

export default function Header() {
  useEffect(() => {
    if (!localStorage.getItem("tempOrgID")) {
      localStorage.setItem("tempOrgID", ulid());
    }
  }, []);

  return (
    <header className="bg-white shadow-sm w-full">
      <div className="mx-auto py-4 px-4 sm:px-6 lg:px-8 flex justify-between items-center">
        <h1 className="text-2xl font-bold text-gray-900">pollis</h1>
        <SignedOut>
          <SignInButton forceRedirectUrl={"/app"} />
        </SignedOut>
        <SignedIn>
          <UserButton />
        </SignedIn>
      </div>
    </header>
  );
}
