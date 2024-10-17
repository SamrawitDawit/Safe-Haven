import NextAuth, { DefaultSession } from "next-auth";

// Extend the Session interface to include accessToken
declare module "next-auth" {
  interface Session {
    accessToken?: string; // Adding the accessToken field to the Session type
    refreshToken?: string; // Adding the refreshToken field to the Session type
    user: {
      id: string;
      name: string;
      email: string;
    } & DefaultSession["user"];
  }
  interface User {
    id: string;
    accessToken?: string;
    refreshToken?: string;
  }
}

declare module "next-auth/jwt" {
  interface JWT {
    accessToken?: string;
    refreshToken?: string;
  }
}


