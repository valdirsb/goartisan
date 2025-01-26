package {{.Entity}}

import (
	domain "{{.projectName}}/modules/{{.Entity}}/domain"
)

type {{.Entity}}Service interface {
	Find(ID int32) (*domain.{{.Entity}}, error)
}
