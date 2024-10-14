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
import SignInWithGoogle from "../signup/components/SignInWithGoogle";
import { signIn } from "next-auth/react";

const formSchema = z
  .object({
    fullName: z.string().min(2, {
      message: "Full name must be at least 2 characters.",
    }),
    phoneNumber: z.string().nonempty({
      message: "Invalid phone.",
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

export default function SignUpwithPhone() {
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const [successMessage, setSuccessMessage] = useState<string | null>(null);

  const form = useForm<FormData>({
    resolver: zodResolver(formSchema),
  });
  const router = useRouter();

  // Handle form submission for email registration
  const onSubmit = async (data: FormData) => {
    console.log(data);
    try {
      const registerData: RegisterUserData = {
        fullName: data.fullName,
        phoneNumber: data.phoneNumber, // Email registration only
        password: data.password,
        userType: "normal", // Always "normal"
        category: data.category,
        language: data.language,
        bio: data.bio, // Optional bio
      };
      console.log(registerData, "registered");
      // Call the registration API for email-based registration
      const response = await registerUser(registerData);
      setSuccessMessage("Registration successful!");
      setErrorMessage(null);

      await signIn("credentials", {
        redirect: false,
        identifier: data.phoneNumber, // Sign in using phone number
        password: data.password,
      });
      router.push("/");
    } catch (error: any) {
      setErrorMessage("Failed to register. Please try again.");
      setSuccessMessage(null);
    }
  };

  return (
    <div className="container mx-auto w-2/5 py-4 h-fit">
      <p className="text-3xl text-center font-bold font-sans"> Register </p>
      <SignInWithGoogle />
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
            name="phoneNumber"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Phone Number</FormLabel>
                <FormControl>
                  <Input type="string" placeholder="" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <Link href="/auth/signup" className="text-slate-500 text-sm">
            Use Email instead
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
                    <SelectTrigger className="">
                      <SelectValue placeholder="Select Preferred language" />
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
                    <SelectTrigger className="">
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
            className="bg-customprimary text-white p-3 rounded"
          >
            Register
          </button>
          <br />

          <div className="flex flex-col justify-between mt-4">
            <Link href="/login" className="text-customprimary">
              Already have an account? Login
            </Link>
          </div>
          {errorMessage && <p className="text-red-500">{errorMessage}</p>}
          {successMessage && <p className="text-green-500">{successMessage}</p>}
        </form>
      </Form>
    </div>
  );
}
