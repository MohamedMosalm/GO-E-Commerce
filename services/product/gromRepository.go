package product

import (
	"github.com/MohamedMosalm/GO-E-Commerce/models"
	"gorm.io/gorm"
)

type GromProductRepository struct {
	db *gorm.DB
}

func NewGromProductRepository(db *gorm.DB) *GromProductRepository {
	return &GromProductRepository{db: db}
}

func (r *GromProductRepository) CreateProduct(product *models.Product, CategoriesIds []int) error {
	if err := r.db.Create(product).Error; err != nil {
		return err
	}
	for _, id := range CategoriesIds {
		var category models.Category
		if err := r.db.Where("category_id = ?", id).First(&category).Error; err != nil {
			return err
		}
		r.db.Model(product).Association("Categories").Append(&category)
	}
	return nil
}

func (r *GromProductRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product

	if err := r.db.Preload("Categories").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *GromProductRepository) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := r.db.Preload("Categories").Where("product_id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *GromProductRepository) UpdateProduct(id uint, updates *models.Product) (*models.Product, error) {
	product, err := r.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	if updates.Name != "" {
		product.Name = updates.Name
	}
	if updates.Description != "" {
		product.Description = updates.Description
	}
	if updates.Price != 0 {
		product.Price = updates.Price
	}
	if updates.Stock != 0 {
		product.Stock = updates.Stock
	}

	err = r.db.Save(product).Error

	return product, err
}

func (r *GromProductRepository) DeleteProduct(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}
