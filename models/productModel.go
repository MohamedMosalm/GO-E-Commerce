package models

type Product struct {
	ProductID   int        `gorm:"column:product_id;primaryKey;autoIncrement" json:"product_id"`
	Name        string     `gorm:"column:name;size:100;not null" json:"name"`
	Description string     `gorm:"column:description;type:text" json:"description"`
	Price       float64    `gorm:"column:price;type:decimal(10,2);not null" json:"price"`
	Stock       int        `gorm:"column:stock;not null" json:"stock"`
	Categories  []Category `gorm:"many2many:product_category;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"categories"`
	Orders      []Order    `gorm:"many2many:order_product;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"orders"`
	Reviews     []Review   `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"reviews"`
}
