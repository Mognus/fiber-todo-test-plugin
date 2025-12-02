module github.com/yourcompany/todo-module

go 1.23

require (
	github.com/gofiber/fiber/v2 v2.52.10
	github.com/yourcompany/backend-core v0.0.0
	gorm.io/gorm v1.31.1
)

// For local development
replace github.com/yourcompany/backend-core => ../backend-core
