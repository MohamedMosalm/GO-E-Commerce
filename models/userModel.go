package models

import "time"

type User struct {
	UserID         int      `gorm:"column:user_id;primaryKey;autoIncrement" json:"user_id"`
	FirstName      string   `gorm:"column:first_name;not null" json:"first_name"`
	LastName       string   `gorm:"column:last_name;not null" json:"last_name"`
	Email          string   `gorm:"column:email;size:100;not null;unique" json:"email"`
	Password       string   `gorm:"-" json:"password"`
	HashedPassword string   `gorm:"column:hashed_password;size:255;not null" json:"-"`
	Reviews        []Review `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"reviews"`
	Orders         []Order  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"orders"`
}

type PasswordResetToken struct {
	ID        uint      `gorm:"column:token_id;primaryKey"`
	Email     string    `gorm:"column:email;"`
	Token     string    `gorm:"column:token;uniqueIndex"`
	ExpiresAt time.Time `gorm:"column:expires_at"`
}
