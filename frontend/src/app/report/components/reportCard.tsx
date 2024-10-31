import Image, { StaticImageData } from "next/image";

interface ReportCardProps {
  title: string;
  description: string;
  status: string;
  location: string;
  image: StaticImageData;
  resolution_status: string;
  category: string;
  evidence: string;
}

export default function ReportCard({
  title,
  description,
  status,
  location,
  image,
  resolution_status,
  category,
  evidence,
}: ReportCardProps) {
  return (
    <div className="flex bg-lightgreen max-w-xl rounded-md p-4">
      <div className="flex flex-col">
        <p className="text-white"> Case: {title}</p>
        <p className="text-white">Description: {description}</p>
        <p className="text-white">Location: {location}</p>
        <p className="text-white"> Category: {category}</p>
        <p className="text-white">Evidence: {evidence}</p>
      </div>
      <div className="flex flex-col">
        <Image src={image} alt="profile" width={100} height={100} />
        <p className="text-white">Case Status: {status}</p>
        <p className="text-white">Resolution Status: {resolution_status}</p>
      </div>
    </div>
  );
}
