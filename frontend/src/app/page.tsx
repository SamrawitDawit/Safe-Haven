import Image from "next/image";
import LoginEmail from "./loginEmail/page";

export default function Home() {
  return (
    <div className=" items-center justify-items-center  p-8 pb-20 gap-16 font-[family-name:var(--font-geist-sans)]">
      <LoginEmail/>
    </div>
  );
}
