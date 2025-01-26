package repositories

import "models"

type {{.Entity}}Repository struct {}

func (r *{{.Entity}}Repository) FindAll() []models.{{.Entity}} {
    return []models.{{.Entity}}{}
}

func (r *{{.Entity}}Repository) Save(entity models.{{.Entity}}) error {
    return nil
}