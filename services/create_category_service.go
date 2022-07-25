package services

import (
	"errors"
	"rentx/models"
	"rentx/repositories"
)

type Request struct {
	Name        string
	Description string
}

type CreateCategoryService interface {
	Execute(request Request) (*models.Category, error)
}

type CreateCategoryServiceImpl struct {
	categoryRepository repositories.CategoriesRepository
}

func NewCreateCategoryService(categoryRepository repositories.CategoriesRepository) CreateCategoryService {
	return &CreateCategoryServiceImpl{categoryRepository}
}

func (c *CreateCategoryServiceImpl) Execute(request Request) (*models.Category, error) {
	category := repositories.CreateCategoryDTO{
		Name:        request.Name,
		Description: request.Description,
	}

	checkAlreadyExists, err := c.categoryRepository.GetByName(category.Name)

	if err != nil {
		return nil, err
	}

	if checkAlreadyExists != nil {
		return nil, errors.New("category already exists")
	}

	newCategory, err := c.categoryRepository.Create(&category)

	if err != nil {
		return nil, err
	}

	return newCategory, nil
}
