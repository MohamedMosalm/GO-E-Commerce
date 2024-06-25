package user

import (
	"time"

	"github.com/MohamedMosalm/GO-E-Commerce/models"
	"gorm.io/gorm"
)

type GromPasswordResetTokenRepository struct {
	db *gorm.DB
}

func NewGromPasswordResetTokenRepository(db *gorm.DB) *GromPasswordResetTokenRepository {
	return &GromPasswordResetTokenRepository{db: db}
}

func (r *GromPasswordResetTokenRepository) CreatePasswordResetToken(resetToken *models.PasswordResetToken) error {
	return r.db.Create(resetToken).Error
}

func (r *GromPasswordResetTokenRepository) FindResetToken(token string) (*models.PasswordResetToken, error) {
	var resetToken models.PasswordResetToken
	if err := r.db.Where("token = ? AND expires_at > ?", token, time.Now()).First(&resetToken).Error; err != nil {
		return nil, err
	}
	return &resetToken, nil
}
