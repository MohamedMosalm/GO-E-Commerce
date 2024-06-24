package db

import (
	"log"

	"github.com/MohamedMosalm/GO-E-Commerce/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	log.Println("Connected to the database successfully!")
	DB = db
	return DB
}

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&models.User{}, &models.Product{}, &models.Review{}, &models.Category{}, &models.Order{}); err != nil {
		panic(err)
	}
}

func DropUslessForeignKeys() {
	DB.Exec(`
	ALTER TABLE products DROP CONSTRAINT fk_reviews_product;
	ALTER TABLE users DROP CONSTRAINT fk_orders_user;
	ALTER TABLE users DROP CONSTRAINT fk_reviews_user;
	`)
}
