# JWT Authentication Setup

## Overview
This application now includes a complete JWT (JSON Web Token) authentication system with user registration, login, and protected routes.

## Environment Variables

Create a `.env` file in the backend directory with the following variables:

```env
PORT=8080
DATABASE_URL=user:password@tcp(localhost:3306)/sociomile_db?charset=utf8mb4&parseTime=True&loc=Local
JWT_SECRET=your-secret-key-here-change-in-production
```

**Important:** Generate a secure JWT_SECRET for production using:
```bash
openssl rand -base64 32
```

## API Endpoints

### Public Endpoints (No Authentication Required)

#### 1. Register a New User
**POST** `/api/v1/auth/register`

Request Body:
```json
{
  "email": "user@example.com",
  "password": "password123",
  "name": "John Doe"
}
```

Response (201 Created):
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "email": "user@example.com",
    "name": "John Doe"
  }
}
```

#### 2. Login
**POST** `/api/v1/auth/login`

Request Body:
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

Response (200 OK):
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "email": "user@example.com",
    "name": "John Doe"
  }
}
```

#### 3. Refresh Token
**POST** `/api/v1/auth/refresh`

Request Body:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

Response (200 OK):
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Protected Endpoints (Authentication Required)

#### 4. Get User Profile
**GET** `/api/v1/auth/profile`

Headers:
```
Authorization: Bearer <your-jwt-token>
```

Response (200 OK):
```json
{
  "id": 1,
  "email": "user@example.com",
  "name": "John Doe"
}
```

## Using JWT Middleware

To protect any route with JWT authentication, use the `middleware.JWTAuth` middleware:

```go
import (
    "DewaSRY/sociomile-app/pkg/middleware"
    "github.com/go-chi/chi/v5"
)

func RegisterProtectedRoutes(r chi.Router) {
    r.Group(func(r chi.Router) {
        r.Use(middleware.JWTAuth)
        
        // All routes in this group require authentication
        r.Get("/protected-resource", yourHandler)
        r.Post("/protected-action", yourHandler)
    })
}
```

## Accessing User Info in Handlers

When a request passes through the JWT middleware, user information is stored in the request context:

```go
func YourHandler(w http.ResponseWriter, r *http.Request) {
    // Get authenticated user ID
    userID, ok := r.Context().Value("userID").(uint)
    if !ok {
        // Handle error
        return
    }
    
    // Get authenticated user email
    email, ok := r.Context().Value("email").(string)
    if !ok {
        // Handle error
        return
    }
    
    // Use userID and email...
}
```

## Testing with cURL

### Register a new user:
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123",
    "name": "Test User"
  }'
```

### Login:
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'
```

### Access protected endpoint:
```bash
curl -X GET http://localhost:8080/api/v1/auth/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE"
```

## Security Features

1. **Password Hashing**: Passwords are hashed using bcrypt before storage
2. **Token Expiration**: JWT tokens expire after 24 hours
3. **Secure Token Validation**: Tokens are validated with HMAC-SHA256
4. **Context-Based Auth**: User info is passed via request context, not global variables

## Running the Application

```bash
cd backend
go run cmd/main.go
```

The server will start on the port specified in your `.env` file (default: 8080).

## Database Migration

The application automatically creates the `users` table on startup using GORM's AutoMigrate feature.

## Project Structure

```
backend/
├── cmd/
│   └── main.go                    # Application entry point
├── internal/
│   ├── database/
│   │   └── connection_database.go # Database connection
│   ├── handlers/
│   │   └── auth_handler.go        # Authentication HTTP handlers
│   ├── routers/
│   │   └── authentication_router.go # Auth route definitions
│   └── services/
│       └── auth_service.go        # Authentication business logic
├── pkg/
│   ├── dtos/
│   │   ├── requestdto/
│   │   │   └── auth_request.go    # Request DTOs
│   │   └── responsedto/
│   │       └── auth_response.go   # Response DTOs
│   ├── lib/
│   │   └── jwt/
│   │       └── jwt.go             # JWT utility functions
│   ├── middleware/
│   │   └── middleware.go          # JWT authentication middleware
│   └── models/
│       └── user.go                # User model
└── .env.example                   # Environment variables template
```
