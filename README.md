# Smart Task Planner API

Smart Task Planner is a RESTful backend service built with Go, Echo v5, GORM, and PostgreSQL. It provides JWT-based authentication and task management functionality, allowing users to securely create, retrieve, update, and delete tasks.

## Technology Stack

* Go (Golang)
* Echo v5 (HTTP Framework)
* GORM (ORM)
* PostgreSQL
* JWT (Authentication)

---

## Features

* User registration
* User login with JWT token generation
* Protected task routes
* Create task
* Get all tasks (user-specific)
* Update task (partial updates supported)
* Delete task
* Ownership protection (users can access only their tasks)

---

## Environment Variables

Create a `.env` file or configure environment variables:

```
PORT=3000
DB_DSN=host=your_host user=your_user password=your_password dbname=your_db port=5432 sslmode=require TimeZone=UTC
JWT_KEY=your_secret_key
```

---

## Installation

### 1. Clone the repository

```
git clone https://github.com/JagTheFriend/Smart-Task-Planner-App.git
cd smart-task-planner
```

### 2. Install dependencies

```
go mod tidy
```

### 3. Run the application

```
go run main.go
```

Server will start on:

```
http://localhost:3000
```

---

## API Endpoints

Base path:

```
/api/v1
```

---

### Authentication

#### Signup

```
POST /api/v1/auth/signup
```

Request Body:

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

---

#### Login

```
POST /api/v1/auth/login
```

Request Body:

```json
{
  "email": "john@example.com",
  "password": "password123"
}
```

Response:

```json
{
  "message": "<JWT_TOKEN>"
}
```

---

### Task Management

All task routes require:

```
Authorization: Bearer <JWT_TOKEN>
```

---

#### Create Task

```
POST /api/v1/task
```

Request Body:

```json
{
  "title": "Complete API",
  "description": "Implement CRUD operations",
  "deadline": "2026-02-20T18:00:00Z"
}
```

---

#### Get Tasks

```
GET /api/v1/task
```

Returns all tasks belonging to the authenticated user.

---

#### Update Task

```
PUT /api/v1/task
```

Request Body:

```json
{
  "id": 1,
  "title": "Updated Title",
  "completed": true
}
```

Only `id` is required. Other fields are optional.

---

#### Delete Task

```
DELETE /api/v1/task?id=1
```

---

## Security

* JWT-based authentication
* Route protection using middleware
* User-specific task isolation
* Secure password handling (recommended: hash before storing)

---

## Error Handling

* Structured error responses
* Proper HTTP status codes
* Panic recovery middleware enabled

