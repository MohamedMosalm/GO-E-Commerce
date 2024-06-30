package product

import (
	"github.com/MohamedMosalm/GO-E-Commerce/models"
)

type ProductRepository interface {
	CreateProduct(product *models.Product, CategoriesIds []int) error
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	UpdateProduct(id uint, updates *models.Product) (*models.Product, error)
	DeleteProduct(id uint) error
}
