package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MohamedMosalm/GO-E-Commerce/cmd/api"
	"github.com/MohamedMosalm/GO-E-Commerce/db"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error fetching database credentials")
		panic(err.Error())
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"), os.Getenv("SSLMode"),
	)

	db.InitDB(dsn)
	// db.Migrate(db.DB)
	// db.DropUslessForeignKeys()

	server := api.NewAPIServer(":8080", db.DB)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
