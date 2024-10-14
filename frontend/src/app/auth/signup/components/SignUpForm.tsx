"use client";

import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { z } from "zod";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import {
  Form,
  FormField,
  FormItem,
  FormLabel,
  FormControl,
  FormMessage,
} from "@/components/ui/form";
import Link from "next/link";
import { useState } from "react";
import { registerUser, RegisterUserData } from "@/api/auth"; // Import API functions
import { useRouter } from "next/navigation";
import SignInWithGoogle from "./SignInWithGoogle";
import { signIn } from "next-auth/react";

// Zod Schema
const formSchema = z
  .object({
    fullName: z.string().min(2, {
      message: "Full name must be at least 2 characters.",
    }),
    email: z.string().email({
      message: "Invalid email address.",
    }),
    password: z.string().min(6, {
      message: "Password must be at least 6 characters.",
    }),
    confirmPassword: z.string().min(6, {
      message: "Confirm password must be at least 6 characters.",
    }),
    language: z.string().nonempty({
      message: "Preferred language is required.",
    }),
    category: z.enum(["General", "Victim"], {
      required_error: "Category is required.",
    }),
    bio: z.string().optional(), // Optional bio field
  })
  .refine((data) => data.password === data.confirmPassword, {
    path: ["confirmPassword"],
    message: "Passwords do not match.",
  });

type FormData = z.infer<typeof formSchema>;

export default function SignUpForm() {
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const [successMessage, setSuccessMessage] = useState<string | null>(null);

  const form = useForm<FormData>({
    resolver: zodResolver(formSchema),
  });
  const router = useRouter();

  // Handle form submission for email registration
  const onSubmit = async (data: FormData) => {
    try {
      const registerData: RegisterUserData = {
        fullName: data.fullName,
        email: data.email, // Email registration only
        password: data.password,
        userType: "normal", // Always "normal"
        category: data.category,
        language: data.language,
        bio: data.bio, // Optional bio
      };

      // Call the registration API for email-based registration
      const response = await registerUser(registerData);
      setSuccessMessage("Registration successful!");
      setErrorMessage(null);

      await signIn("credentials", {
        redirect: false,
        identifier: data.email, // Sign in using email
        password: data.password,
      });

      router.push("/");
    } catch (error: any) {
      const message = error.message || "Failed to register. Please try again.";
      setErrorMessage(message);
      setSuccessMessage(null);
    }
  };

  return (
    <div className="max-w-lg mx-auto p-4 sm:p-6">
      <h1 className="text-2xl font-bold text-center mb-4">Register</h1>
      <SignInWithGoogle />

      <div className="flex items-center justify-center my-4">
        <hr className="w-full border-t border-gray-300" />
        <span className="px-2 text-gray-500">Or</span>
        <hr className="w-full border-t border-gray-300" />
      </div>

      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
          <FormField
            control={form.control}
            name="fullName"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Full Name</FormLabel>
                <FormControl>
                  <Input placeholder="John Doe" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="email"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Email Address</FormLabel>
                <FormControl>
                  <Input
                    type="email"
                    placeholder="example@example.com"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <Link href="/signup-with-phone" className="text-slate-500 text-sm">
            Use Phone number instead
          </Link>

          <FormField
            control={form.control}
            name="password"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Password</FormLabel>
                <FormControl>
                  <Input type="password" placeholder="********" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="confirmPassword"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Confirm Password</FormLabel>
                <FormControl>
                  <Input type="password" placeholder="********" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="language"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Preferred Language</FormLabel>
                <FormControl>
                  <Select onValueChange={field.onChange}>
                    <SelectTrigger className="border border-gray-300 p-2 rounded-md hover:border-gray-400 transition duration-150">
                      <SelectValue placeholder="Select Preferred Language" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="English">English</SelectItem>
                      <SelectItem value="Amharic">Amharic</SelectItem>
                    </SelectContent>
                  </Select>
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="category"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Category</FormLabel>
                <FormControl>
                  <Select onValueChange={field.onChange}>
                    <SelectTrigger className="border border-gray-300 p-2 rounded-md hover:border-gray-400 transition duration-150">
                      <SelectValue placeholder="Select Category" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="General">General</SelectItem>
                      <SelectItem value="Victim">Victim</SelectItem>
                    </SelectContent>
                  </Select>
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="bio"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Bio (Optional)</FormLabel>
                <FormControl>
                  <Input placeholder="A little about yourself..." {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <button
            type="submit"
            className={`bg-customprimary text-white p-3 rounded flex items-center justify-center space-x-2 ${
              form.formState.isSubmitting ? "opacity-50 cursor-not-allowed" : ""
            }`}
            disabled={form.formState.isSubmitting}
          >
            {form.formState.isSubmitting ? (
              <svg
                className="animate-spin h-5 w-5 text-white"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
              >
                <circle
                  className="opacity-25"
                  cx="12"
                  cy="12"
                  r="10"
                  stroke="currentColor"
                  strokeWidth="4"
                ></circle>
                <path
                  className="opacity-75"
                  fill="currentColor"
                  d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"
                ></path>
              </svg>
            ) : (
              "Create Account"
            )}
          </button>

          <br />
          <div className="flex flex-col justify-between mt-4">
            <Link href="/auth/login" className="text-customprimary">
              Already have an account? Login
            </Link>
          </div>

          {errorMessage && (
            <div className="bg-red-100 text-red-700 p-3 rounded mt-4">
              <p>{errorMessage}</p>
            </div>
          )}
          {successMessage && (
            <div className="bg-green-100 text-green-700 p-3 rounded mt-4">
              <p>{successMessage}</p>
            </div>
          )}
        </form>
      </Form>
    </div>
  );
}
