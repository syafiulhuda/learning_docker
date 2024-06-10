package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	wiseWord := "The journey of a thousand miles begins with a single step."

	router := gin.Default()

	handler := func(c *gin.Context) {
		c.String(http.StatusOK, wiseWord) 
		fmt.Println("Someone accessed the server!")
	}

	router.GET("/", handler)

	fmt.Println("Server listening on port 80...")
	err := router.Run(":80") 
	if err != nil {
		panic(err)
	}
}
