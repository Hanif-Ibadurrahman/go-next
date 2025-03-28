## Customer Repository Backend

This project implements a user repository backend in Go, providing functionality to search users, bank accounts, pockets, and term deposits. The repository is located in backend/pkg/api/v1/user/repository/impl

## Setup

1. Clone the Repository

```bash
   git clone <repository-url>
   cp ./backend/.env.example ./backend/.env
   docker compose up -d --build
```

Frontend
http://localhost:3000

Backend
http://localhost:8000

Swagger
http://localhost:8000/docs/

## Display

![Login](/docs/login.png)
Account
Username: customer1
Password: Password123

![Feature Search](/docs/search.png)
Steps

- Input data search
- Click button search

## Running Unit Tests

Unit tests are located in user_impl_test.go and use go-sqlmock to simulate database interactions, eliminating the need for a real database during testing. The tests cover key functions: Search

Steps
Navigate to the Package:

```bash
    cd backend/pkg/api/v1/user/repository/impl
```

Run all tests with verbose output:

```bash
    go test -v
```
