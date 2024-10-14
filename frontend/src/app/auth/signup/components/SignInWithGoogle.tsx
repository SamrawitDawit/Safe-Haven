import { useState } from "react";

const SignInButton = () => {
  const [user, setUser] = useState(null);

  const handleGoogleSignIn = async () => {
    try {
      const response = await fetch("http://localhost:8080/auth/google", {
        method: "GET",
      });

      if (response.ok) {
        const data = await response.json();

        localStorage.setItem("accessToken", data.data.accessToken);
        localStorage.setItem("refreshToken", data.data.refreshToken);

        setUser(data.data.user);

        window.location.href = "/";
      } else {
        console.error("Google login failed:", response.statusText);
      }
    } catch (error) {
      console.error("Error during Google login:", error);
    }
  };

  return (
    <div className="flex justify-center">
      <button
        onClick={handleGoogleSignIn}
        type="button"
        className="google-sign-in-button"
      >
        Use Google
      </button>
    </div>
  );
};

export default SignInButton;
