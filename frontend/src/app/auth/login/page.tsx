"use client";
import React, { useState } from "react";
import { useForm, SubmitHandler } from "react-hook-form";
import { useRouter } from "next/navigation";
import { signIn } from "next-auth/react";
import Link from "next/link";
import SignInWithGoogle from "../../signup/components/SignInWithGoogle";

interface LoginValues {
  identifier: string; // Email or phone number
  password: string;
}

const LoginEmail = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginValues>();
  const router = useRouter();
  const [errorMessage, setErrorMessage] = useState<string | null>(null);

  const onSubmit: SubmitHandler<LoginValues> = async (data) => {
    const res = await signIn("credentials", {
      redirect: false, // Prevent automatic redirects to handle errors manually
      identifier: data.identifier,
      password: data.password,
    });

    if (res?.error) {
      setErrorMessage("Failed to login. Please check your credentials.");
    } else {
      // If login is successful, redirect to the homepage
      router.push("/");
    }
  };

  return (
    <div className="flex justify-center mt-0">
      <div className="flex flex-col gap-4 w-[480px] mt-50">
        <h2 className="text-3xl text-center font-bold font-sans my-4">
          Log In
        </h2>

        <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-1">
          {/* Identifier Field (Email or Phone) */}
          <label htmlFor="identifier" className="input-label">
            Email or Phone Number:
          </label>
          <input
            type="text"
            id="identifier"
            {...register("identifier", {
              required: "Email or phone number is required",
            })}
            placeholder="Enter your email or phone number"
            className="border rounded p-2"
          />
          {errors.identifier && (
            <p className="text-red-500">{errors.identifier.message}</p>
          )}

          {/* Password Field */}
          <label htmlFor="password" className="input-label">
            Password:
          </label>
          <input
            type="password"
            id="password"
            {...register("password", { required: "Password is required" })}
            placeholder="Enter your password"
            className="border rounded p-2"
          />
          {errors.password && (
            <p className="text-red-500">{errors.password.message}</p>
          )}

          {/* Display Error Message */}
          {errorMessage && <p className="text-red-500">{errorMessage}</p>}

          {/* Submit Button */}
          <input
            type="submit"
            value="Log In"
            className="bg-[#4640DE] px-6 py-3 rounded-[80px] text-white font-epilogue font-bold text-[12px] mt-4"
          />
        </form>

        {/* Sign In with Google */}
        <SignInWithGoogle />

        {/* Sign Up Link */}
        <div className="w-fit h-fit flex gap-3">
          <p className="text-[16px] font-epilogue font-normal leading-[25.6px] text-[#202430] w-[203px] h-[26px] opacity-[0.7]">
            Don&apos;t have an account?
          </p>
          <Link
            href="/signup"
            className="text-[16px] font-inter font-semibold leading-[24px] text-[#4640DE] h-[24px]"
          >
            Sign up
          </Link>
        </div>
        <div className="w-fit h-fit flex gap-3">
          <p className="text-[16px] font-epilogue font-normal leading-[25.6px] text-[#202430] w-[203px] h-[26px] opacity-[0.7]">
            Don't remember your password?
          </p>
          <Link
            href="/auth/forgot-password"
            className="text-[16px] font-inter font-semibold leading-[24px] text-[#4640DE] h-[24px]"
          >
            Forgot Password
          </Link>
        </div>
      </div>
    </div>
  );
};

export default LoginEmail;
