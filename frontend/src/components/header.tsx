"use client";
import Image from "next/image";
import logo from "../../public/logo.png";
import { Button } from "./ui/button";
import { useRouter } from "next/navigation";

export default function Header() {
  const router = useRouter();
  return (
    <div className="bg-customprimary flex justify-between p-4">
      <div className="flex gap-3">
        <Image src={logo} alt="safe-haven logo" />
        <p className="text-white">Safe Haven</p>
      </div>
      <nav>
        <ul className="flex text-white gap-4 mx-16">
          <li>Home</li>
          <li>About Us</li>
          <li>Education</li>
          <li>Services</li>
        </ul>
      </nav>
      <div className="gap-2">
        <Button
          className=" rounded-full bg-transparent border border-black"
          onClick={() => router.push("/auth/login")}
        >
          Login
        </Button>

        <Button
          className="border-black rounded-full bg-white text-black hover:bg-black hover:text-white"
          onClick={() => router.push("/auth/signup")}
        >
          Sign Up
        </Button>
      </div>
    </div>
  );
}
