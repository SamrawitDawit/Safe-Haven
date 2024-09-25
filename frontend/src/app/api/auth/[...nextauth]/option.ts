import type { NextAuthOptions, User as NextAuthUser } from "next-auth";
import GoogleProvider from "next-auth/providers/google";
import CredentialsProvider from "next-auth/providers/credentials";

interface ExtendedUser extends NextAuthUser {
  accessToken?: string;
  refreshToken?: string;
}
interface CustomToken {
  accessToken?: string;
  refreshToken?: string;
}

interface CustomSession {
  accessToken?: string;
  refreshToken?: string;
}
export const options: NextAuthOptions = {
  providers: [
    GoogleProvider({
      clientId: process.env.GOOGLE_CLIENT_ID as string,
      clientSecret: process.env.GOOGLE_CLIENT_SECRET as string,
    }),
    CredentialsProvider({
      name: "Credentials",
      credentials: {
        identifier: { label: "Email or Phone", type: "text" },
        password: { label: "Password", type: "password" },
      },
      async authorize(credentials) {
        // Determine if identifier is email or phoneNumber
        const isEmail = credentials?.identifier.includes("@");

        // Prepare request body based on input type
        const body = isEmail
          ? { email: credentials?.identifier, password: credentials?.password }
          : {
              phoneNumber: credentials?.identifier,
              password: credentials?.password,
            };

        // Call your backend API for authentication
        const res = await fetch("http://localhost:8080/auth/login", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(body),
        });

        if (!res.ok) {
          // Return null if the login fails
          return null;
        }

        const user = await res.json();

        // If user exists and tokens are available, return the user object
        if (user?.data?.accessToken && user?.data?.refreshToken) {
          return {
            id: user.data.id,
            name: user.data.name,
            email: user.data.email, // Use email if provided, or phone number if available
            accessToken: user.data.accessToken,
            refreshToken: user.data.refreshToken,
          };
        }

        return null; // Return null if no user is found
      },
    }),
  ],
  pages: {
    signIn: "/auth/login", // Redirect to your custom sign-in page
    // Error page if there are login issues
  },
  callbacks: {
    // Add the JWT callback to attach the tokens to the session
    async jwt({ token, user }) {
      if (user) {
        const extendedUser = user as ExtendedUser;
        token.accessToken = extendedUser.accessToken;
        token.refreshToken = extendedUser.refreshToken;
      }
      return token;
    },
    // Add the session callback to include the accessToken in the session
    async session({ session, token }) {
      if (token?.accessToken) {
        session.accessToken = token.accessToken as string;
        session.refreshToken = token.refreshToken as string;
      }
      return session;
    },
  },
  session: {
    strategy: "jwt", // Use JWT strategy to handle sessions
  },
};
