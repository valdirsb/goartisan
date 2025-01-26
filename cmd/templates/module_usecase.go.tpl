package usecase

import (

	domain "{{.projectName}}/modules/{{.Entity}}/domain"
	"{{.projectName}}/modules/{{.Entity}}/repository"
	modules "{{.projectName}}/modules/all_modules"
)

type {{.Entity}}UseCase struct {
	repo repository.{{.Entity}}Repository
	*modules.AllModules
}

func New{{.Entity}}UseCase({{.Entity}}Repo repository.{{.Entity}}Repository, modulos *modules.AllModules) *{{.Entity}}UseCase {
	return &{{.Entity}}UseCase{
		repo:       {{.Entity}}Repo,
		AllModules: modulos,
	}
}

func (a *{{.Entity}}UseCase) Find(ID int32) (*domain.{{.Entity}}, error) {

	object, err := a.repo.GetByID(ID)

	if err != nil {
		return nil, err
	}

	return object, nil

}
