package models

type Review struct {
	ReviewID  int     `gorm:"column:review_id;primaryKey;autoIncrement" json:"review_id"`
	Rating    int     `gorm:"column:rating;not null" json:"rating"`
	Comment   string  `gorm:"column:comment;type:text" json:"comment"`
	ProductID int     `gorm:"column:product_id" json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product"`
	UserID    int     `gorm:"column:user_id" json:"user_id"`
	User      User    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}
