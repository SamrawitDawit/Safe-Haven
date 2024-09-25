import React from "react";
import SignUpForm from "./components/SignUpForm";

const SignUp = () => {
  return (
    <div className="container mx-auto w-2/5 py-4 h-fit">
      <p className="text-3xl text-center font-bold font-sans my-4">
        {" "}
        Register{" "}
      </p>

      <SignUpForm />
    </div>
  );
};

export default SignUp;
