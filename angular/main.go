package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	InitDatabase()
	r := SetupRouter()

	r.GET("/", func(c *gin.Context) {
		books := Book{}.DbList()
		movies := Movie{}.DbList()
		c.HTML(http.StatusOK, "index.html", gin.H{
			"books":  books,
			"movies": movies,
		})
	})

	r.Run("127.0.0.1:8080")
}
