'use client'
import React, { useState, useEffect} from 'react'
import { useSearchParams } from 'next/navigation'
import { useForm, SubmitHandler } from 'react-hook-form'
import { useRouter, usePathname } from 'next/navigation'
import { signIn } from 'next-auth/react'
import Link from "next/link";

interface LoginValues {
  email: string;
  password: string;
}

const LoginEmail = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginValues>();
  const router = useRouter();
  const [error, setError] = useState<string | null>(null);
  const [usertype, setUsertype] = useState<string>('');

  const searchParams = useSearchParams();
  const message = searchParams ? searchParams.get("message") : null;

  const onSubmit: SubmitHandler<LoginValues> = async (data) => {
    try {
      const res = await signIn("credentials", {
        redirect: false,
        email: data.email,
        password: data.password,
        usertype: 'normal',
      });

      if (res?.error) {
        setError(res.error);
      } else if (res?.ok) {
        setUsertype('normal');
        router.push("/");
      }
    } catch (error) {
      setError("Login Failed! Please try again. ");
    }
  };

  const handleGoogleSignIn = async () => {
    const res = await signIn("google", {
      callbackUrl: "/",
      usertype: 'normal',
    });
    setUsertype('normal');
  };


  return (
    <div className="flex justify-center mt-0 ">
      <div className="flex flex-col gap-4 w-[480px] mt-50">
        <div className="flex justify-center">
          <h4 className="text-red-600 font-bold text-[20px] font-epilogue">{message}</h4>
        </div>
        <h2 className="text-[36px] font-black font-poppins">Log In</h2>

        <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-1">
          <label htmlFor="email" className="input-label">
            Email:
          </label>
          <input
            type="email"
            id="email"
            {...register("email", { required: "Email is required" })}
            placeholder="Enter your email"
            className="border rounded p-2"
          />
          {errors.email && (
            <p className="text-red-500">{errors.email.message}</p>
          )}

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

          {error && <p className="text-red-500">{error}</p>}


          <div className="border rounded-md border-[#CCCCF5] border-3 rounded- px-10 py-3">
            <Link 
            href="/loginPhone"
            className='flex gap-2 text-[#4640DE] font-epilogue font-semibold text-base w-[100%] justify-center'
            >
                Log in with Phone number
            </Link>
          </div>

          <div className="border rounded-md border-[#CCCCF5] border-3 rounded- px-10 py-3">
            <Link 
            href="/loginAnonymous"
            className='flex gap-2 text-[#4640DE] font-epilogue font-semibold text-base w-[100%] justify-center'
            >
                Log in Anonymously 
            </Link>
          </div>

          <div className="border rounded-md border-[#CCCCF5] border-3 rounded- px-10 py-3">
            <button
              onClick={() => handleGoogleSignIn()}
              className="flex gap-2 text-[#4640DE] font-epilogue font-semibold text-base w-[100%] justify-center"
            >
              <GoogleBar />
              Sign in with Google
            </button>
          </div>

          <input
            type="submit"
            value="Sign In"
            className="bg-[#4640DE] px-6 py-3 rounded-[80px] text-white font-epilogue font-bold text-[12px] mt-4"
          />
        </form>
        <div className="w-fit h-fit flex gap-3">
          <p className="text-[16px] font-epilogue font-normal leading-[25.6px] text-[#202430] w-[203px] h-[26px] opacity-[0.7]">
            Don't have an account?
          </p>
          <Link
            href="/signup" 
            className="text-[16px] font-inter font-semibold leading-[24px] text-[#4640DE] h-[24px]"
          >
            Sign up
          </Link>
        </div>
      </div>
    </div>
  )
}


function GoogleBar() {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      x="0px"
      y="0px"
      width="20"
      height="20"
      viewBox="0 0 48 48"
    >
      <path
        fill="#FFC107"
        d="M43.611,20.083H42V20H24v8h11.303c-1.649,4.657-6.08,8-11.303,8c-6.627,0-12-5.373-12-12c0-6.627,5.373-12,12-12c3.059,0,5.842,1.154,7.961,3.039l5.657-5.657C34.046,6.053,29.268,4,24,4C12.955,4,4,12.955,4,24c0,11.045,8.955,20,20,20c11.045,0,20-8.955,20-20C44,22.659,43.862,21.35,43.611,20.083z"
      ></path>
      <path
        fill="#FF3D00"
        d="M6.306,14.691l6.571,4.819C14.655,15.108,18.961,12,24,12c3.059,0,5.842,1.154,7.961,3.039l5.657-5.657C34.046,6.053,29.268,4,24,4C16.318,4,9.656,8.337,6.306,14.691z"
      ></path>
      <path
        fill="#4CAF50"
        d="M24,44c5.166,0,9.86-1.977,13.409-5.192l-6.19-5.238C29.211,35.091,26.715,36,24,36c-5.202,0-9.619-3.317-11.283-7.946l-6.522,5.025C9.505,39.556,16.227,44,24,44z"
      ></path>
      <path
        fill="#1976D2"
        d="M43.611,20.083H42V20H24v8h11.303c-0.792,2.237-2.231,4.166-4.087,5.571c0.001-0.001,0.002-0.001,0.003-0.002l6.19,5.238C36.971,39.205,44,34,44,24C44,22.659,43.862,21.35,43.611,20.083z"
      ></path>
    </svg>
  );
}

export default LoginEmail