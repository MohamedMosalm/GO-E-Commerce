package category

import (
	"github.com/MohamedMosalm/GO-E-Commerce/models"
)

type CategoryRepository interface {
	CreateCategory(category *models.Category, ProductsIds []int) error
	GetAllCategories() ([]models.Category, error)
	GetCategoryByID(id uint) (*models.Category, error)
	UpdateCategory(id uint, updates *models.Category) (*models.Category, error)
	DeleteCategory(id uint) error
}
