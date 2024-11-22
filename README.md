# Simple REST API Server with JWT Authentication

This is a **basic** REST API server written in Go (Golang). It was created as a learning project to demonstrate fundamental skills in building web servers, working with JSON, handling HTTP requests, and implementing JWT-based authentication.

> ⚠️ **Note:** This is not a production-ready application but a simple demonstration of RESTful API principles. The focus is on clean code organization and modular structure rather than advanced features or optimizations.

## Features

- **User Management**: 
  - Add, retrieve, update, and delete user records (CRUD operations).
- **JWT Authentication**: 
  - Secure endpoints using JSON Web Tokens (JWT).
- **Basic Routing**: 
  - Organized routing using the gorilla/mux router.

## Project Structure

The project is divided into clear modules to follow good practices:

go-rest-api/ │ ├── main.go # Entry point of the application ├── handlers/ # Request handlers for API endpoints │ ├── user_handler.go # Handlers for user CRUD operations │ ├── auth_handler.go # Handlers for login and authentication │ ├── middleware/ # Middleware logic │ └── auth_middleware.go # JWT authentication middleware │ ├── models/ # Data structures and in-memory storage │ ├── user.go # User and book models │ ├── auth.go # Login data and predefined user │ ├── utils/ # Utility functions │ └── jwt_utils.go # JWT generation logic 


## Endpoints

### Public Endpoints
1. POST /login  
   Authenticate a user and return a JWT token.
   user: 1
   pass: 1 
   **Payload**:
   
json
   {
       "username": "string",
       "password": "string"
   }
Response:

On success: Token string.
On failure: Error message.
Protected Endpoints (require Bearer token in the Authorization header)
GET /usersBook
Retrieve all users and their books.

GET /usersBook/{id}
Retrieve a user by ID.

POST /usersBook
Add a new user.
Payload:

{
    "id": "string",
    "name": "string",
    "age": "number",
    "name_book": "string",
    "book": {
        "theme": "string",
        "year": "number"
    }
}
PUT /usersBook/{id}
Update an existing user by ID.

DELETE /usersBook/{id}
Delete a user by ID.

How to Run
Prerequisites
Go installed (1.20+ recommended).
Steps
Clone the repository:

git clone https://github.com/Role1776/awesomeProject.git
cd go-rest-api
Run the server:

go run main.go
The server will be available at http://localhost:8080.

Notes
Data Storage: Currently, user data is stored in memory (models.Users). Data will be lost when the server restarts.
JWT Secret: The signing key (MySignKey) is hardcoded and not environment-secure.
Error Handling: Simplistic error handling, mainly for learning purposes.
Next Steps (Future Improvements)
Replace in-memory storage with a database (e.g., PostgreSQL).
Use environment variables for sensitive data like JWT signing keys.
Add unit tests to improve code quality.
Add Docker support for easier deployment.
Purpose
This project is intentionally simple and serves as a demonstration of basic concepts. It is meant as a starting point for more advanced server development.
