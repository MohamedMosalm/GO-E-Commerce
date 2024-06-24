package models

type Category struct {
	CategoryID  int       `gorm:"column:category_id;primaryKey;autoIncrement" json:"category_id"`
	Name        string    `gorm:"column:name;size:50;not null" json:"name"`
	Description string    `gorm:"column:description;type:text" json:"description"`
	Products    []Product `gorm:"many2many:product_category;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"products"`
}
