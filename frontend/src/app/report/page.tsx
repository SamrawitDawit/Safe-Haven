import Header from "@/components/header";
import girl from "./assets/image.png";
import ReportCard from "./components/reportCard";

export default function Report() {
  return (
    <>
      <Header />
      <div className="grid grid-cols-2 gap-2 p-2 bg-customprimary">
        <ReportCard
          title="case 1"
          description="Lorem ipsum dolor sit amet, sonsectetur elit. sed do eiusmod tempor incididnt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nosturd execitation ullamco"
          location="undisclosed"
          category="undisclosed"
          evidence="none"
          resolution_status="unresolved"
          image={girl}
          status="counsellor assigned"
        />
        <ReportCard
          title="case 1"
          description="Lorem ipsum dolor sit amet, sonsectetur elit. sed do eiusmod tempor incididnt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nosturd execitation ullamco"
          location="undisclosed"
          category="undisclosed"
          evidence="none"
          resolution_status="unresolved"
          image={girl}
          status="counsellor assigned"
        />
        <ReportCard
          title="case 1"
          description="Lorem ipsum dolor sit amet, sonsectetur elit. sed do eiusmod tempor incididnt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nosturd execitation ullamco"
          location="undisclosed"
          category="undisclosed"
          evidence="none"
          resolution_status="unresolved"
          image={girl}
          status="counsellor assigned"
        />
        <ReportCard
          title="case 1"
          description="Lorem ipsum dolor sit amet, sonsectetur elit. sed do eiusmod tempor incididnt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nosturd execitation ullamco"
          location="undisclosed"
          category="undisclosed"
          evidence="none"
          resolution_status="unresolved"
          image={girl}
          status="counsellor assigned"
        />
      </div>
    </>
  );
}
