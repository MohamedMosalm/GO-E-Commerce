package models

import "time"

type Order struct {
	OrderID    int       `gorm:"column:order_id;primaryKey;autoIncrement" json:"order_id"`
	TotalPrice float64   `gorm:"column:total_price;type:decimal(10,2);not null" json:"total_price"`
	OrderDate  time.Time `gorm:"column:order_date;default:CURRENT_TIMESTAMP" json:"order_date"`
	Status     string    `gorm:"column:status;size:50;not null" json:"status"`
	UserID     int       `gorm:"column:user_id" json:"user_id"`
	User       User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	Products   []Product `gorm:"many2many:order_product;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"products"`
}
