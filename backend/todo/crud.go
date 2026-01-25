package todo

import (
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

// CRUD Operations using default implementations

func (m *Module) List(filters map[string]string, page, limit int) (crud.ListResponse, error) {
	return crud.DefaultList(m.db, &Todo{}, filters, page, limit)
}

func (m *Module) Get(id string) (any, error) {
	return crud.DefaultGet(m.db, &Todo{}, id)
}

func (m *Module) Create(data map[string]any) (any, error) {
	return crud.DefaultCreate(m.db, &Todo{}, data)
}

func (m *Module) Update(id string, data map[string]any) (any, error) {
	return crud.DefaultUpdate(m.db, &Todo{}, id, data)
}

func (m *Module) Delete(id string) error {
	return crud.DefaultDelete(m.db, &Todo{}, id)
}
