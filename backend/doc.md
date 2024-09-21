# Backend Documentation
## Overview
This backend service provides RESTful APIs for handling user authentication and authorization, including traditional email/password login as well as OAuth with Google. It also features role-based access control (RBAC), secure token handling (JWT), and email services for account recovery.

The project is built using Go with the Gin web framework. It uses MongoDB as the database and supports secure and scalable JWT-based authentication with role-based middleware.

## Project Structure
The main directories are structured as follows:

```graphql
backend/
│
├── config/             # Configuration for environment variables and app settings
├── controllers/        # Controllers that handle HTTP requests and responses
├── delivery/           # Main entry point for HTTP server and routes
├── infrastructure/     # Core services like JWT, middleware, and email service
├── usecases/           # Business logic and application flow
├── utils/              # Utility functions and error response handling
└── tests/              # Test cases for unit testing
```
## Key Features
### Authentication
   * Register & Login: Supports user registration and login with email and password.
   * Google OAuth: Allows users to log in via their Google account.
   * JWT-based Authentication: Issues JWT tokens upon successful login and supports token refresh.
   * Role-Based Access Control (RBAC): Secures routes based on user roles (e.g., Admin, Counselor).
### Middleware
   * CORS Middleware: Enables CORS for cross-origin requests.
   * Auth Middleware: Protects sensitive routes, requiring a valid JWT token.
   * Role-Based Middleware: Controls access to routes based on the user's role (e.g., Admin-only routes).
### Email Service
   * Password Recovery: Supports sending recovery emails to users who have forgotten their passwords.
### Error Handling
   * A consistent error handling mechanism is used throughout the API, leveraging utility functions to send uniform error responses.
### Environment Variables
The application relies on various environment variables. Make sure to define them properly before running the project:

   * `PORT`: The port the application runs on.
   * `DB_URI`: The MongoDB URI for the database connection.
   * `DB_NAME`: The name of the MongoDB database.
   * `JWT_SECRET`: Secret key for signing JWT tokens.
   * `EMAIL_HOST`: The email service host.
   * `EMAIL_USERNAME`: The username for the email service.
   * `EMAIL_PASSWORD`: The password for the email service.
   * `EMAIL`: The sender email address for account recovery emails.
   * `GOOGLE_CLIENT_ID`: Google OAuth Client ID.
   * `GOOGLE_CLIENT_SECRET`: Google OAuth Client Secret.
   * `USER_COLLECTION`: MongoDB collection for user data.
   * `ENCRYPT_KEY`: Key used for encryption.
   * `ALLOWED_ORIGIN`: Allowed origins for CORS.


## Setup and Installation
  ### Prerequisites
   1. Install Go (https://golang.org/)
   2. Install MongoDB or set up a MongoDB Atlas instance.
   3. Set up a Google OAuth project to obtain the `GOOGLE_CLIENT_ID` and `GOOGLE_CLIENT_SECRET`.
   4. Set up an SMTP email service (e.g., Gmail or SendGrid).
### Installation Steps
   1. Clone the repository:

```bash
    git clone https://github.com/your-repo/backend.git
```
   2. Navigate to the project directory:

```bash
cd backend
```
   3. Create a `.env` file in the root directory and fill in the environment variables as described in the Environment Variables section.

   4. Install the required Go modules:

```bash
go mod tidy
```
   5. Run the application:

```bash
go run main.go
```
## Routes
  ### Authentication Routes
   * `POST /auth/register`: Register a new user.
   * `POST /auth/login`: Login with email and password.
   * `POST /auth/refresh-token`: Refresh the JWT token.
   * `POST /auth/forgot-password`: Send a password reset email.
   * `POST /auth/reset-password`: Reset the user password.
   * `GET /auth/google`: Redirect to Google OAuth login.
   * `GET /auth/google/callback`: Handle Google OAuth callback and log in the user.

## JWT and Role-Based Access Control
  ### JWT Tokens:

   * On successful login, an access token and a refresh token are issued.
   * The access token is used to authenticate subsequent requests.
   * If the access token expires, the client can use the refresh token to get a new one.
  ### Role-Based Access Control (RBAC):

   * Some routes are restricted based on user roles (e.g., "admin", "counselor"). The `AuthMiddleware` extracts the role from the JWT token and checks whether the user has permission to access the route.
## Middleware
   * CORS Middleware: Adds the necessary headers to allow cross-origin requests.
   * Auth Middleware: Checks for a valid JWT in the `Authorization` header and validates it.
   * Admin Middleware: Restricts certain routes to users with an "admin" role.
   * Counselor Middleware: Restricts access to counselor-specific routes.
## Security Considerations
   1. Google OAuth State Parameter: Protect against CSRF attacks by generating a unique, random state token when redirecting for Google login, and validating it upon callback.
   2. JWT Security: Ensure the JWT_SECRET is a strong, unpredictable key and rotate it periodically.
## Testing
Unit tests are located in the tests directory. You can run the tests using:

```bash
go test ./...
```

## Contributing
Feel free to contribute to this project by submitting issues or pull requests. Before submitting a PR, please ensure that all tests pass.