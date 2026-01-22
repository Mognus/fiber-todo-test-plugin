package todo

import (
	"fmt"
	"strconv"

	"template/modules/core/pkg/crud"
)

// GetModelName implements crud.CRUDProvider
func (m *Module) GetModelName() string {
	return "todos"
}

// GetSchema implements crud.CRUDProvider
func (m *Module) GetSchema() crud.Schema {
	return crud.Schema{
		Name:        "todos",
		DisplayName: "Todos",
		Fields: []crud.Field{
			{Name: "id", Type: "number", Label: "ID", Readonly: true, Editable: true},
			{Name: "title", Type: "string", Label: "Title", Required: true, Editable: true},
			{Name: "completed", Type: "boolean", Label: "Completed", Editable: true},
			{Name: "created_at", Type: "date", Label: "Created", Readonly: true, Editable: true},
			{Name: "updated_at", Type: "date", Label: "Updated", Readonly: true, Editable: true},
		},
		Filterable: []string{"completed"},
		Searchable: []string{"title"},
	}
}

// List implements crud.CRUDProvider
func (m *Module) List(filters map[string]string, page, limit int) (crud.ListResponse, error) {
	var todos []Todo
	query := m.db.Model(&Todo{})

	// Apply filters
	if completed, ok := filters["completed"]; ok {
		isCompleted := completed == "true"
		query = query.Where("completed = ?", isCompleted)
	}

	// Apply search if provided
	if search, ok := filters["search"]; ok && search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where("title ILIKE ?", searchPattern)
	}

	// Get total count
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return crud.ListResponse{}, err
	}

	// Apply pagination
	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&todos).Error; err != nil {
		return crud.ListResponse{}, err
	}

	// Convert to []any
	items := make([]any, len(todos))
	for i, todo := range todos {
		items[i] = todo
	}

	return crud.ListResponse{
		Items: items,
		Total: total,
		Page:  page,
		Limit: limit,
	}, nil
}

// Get implements crud.CRUDProvider
func (m *Module) Get(id string) (any, error) {
	var todo Todo
	todoID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid todo ID: %v", err)
	}

	if err := m.db.First(&todo, todoID).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

// Create implements crud.CRUDProvider
func (m *Module) Create(data map[string]any) (any, error) {
	todo := Todo{
		Completed: false, // Default to not completed
	}

	// Map data to todo fields
	if title, ok := data["title"].(string); ok {
		todo.Title = title
	}
	if completed, ok := data["completed"].(bool); ok {
		todo.Completed = completed
	}

	if err := m.db.Create(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

// Update implements crud.CRUDProvider
func (m *Module) Update(id string, data map[string]any) (any, error) {
	todoID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid todo ID: %v", err)
	}

	var todo Todo
	if err := m.db.First(&todo, todoID).Error; err != nil {
		return nil, err
	}

	// Update fields that are provided
	updates := make(map[string]interface{})

	if title, ok := data["title"].(string); ok {
		updates["title"] = title
	}
	if completed, ok := data["completed"].(bool); ok {
		updates["completed"] = completed
	}

	if err := m.db.Model(&todo).Updates(updates).Error; err != nil {
		return nil, err
	}

	// Reload todo to get updated data
	if err := m.db.First(&todo, todoID).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

// Delete implements crud.CRUDProvider
func (m *Module) Delete(id string) error {
	todoID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return fmt.Errorf("invalid todo ID: %v", err)
	}

	// Soft delete
	if err := m.db.Delete(&Todo{}, todoID).Error; err != nil {
		return err
	}

	return nil
}
