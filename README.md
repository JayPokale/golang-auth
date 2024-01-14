# Golang-auth

## Description

This is a golang project for authentication

## Technologies Used

Go
Fiber (or other web framework)
MongoDB (or your preferred database)
JWT for authentication
Bcrypt

## Getting Started

Provide instructions on how to set up and run your project locally.

### Clone the repository:

```sh
git clone https://github.com/jaypokale/golang-auth.git
```

### Install dependencies:

```sh
cd golang-auth
go mod download
```

Set up environment variables:
Create a `.env` file and configure the required environment variables.

Run the application:

```sh
go run main.go
```

Your project should now be running locally.

## API Endpoints
List and describe the main API endpoints of your application.
│
├── Auth
│   ├── POST http://localhost:3000/auth/signup
│   │   └── Body: {
│   │               "name": "John Doe",
│   │               "email": "john@example.com",
│   │               "password": "password",
│   │               "phone": "1234567890",
│   │               "gender": "Male",
│   │               "howDidYouHear": "Friends",
│   │               "city": "Mumbai",
│   │               "state": "Maharashtra"
│   │             }
│   └── POST http://localhost:3000/auth/login
│       └── Body: {
│                   "email": "john@example.com",
│                   "password": "password",
│                 }
│
├── User
│   ├── GET http://localhost:3000/users
│   │   └── Header: { key }
│   ├── PUT http://localhost:3000/users
│   │   └── Header: { key }
│   └── DELETE http://localhost:3000/users
│       └── Header: { key }
│
├── Admin
│   ├── GET http://localhost:3000/admin
│   │   └── Header: { key }
│   ├── GET http://localhost:3000/admin/{id}
│   │   └── Header: { key }
│   ├── PUT http://localhost:3000/admin/{id}
│   │   ├── Header: { key }
│   │   └── Body: {
│   │               "name": "Changed Name",
│   │               "email": "changed@example.com",
│   │             }
│   └── DELETE http://localhost:3000/admin/{id}
│       └── Header: { key }