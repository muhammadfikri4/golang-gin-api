package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
	Data    struct {
		Name string `json:"name"`
	} `json:"data"`
}

func main() {
	app := gin.Default()

	// app.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello, World!",
	// 	})
	// })

	app.GET("/", func(ctx *gin.Context) {

		name := ctx.Query("name")

		ctx.JSON(200, gin.H{
			"name": name,
		})
	})

	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	type Response struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	app.POST("/", func(c *gin.Context) {

		var user User

		err := c.ShouldBindJSON(&user)

		if err != nil {
			log.Fatal(err)
		}

		if user.Email == "" || user.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Email and Name are required",
			})
			return
		}

		res := Response{
			Message: "Hello, World!",
			Data:    user,
		}

		c.JSON(200, res)
	})

	app.Run(":3000")
}
