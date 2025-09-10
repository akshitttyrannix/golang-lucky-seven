package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!!")

	server := gin.Default()

	server.Run(":8080")
}
