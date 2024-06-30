package category

import (
	"github.com/MohamedMosalm/GO-E-Commerce/models"
	"gorm.io/gorm"
)

type GromCategoryRepository struct {
	db *gorm.DB
}

func NewGromCategoryRepository(db *gorm.DB) *GromCategoryRepository {
	return &GromCategoryRepository{db: db}
}

func (r *GromCategoryRepository) CreateCategory(category *models.Category, ProductsIds []int) error {
	if err := r.db.Create(category).Error; err != nil {
		return err
	}
	for _, id := range ProductsIds {
		var product models.Product
		if err := r.db.Where("product_id = ?", id).First(&product).Error; err != nil {
			return err
		}
		r.db.Model(category).Association("Products").Append(&product)
	}
	return nil
}

func (r *GromCategoryRepository) GetAllCategories() ([]models.Category, error) {
	var categorys []models.Category

	if err := r.db.Preload("Products").Find(&categorys).Error; err != nil {
		return nil, err
	}
	return categorys, nil
}

func (r *GromCategoryRepository) GetCategoryByID(id uint) (*models.Category, error) {
	var category models.Category
	if err := r.db.Preload("Products").Where("category_id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *GromCategoryRepository) UpdateCategory(id uint, updates *models.Category) (*models.Category, error) {
	category, err := r.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	if updates.Name != "" {
		category.Name = updates.Name
	}
	if updates.Description != "" {
		category.Description = updates.Description
	}
	if updates.Description != "" {
		category.Description = updates.Description
	}

	err = r.db.Save(category).Error

	return category, err
}

func (r *GromCategoryRepository) DeleteCategory(id uint) error {
	return r.db.Delete(&models.Category{}, id).Error
}
