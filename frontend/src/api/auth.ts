import { signIn } from "next-auth/react";

// Interface for traditional registration data
export interface RegisterUserData {
  fullName?: string;
  email?: string;
  phoneNumber?: string;
  anonymousDifferentiator?: string;
  password: string;
  userType: "normal" | "anonymous";
  category?: "General" | "Victim";
  bio?: string; // Optional bio field
  language?: string; // Optional language field
}

// Function for user registration
export const registerUser = async (userData: RegisterUserData) => {
  const body: any = {
    fullName: userData.fullName,
    password: userData.password,
    userType: userData.userType,
    category: userData.category, // General or Victim
  };

  // Conditionally add email, phone number, or anonymous differentiator
  if (userData.email) {
    body.email = userData.email;
  }
  if (userData.phoneNumber) {
    body.phoneNumber = userData.phoneNumber;
  }
  if (userData.anonymousDifferentiator) {
    body.anonymousDifferentiator = userData.anonymousDifferentiator;
  }

  // Add optional fields if provided
  if (userData.bio) {
    body.bio = userData.bio;
  }
  if (userData.language) {
    body.language = userData.language;
  }

  // Make the API call to the backend
  const response = await fetch("http://localhost:8080/auth/register", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(body),
  });

  if (!response.ok) {
    const errorResponse = await response.json();
    const errorMessage = errorResponse.error || "Failed to register user";
    console.log(errorMessage);
    throw new Error(errorMessage);
  }

  return await response.json();
};

// Google OAuth via NextAuth
export const signInWithGoogle = () => {
  signIn("google", {
    callbackUrl: "/", // Redirect after successful login
  });
};

// Interface for login data
export interface LoginUserData {
  email?: string; // Either email or phone number can be used
  phoneNumber?: string;
  password: string;
}

// Function for user login
export const loginUser = async (userData: LoginUserData) => {
  const body: any = {
    password: userData.password,
  };

  // Conditionally add email or phone number to the body
  if (userData.email) {
    body.email = userData.email;
  }
  if (userData.phoneNumber) {
    body.phoneNumber = userData.phoneNumber;
  }

  // Make the API call to the backend
  const response = await fetch("http://localhost:8080/auth/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(body),
  });

  if (!response.ok) {
    throw new Error("Failed to login");
  }

  // Parse the JSON response (accessToken and refreshToken)
  const data = await response.json();
  console.log(data);

  await signIn("credentials", {
    redirect: false,
    accessToken: data.data.accessToken,
    refreshToken: data.data.refreshToken,
  });
  return data;
};
