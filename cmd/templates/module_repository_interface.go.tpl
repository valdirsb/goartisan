package repository

import (
	domain "{{.projectName}}/modules/{{.Entity}}/domain"
)

type {{.Entity}}Repository interface {
	GetByID(ID int32) (*domain.{{.Entity}}, error)
}
