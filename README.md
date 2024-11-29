# Buildas Technical Assessment - Go Project

## Overview

This project provides a basic user management system with features such as:

    User registration
    User login
    User listing
    Password verification
    Mock service testing for user registration

It uses the Gin framework for handling HTTP requests, bcrypt for hashing passwords, and PostgresSQL as the database. The application includes service, repository, and handler layers to handle the business logic, data access, and HTTP routing respectively.

## Features

- **User Registration**: Users can register by providing a username and password.
- **User Authentication**: Login functionality with password validation.
- **Get All Users**: Fetch all registered users from the database.
- **Mocking Services**: Use of mocks for testing services without depending on a real database.

## Project Structure

```
my-gin-app/
├── cmd/                              # Entry point for the application
│   └── my-gin-app/                   # Main application files
│       └── main.go                   # Main entry point
├── pkg/                              # Application logic
│   └── models/                       # Data models (e.g., User model)
│   └── services/                     # Service layer logic
│       └── user_service.go           # User service logic
│       └── user_service_interface.go # Interface for user service
│   └── handlers/                     # Handlers for HTTP routes
│       └── user_handler.go           # Handler for user routes
│   └── middlewares/                  # Handlers for HTTP routes
│       └── auth_middleware.go        # Handler for user routes
│   └── utils/                        # Utility functions
├── internal/                         # Internal code, private to the application
│   └── config/                       # Repository layer for data access
│       └── config.go                 # User data access logic
│       └── load.go                   # Load the environment variables from a .env file    
│   └── database/                     # Repository layer for data access
│       └── db.go                     # It opens a connection to the database using GORM
│   └── repository/                   # Repository layer for data access
│       └── user_repo.go              # User data access logic    
├── migrations/                       # Database migrations
├── tests/                            # Test cases
│   └── pkg/                           
│       └── handlers_test/             
│           └── user_handler_test.go  # Test cases for user handler
├── web/                              
│   └── assets/                       
│       └── styles.css                # Stylesheet for the web interface
│   └── templates/
│       └── home.html                 # HTML template for the home interface
│       └── error.html                # HTML template for error pages
│       └── login.html                # HTML template for the login form
│       └── register.html             # HTML template for the registration form
├── .env                              # Environment variables
├── docker-compose.yml                # Docker compose file
├── Dockerfile                        # Dockerfile
├── go.mod                            # Go module definition
├── go.sum                            # Go checksum file
└── README.md                         # Project README

```

## Requirements

- Go 1.20.
- PostgresSQL for the database.
- Docker for containerization.

## Setup

### 1. Clone the repository

```bash
 git clone https://github.com/tdazle007/buildas-assessment.git
cd buildas-assessment-go

```

### 2. Install dependencies

```bash
  go mod tidy
```

### 3. Setup database

Ensure you have PostgresSQL set up and running. Update the database connection details in .env.

Example .env file:

```
POSTGRES_USER=postgres
POSTGRES_PASSWORD=secret
POSTGRES_DB=go_crud_app
POSTGRES_HOST=database
POSTGRES_PORT=5432
```

### 4. Build the Docker image

```bash
 buildas-assessment-go> docker compose up --build 
```

### 5. Test the application

```bash 
 buildas-assessment-go> go test -v ./...
```

### 6. Access the application

You can access the application at http://localhost:8080

### 7. Endpoints

The application provides the following endpoints:

- **/api/v1/user/register**: Register a new user.
- **/api/v1/user/login**: Authenticate a user.
- **/api/v1/user/home**: Get all registered/add more users.

### 8. Services

The application follows a layered architecture with services encapsulating the business logic.
- **UserService**: Handles user-related operations such as registration, fetching users, and password checks.
- **UserRepository**: Interfaces with the database to persist and retrieve user data.
- **UserServiceInterface**: Defines the methods required for user management, which are then implemented by ```UserService``` and mocked in tests.

### 9. Handlers

The handler package contains the HTTP handler functions that interact with the service layer:

- **/register**: Handles user registration by receiving data via a POST request, invoking ```RegisterUser```, and redirecting the user to a success page on success.
- **/login**: Handles user login by receiving data via a POST request, invoking ```LoginUser```, and redirecting the user to a success page on success.
- **/home**: Handles the home page request, invoking ```GetAllUsers```, and rendering the home template with the list of users.

### 10. Testing

Testing is implemented using the ```testify``` package for mocking and assertions.

This ```README.md``` provides a comprehensive overview of the project, installation steps, and explanations of the components and testing. It will guide other developers through setting up, running, and understanding the codebase effectively. 😊
