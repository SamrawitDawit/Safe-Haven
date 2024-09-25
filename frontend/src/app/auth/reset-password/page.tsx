"use client";
import React, { useState } from "react";
import { useRouter, useSearchParams } from "next/navigation";

const ResetPasswordPage = () => {
  const [newPassword, setNewPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");
  const searchParams = useSearchParams(); // For extracting query parameters
  const token = searchParams.get("token"); // Assuming token is passed as a query parameter

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    // Reset any previous error or success messages
    setError("");
    setSuccess("");

    if (newPassword !== confirmPassword) {
      setError("Passwords do not match!");
      return;
    }

    if (!token) {
      setError("Invalid or missing token.");
      return;
    }

    try {
      const response = await fetch(
        "http://localhost:8080/auth/reset-password",
        {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ newPassword, token }),
        }
      );

      const data = await response.json();

      if (response.ok) {
        setSuccess("Password has been successfully updated.");
      } else if (data.error === "Invalid token") {
        setError("Invalid or expired token. Please request a new reset link.");
      } else {
        setError(data.message || "Failed to reset password.");
      }
    } catch (error) {
      setError("Something went wrong. Please try again later.");
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-white">
      <div className="bg-white p-8 rounded-lg shadow-md w-full max-w-md">
        <h2 className="text-2xl font-semibold text-gray-800 mb-6">
          Set a new password
        </h2>
        <p className="text-gray-600 mb-6">
          Create a new password. Ensure it differs from previous ones for
          security.
        </p>

        {error && <div className="text-red-500 text-sm mb-4">{error}</div>}
        {success && (
          <div className="text-green-500 text-sm mb-4">{success}</div>
        )}

        <form onSubmit={handleSubmit} className="space-y-6">
          <div>
            <label
              htmlFor="new-password"
              className="block text-sm font-medium text-gray-700"
            >
              New Password<span className="text-red-500">*</span>
            </label>
            <input
              type="password"
              id="new-password"
              value={newPassword}
              onChange={(e) => setNewPassword(e.target.value)}
              className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
              placeholder="Enter a new password"
              required
            />
          </div>

          <div>
            <label
              htmlFor="confirm-password"
              className="block text-sm font-medium text-gray-700"
            >
              Confirm Password<span className="text-red-500">*</span>
            </label>
            <input
              type="password"
              id="confirm-password"
              value={confirmPassword}
              onChange={(e) => setConfirmPassword(e.target.value)}
              className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
              placeholder="Rewrite the password"
              required
            />
          </div>

          <button
            type="submit"
            className="w-full bg-[#96D1D5] text-white py-2 px-4 rounded-md hover:bg-teal-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-teal-500"
          >
            Update Password
          </button>
        </form>
      </div>
    </div>
  );
};

export default ResetPasswordPage;
