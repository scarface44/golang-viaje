package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "You Want A Job, Ernie?")
	})

	r.Run(":3000")
}
