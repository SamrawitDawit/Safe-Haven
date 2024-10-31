import Header from "@/components/header";
import land from "./assets/landingimg.png";
import Image from "next/image";

export default function Landing() {
  return (
    <div>
      {/* the header  */}
      <Header />
      <div className="flex justify-between items-center p-4 bg-customprimary">
        <div className="space-y-12">
          <div className="text-container">
            <p className="text-white text-center text-4xl whitespace-normal ">
              Empowering Your Journey to Healing and Justice
            </p>
          </div>
          <p className="text-white">
            Welcome to your safe-haven, a safe and supportive space designed to
            guide you through your healing journey. Our app offers valuable
            resources, practical tips, and educational content to help you
            understand your options, build resilience, and connect with support
            networks. Whether you're seeking information or need someone to turn
            to, we are here to provide the knowledge and support you need every
            step of the way.
          </p>
        </div>
        <Image src={land} alt="support" className="object-contain" />
      </div>
    </div>
  );
}
