"use client";
import { useState } from "react";

const ForgotPasswordPage = () => {
  const [email, setEmail] = useState("");
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    // Reset any previous error or success messages
    setError("");
    setSuccess("");

    if (!email) {
      setError("Email is required");
      return;
    }

    try {
      // Log the email before sending to ensure the right data is being sent
      console.log("Sending email:", email);

      // Sending the email as a plain text string
      const response = await fetch(
        "http://localhost:8080/auth/forgot-password",
        {
          method: "POST",
          headers: { "Content-Type": "text/plain" }, // Set content-type to text/plain
          body: JSON.stringify(email), // Send plain text email
        }
      );

      if (response.ok) {
        setSuccess("Password reset email has been sent!");
      } else {
        const data = await response.json();
        console.log("Error response:", data);
        setError(data.message || "Failed to send email.");
      }
    } catch (error) {
      setError("Something went wrong. Please try again later.");
    }
  };

  return (
    <div className="flex flex-col items-center justify-center h-screen bg-gray-50">
      <div className="w-full max-w-md p-6 bg-white rounded-lg shadow-md">
        <h2 className="mb-4 text-2xl font-semibold text-center text-gray-800">
          Forgot password
        </h2>
        <p className="mb-6 text-sm text-center text-gray-600">
          Please enter your email to reset the password
        </p>
        <form onSubmit={handleSubmit} className="space-y-6">
          <div>
            <label
              htmlFor="email"
              className="block text-sm font-medium text-gray-700"
            >
              Email Address<span className="text-red-500">*</span>
            </label>
            <input
              type="email"
              id="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
              className="block w-full px-3 py-2 mt-1 border rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500 sm:text-sm border-gray-300"
              placeholder="Enter email address"
            />
          </div>
          {error && <p className="text-sm text-red-500">{error}</p>}
          {success && <p className="text-sm text-green-500">{success}</p>}
          <div>
            <button
              type="submit"
              className="w-full px-4 py-2 font-medium text-white bg-[#96D1D5] rounded-md hover:bg-green-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
            >
              Send Email
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default ForgotPasswordPage;
