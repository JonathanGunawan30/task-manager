# Task Manager RESTful API

A simple RESTful API for managing projects and tasks, built with Go following Clean Architecture principles.

## Features
- **User Management**: Simple registration and retrieval.
- **Project Management**: Full CRUD (Create, Read, Update, Delete) operations for projects.
- **Task Tracking**: Manage tasks within projects with status, priority, and deadlines.
- **Security**:
    - API Key Authentication via `X-API-Key` header.
    - Ownership protection: Users can only access and modify their own projects and tasks (checked via `X-User-ID` header).
- **Input Validation**: Strict validation for all requests.
- **Environment-based configuration**: Easily switch between environments using `.env`.

## Tech Stack
- **Language**: Go 1.26+
- **Framework/Router**: Standard Library `http.ServeMux` (Go 1.22+ syntax)
- **Database**: MySQL with GORM
- **Logging**: `logrus`
- **Configuration**: `viper`
- **Validation**: `go-playground/validator`
- **Testing**: JetBrains HTTP Client (`.http` files)

## Prerequisites
- Go 1.26 or later
- MySQL Database

## Setup
1. **Clone the repository**:
   ```bash
   git clone https://github.com/JonathanGunawan30/task-manager
   cd task-manager
   ```
2. **Environment Configuration**:
   Create a `.env` file based on `.env.example` (if available, otherwise create one):
   ```bash
   # Example .env content
   APP_NAME=task-manager
   APP_PORT=3000
   X_API_KEY=your-secret-api-key
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=root
   DB_PASSWORD=
   DB_NAME=task_manager
   ```

3. **Database Setup**:
    - Create a database named `task_manager` in MySQL.
    - Run the migrations found in `database/migrations/` (Up SQL files).

4. **Run the Application**:
   ```bash
   go run cmd/main.go
   ```
   The server will start on port `3000` (default).

## Testing
You can perform manual testing using the JetBrains HTTP Client.
- Open `test/manual-test.http`.
- Configure the `@api_key` and `@user_id` variables at the top of the file.
- Run requests directly from your IDE.

## API Documentation
The API specification is available in `docs/api.json` (OpenAPI format).

### Endpoints:
#### Users
- `POST /api/users/register` - Register a new user
- `GET /api/users` - Get all users

#### Projects (Requires `X-API-Key` & `X-User-ID`)
- `GET /api/projects` - List all projects for the user
- `POST /api/projects` - Create a new project
- `GET /api/projects/{projectID}` - Get project details
- `PUT /api/projects/{projectID}` - Update a project
- `DELETE /api/projects/{projectID}` - Delete a project

#### Tasks (Requires `X-API-Key` & `X-User-ID`)
- `GET /api/tasks` - List all tasks for the user
- `POST /api/tasks` - Create a new task
- `GET /api/tasks/{taskID}` - Get task details
- `PUT /api/tasks/{taskID}` - Update a task
- `DELETE /api/tasks/{taskID}` - Delete a task
