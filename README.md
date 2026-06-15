# TaskFlow API

REST API for task management built with Go, PostgreSQL and JWT authentication.

## Features

* User registration
* User login
* JWT authentication
* Create task
* Get all user tasks
* Get task by ID
* Update task
* Delete task

## Tech Stack

* Go
* PostgreSQL
* JWT
* Chi Router
* pgx
* bcrypt

  

## Project Structure

```text
cmd/api              - application entry point
internal/auth        - JWT generation and validation
internal/config      - application configuration
internal/handler     - HTTP handlers
internal/middleware  - authentication middleware
internal/model       - domain models
internal/repository  - database layer
internal/service     - business logic
migrations           - database migrations
```

## API Endpoints

### Authentication

POST /auth/register

POST /auth/login

### Profile

GET /profile

### Tasks

GET /tasks

GET /tasks/{id}

POST /tasks

PUT /tasks/{id}

DELETE /tasks/{id}

## Author

Bogdan Voronyk
