package user

import (
	"github.com/MohamedMosalm/GO-E-Commerce/models"
	"gorm.io/gorm"
)

type GromUserRepository struct {
	db *gorm.DB
}

func NewGromUserRepository(db *gorm.DB) *GromUserRepository {
	return &GromUserRepository{db: db}
}

func (r *GromUserRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *GromUserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *GromUserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.Where("user_id = ?", id).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GromUserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *GromUserRepository) UpdateUser(id uint, updates *models.User) (*models.User, error) {
	user, err := r.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	if updates.Email == "" {
		updates.Email = user.Email
	}
	if updates.FirstName == "" {
		updates.FirstName = user.FirstName
	}
	if updates.LastName == "" {
		updates.LastName = user.LastName
	}
	updates.HashedPassword = user.HashedPassword
	err = r.db.Save(user).Error

	return user, err
}

func (r *GromUserRepository) DeleteUser(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
