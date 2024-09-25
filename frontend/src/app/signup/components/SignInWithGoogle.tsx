import Image from "next/image";
import { signIn } from "next-auth/react";

const SignInButton = () => {
  const handleGoogleSignIn = async () => {
    await signIn("google", { callbackUrl: "/" }); // Redirect to homepage after login
  };

  return (
    <button
      onClick={handleGoogleSignIn}
      className="border border-customprimary w-full flex justify-center p-4 my-2"
    >
      <Image
        src="/google.png"
        alt="google"
        width={24}
        height={24}
        className="w-6 h-6"
      />
      <p className="text-center font-poppins text-xl text-customblue px-2">
        Sign In with Google
      </p>
    </button>
  );
};

export default SignInButton;
