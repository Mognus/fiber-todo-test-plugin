package todo

import (
	"github.com/gofiber/fiber/v2"
	"template/modules/core/pkg/crud"
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
	// CRUD routes for todos
	todos := router.Group("/todos")
	todos.Get("/", m.ListHandler())
	todos.Get("/schema", m.SchemaHandler())
	todos.Get("/:id", m.GetHandler())
	todos.Post("/", m.CreateHandler())
	todos.Put("/:id", m.UpdateHandler())
	todos.Delete("/:id", m.DeleteHandler())
}

// Migrate runs database migrations
func (m *Module) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Todo{})
}

// Handler method implementations using default helpers

func (m *Module) ListHandler() fiber.Handler {
	return crud.DefaultListHandler(m)
}

func (m *Module) SchemaHandler() fiber.Handler {
	return crud.DefaultSchemaHandler(m)
}

func (m *Module) GetHandler() fiber.Handler {
	return crud.DefaultGetHandler(m)
}

func (m *Module) CreateHandler() fiber.Handler {
	return crud.DefaultCreateHandler(m)
}

func (m *Module) UpdateHandler() fiber.Handler {
	return crud.DefaultUpdateHandler(m)
}

func (m *Module) DeleteHandler() fiber.Handler {
	return crud.DefaultDeleteHandler(m)
}
