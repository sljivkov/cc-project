package main

import (
	"log"
	"master/application"
	"master/handler"
	"master/repository"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	server := gin.Default()
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
	userHandler := initiateUserHandler(db)
	bookHandler, bookRepo := initiateBookHandler(db)
	borrowHandler := initiateBorrowHandler(db, bookRepo)
	server.POST("/user/register", userHandler.Register)
	server.POST("/book/create", bookHandler.Create)
	server.POST("/borrow", borrowHandler.Borrow)
	server.PATCH("/return", borrowHandler.Return)
	err = server.Run(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Exit")
	}
}

func migrateDb(db *gorm.DB) {
	db.AutoMigrate(&repository.UserModel{})
	db.AutoMigrate(&repository.BookModel{})
	db.AutoMigrate(&repository.BorrowModel{})
}

func initiateUserHandler(db *gorm.DB) *handler.UserHandler {
	userRepo := repository.UserRepository{DB: db}
	userService := application.NewUserService(userRepo)
	return handler.NewUserHandler(userService)
}

func initiateBookHandler(db *gorm.DB) (*handler.BookHandler, *repository.BookRepository) {
	repo := repository.BookRepository{DB: db}
	bookService := application.NewBookService(repo)
	return handler.NewBookHandler(bookService), &repo
}

func initiateBorrowHandler(db *gorm.DB, bookRepo *repository.BookRepository) *handler.BorrowHandler {
	repo := repository.NewBorrowRepository(db)
	service := application.NewBorrowService(repo, bookRepo)
	return handler.NewBorrowHandler(service)
}
