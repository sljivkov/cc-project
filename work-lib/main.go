package main

import (
	"log"
	"os"
	"worker/application"
	"worker/handler"
	"worker/repository"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	// Open a connection to the database using the connection URL
	db, err := gorm.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	migrateDb(db)
	server := gin.Default()
	centralAppUrl := os.Getenv("CENTRAL_URL")
	userHandler := handler.UserHandler{ServerUrl: centralAppUrl}
	borrowHandler := initBorrow(db, centralAppUrl)
	server.POST("user/register", userHandler.Register)
	server.POST("borrow", borrowHandler.Borrow)
	server.PATCH("return", borrowHandler.Return)
	err = server.Run(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Exit")
	}

}

func migrateDb(db *gorm.DB) {
	db.AutoMigrate(&repository.BorrowModel{})
}
func initBorrow(db *gorm.DB, centralAppUrl string) *handler.BorrowHandler {
	repo := repository.NewBorrowRepository(db)
	service := application.NewBorrowService(centralAppUrl, repo)
	return handler.NewBorrowHandler(service)
}
