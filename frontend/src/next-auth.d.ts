import NextAuth, { DefaultSession } from "next-auth";

// Extend the Session interface to include accessToken
declare module "next-auth" {
  interface Session {
    accessToken?: string; // Adding the accessToken field to the Session type
    refreshToken?: string; // Adding the refreshToken field to the Session type
  }
}
