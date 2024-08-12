# Task Management System API

## Overview

This project implements a RESTful API for a Task Management System using Go. The API allows users to create, retrieve, update, and delete tasks. Each task has a title, description, and status (e.g., "todo," "in progress," "completed"). The API also supports marking tasks as complete and filtering tasks based on their status.

## Features

- **CRUD Operations**: Create, Retrieve, Update, and Delete tasks.
- **Status Management**: Mark tasks as complete.
- **Pagination and Filtering**: Retrieve tasks with pagination and filtering capabilities.
- **Error Handling**: Basic validation and error handling for invalid requests.
- **JWT Authentication** (Optional): User authentication for creating, updating, or deleting tasks.
- **Dockerization** (Optional): Docker image for easy deployment.

## Running the Application

- Install Dependencies

```bash
go mod tidy
```

- Run the Application

```bash
go run cmd/main.go
```

## Optional: Dockerization

To build and run the Docker container:

- Build the Docker Image

```bash
docker build -t task-manager .
```

- Run the Docker Container

```bash
docker run -p 8080:8080 task-manager
```

## Running Tests

    ```bash
    go test ./...
    ```

## License

This project is licensed under the MIT License.

## Acknowledgments

- Go Programming Language
- Gorilla Mux Router
- Docker
