package user

import (
	"github.com/MohamedMosalm/GO-E-Commerce/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(id uint, updates *models.User) (*models.User, error)
	DeleteUser(id uint) error
	GetUserByEmail(email string) (*models.User, error)
	SaveIntoDB(user *models.User) error
}

type PasswordResetTokenRepository interface {
	CreatePasswordResetToken(resetToken *models.PasswordResetToken) error
	FindResetToken(token string) (*models.PasswordResetToken, error)
}
