import NextAuth from "next-auth";
import { options } from "./option"; // import your authOptions

const handler = NextAuth(options);

export { handler as GET, handler as POST };
