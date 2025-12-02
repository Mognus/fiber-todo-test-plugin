# Todo Module

Simple CRUD todo list functionality.

## Features

- ✅ Create, Read, Update, Delete todos
- ✅ Soft delete support
- ✅ Structured error handling (via backend-core)
- ✅ Validation
- ✅ RESTful API

## Installation

```bash
go get github.com/yourcompany/todo-module
```

## Usage

```go
import (
    "github.com/yourcompany/todo-module/pkg/todo"
    "yourproject/pkg/module"
)

func main() {
    db := connectDatabase()
    
    // Create module registry
    registry := module.NewRegistry()
    
    // Register todo module
    registry.Register(todo.New(db))
    
    // Run migrations
    registry.MigrateAll(db)
    
    // Register routes
    api := app.Group("/api")
    registry.RegisterAll(api)
}
```

## API Endpoints

- `GET /api/todos` - List all todos
- `GET /api/todos/:id` - Get single todo
- `POST /api/todos` - Create todo
- `PUT /api/todos/:id` - Update todo
- `DELETE /api/todos/:id` - Delete todo

## Example Requests

### Create Todo
```bash
curl -X POST http://localhost:8080/api/todos \
  -H "Content-Type: application/json" \
  -d '{"title": "Buy groceries"}'
```

### List Todos
```bash
curl http://localhost:8080/api/todos
```

### Update Todo
```bash
curl -X PUT http://localhost:8080/api/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Buy groceries", "completed": true}'
```

### Delete Todo
```bash
curl -X DELETE http://localhost:8080/api/todos/1
```

## Model

```go
type Todo struct {
    ID        uint      `json:"id"`
    Title     string    `json:"title"`
    Completed bool      `json:"completed"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

## Dependencies

- `github.com/yourcompany/backend-core` - Error handling utilities
- `github.com/gofiber/fiber/v2` - HTTP framework
- `gorm.io/gorm` - ORM
