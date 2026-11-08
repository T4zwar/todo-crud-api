# Todo CRUD API

A RESTful Todo management API built with Go, Gin, and SQLite. This project supports creating, retrieving, updating, and deleting Todo records.

## Technologies Used

- **Go** – Backend programming language
- **Gin** – HTTP web framework
- **SQLite** – Database
- **Thunder Client** – API testing

## Features

- Create a new Todo
- Retrieve all Todos
- Retrieve a Todo by ID
- Update an existing Todo
- Delete a Todo
- Validate incoming requests
- Handle errors and appropriate HTTP status codes
- Store Todo records in SQLite

## Project Structure

```text
todo-crud-api/
├── database/
│   └── db.go
├── handlers/
│   └── todo.go
├── models/
│   └── todo.go
├── .gitignore
├── go.mod
├── go.sum
└── main.go
```

## Installation and Setup

### 1. Clone the repository

```bash
git clone YOUR_REPOSITORY_URL
cd todo-crud-api
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Run the application

```bash
go run .
```

The server will start at:

`http://localhost:8080`

The SQLite database is created automatically when the application starts.

## API Endpoints

| Method | Endpoint | Description |
|---|---|---|
| POST | `/todos` | Create a Todo |
| GET | `/todos` | Retrieve all Todos |
| GET | `/todos/:id` | Retrieve a Todo by ID |
| PUT | `/todos/:id` | Update a Todo |
| DELETE | `/todos/:id` | Delete a Todo |

## HTTP Status Codes

| Status Code | Meaning |
|---|---|
| 200 OK | Request completed successfully |
| 201 Created | Todo created successfully |
| 400 Bad Request | Invalid request data or Todo ID |
| 404 Not Found | Todo does not exist |
| 500 Internal Server Error | Unexpected server error |


## Example Requests

### Create a Todo

**POST** `/todos`

```json
{
  "title": "Learn Golang",
  "completed": false
}
```

### Update a Todo

**PUT** `/todos/1`

```json
{
  "title": "Master Go CRUD API",
  "completed": true
}
```

## Testing

All five CRUD endpoints were tested using Thunder Client in Visual Studio Code.

## Author

Developed as a Go CRUD API project.
