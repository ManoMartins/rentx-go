package repositories

import (
	"github.com/google/uuid"
	"rentx/models"
	"sort"
	"time"
)

type CreateCategoryDTO struct {
	Name        string
	Description string
}

type CategoriesRepository interface {
	GetAll() (*[]models.Category, error)
	GetByID(id string) (*models.Category, error)
	GetByName(name string) (*models.Category, error)
	Create(category *CreateCategoryDTO) (*models.Category, error)
	Update(category *CreateCategoryDTO) (*models.Category, error)
	Delete(id string) error
}

type CategoriesRepositoryImpl struct {
	categories *[]models.Category
}

func NewCategoriesRepository() CategoriesRepository {
	categories := make([]models.Category, 0)

	return &CategoriesRepositoryImpl{
		categories: &categories,
	}
}

func (c *CategoriesRepositoryImpl) GetAll() (*[]models.Category, error) {
	//TODO implement me
	return c.categories, nil
}

func (c *CategoriesRepositoryImpl) GetByID(id string) (*models.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CategoriesRepositoryImpl) GetByName(name string) (*models.Category, error) {
	idx := sort.Search(len(*c.categories), func(i int) bool {
		return (*c.categories)[i].Name >= name
	})

	if idx == len(*c.categories) || (*c.categories)[idx].Name != name {
		return nil, nil
	}

	return &(*c.categories)[idx], nil
}

func (c *CategoriesRepositoryImpl) Create(category *CreateCategoryDTO) (*models.Category, error) {
	createCategory := models.Category{
		ID:          uuid.New().String(),
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   time.Now(),
	}

	*c.categories = append(*c.categories, createCategory)

	return &createCategory, nil
}

func (c *CategoriesRepositoryImpl) Update(category *CreateCategoryDTO) (*models.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CategoriesRepositoryImpl) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
