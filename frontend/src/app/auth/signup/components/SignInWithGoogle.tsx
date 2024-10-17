
import { useRouter } from "next/navigation";


const SignInButton = () => {

  const router = useRouter();

  
  const handleGoogleSignIn = async () => {
    try {
      // Redirect to backend's Google Login endpoint
      window.location.href = "http://localhost:8080/auth/google";
    } catch (error) {
      console.error("Google Sign-In failed:", error);
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
