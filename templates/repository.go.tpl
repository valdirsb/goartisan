package repositories

import "models"

type {{.Name}}Repository struct {}

func (r *{{.Name}}Repository) FindAll() []models.{{.Name}} {
    return []models.{{.Name}}{}
}

func (r *{{.Name}}Repository) Save(entity models.{{.Name}}) error {
    return nil
}