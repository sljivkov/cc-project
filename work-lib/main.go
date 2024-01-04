package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	err := server.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
