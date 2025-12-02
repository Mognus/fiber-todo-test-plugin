package todo

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	apperrors "github.com/yourcompany/backend-core/pkg/errors"
	"gorm.io/gorm"
)

// Module implements the pluggable module interface
type Module struct {
	db *gorm.DB
}

// New creates a new Todo module instance
func New(db *gorm.DB) *Module {
	return &Module{db: db}
}

// Name returns the module name
func (m *Module) Name() string {
	return "todos"
}

// RegisterRoutes registers all todo routes
func (m *Module) RegisterRoutes(router fiber.Router) {
	todos := router.Group("/todos")
	todos.Get("/", m.GetTodos)
	todos.Get("/:id", m.GetTodo)
	todos.Post("/", m.CreateTodo)
	todos.Put("/:id", m.UpdateTodo)
	todos.Delete("/:id", m.DeleteTodo)
}

// Migrate runs database migrations
func (m *Module) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Todo{})
}

// GetTodos - List all todos
// GET /api/todos
func (m *Module) GetTodos(c *fiber.Ctx) error {
	var todos []Todo

	if err := m.db.Find(&todos).Error; err != nil {
		appErr := apperrors.Internal(err)
		return c.Status(appErr.Code).JSON(appErr)
	}

	return c.JSON(todos)
}

// GetTodo - Get single todo
// GET /api/todos/:id
func (m *Module) GetTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo Todo

	if err := m.db.First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			appErr := apperrors.NotFound("Todo")
			return c.Status(appErr.Code).JSON(appErr)
		}
		appErr := apperrors.Internal(err)
		return c.Status(appErr.Code).JSON(appErr)
	}

	return c.JSON(todo)
}

// CreateTodo - Create new todo
// POST /api/todos
func (m *Module) CreateTodo(c *fiber.Ctx) error {
	todo := new(Todo)

	if err := c.BodyParser(todo); err != nil {
		appErr := apperrors.BadRequest("Invalid request body")
		return c.Status(appErr.Code).JSON(appErr)
	}

	// Validation
	if todo.Title == "" {
		appErr := apperrors.ValidationError(map[string]string{
			"title": "Title is required",
		})
		return c.Status(appErr.Code).JSON(appErr)
	}

	if err := m.db.Create(&todo).Error; err != nil {
		appErr := apperrors.Internal(err)
		return c.Status(appErr.Code).JSON(appErr)
	}

	return c.Status(201).JSON(todo)
}

// UpdateTodo - Update existing todo
// PUT /api/todos/:id
func (m *Module) UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo Todo

	// Find existing todo
	if err := m.db.First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			appErr := apperrors.NotFound("Todo")
			return c.Status(appErr.Code).JSON(appErr)
		}
		appErr := apperrors.Internal(err)
		return c.Status(appErr.Code).JSON(appErr)
	}

	// Parse update data
	if err := c.BodyParser(&todo); err != nil {
		appErr := apperrors.BadRequest("Invalid request body")
		return c.Status(appErr.Code).JSON(appErr)
	}

	// Validation
	if todo.Title == "" {
		appErr := apperrors.ValidationError(map[string]string{
			"title": "Title is required",
		})
		return c.Status(appErr.Code).JSON(appErr)
	}

	// Save updates
	if err := m.db.Save(&todo).Error; err != nil {
		appErr := apperrors.Internal(err)
		return c.Status(appErr.Code).JSON(appErr)
	}

	return c.JSON(todo)
}

// DeleteTodo - Delete todo
// DELETE /api/todos/:id
func (m *Module) DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo Todo

	result := m.db.Delete(&todo, id)
	if result.Error != nil {
		appErr := apperrors.Internal(result.Error)
		return c.Status(appErr.Code).JSON(appErr)
	}

	if result.RowsAffected == 0 {
		appErr := apperrors.NotFound("Todo")
		return c.Status(appErr.Code).JSON(appErr)
	}

	return c.Status(204).Send(nil)
}
