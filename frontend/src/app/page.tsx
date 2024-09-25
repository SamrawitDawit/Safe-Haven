"use client";

import Image from "next/image";
import SignUpForm from "./signup/page";
import { useSession, signOut } from "next-auth/react";

export default function Home() {
  const { data: session } = useSession();
  console.log(session);
  return (
    <div>
      {" "}
      <p className="text-3xl"> Safe Haven Welcome </p>
      {session ? (
        <div>
          <p>Welcome, {session.user?.name}!</p>
          <p className="mt-2">You are logged in with {session.user?.email}</p>

          {/* Logout button */}
          <button
            onClick={() => signOut({ callbackUrl: "/auth/login" })} // Change callbackUrl to login page
            className="mt-4 px-4 py-2 bg-red-500 text-white rounded"
          >
            Logout
          </button>
        </div>
      ) : (
        <p>Please login/signin to continue </p>
      )}
    </div>
  );
}
