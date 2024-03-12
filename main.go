package main

import (
	"awesomeProject1/cmd/cors"
	"awesomeProject1/cmd/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	user.Hello()
	router := gin.Default()

	//router.Use(CORSMiddleware())
	corseSettings := cors.CorsSettings()
	router.Use(func(c *gin.Context) {
		corseSettings.HandlerFunc(c.Writer, c.Request)
		c.Next()
	})

	router.GET("/user", homePage)
	router.POST("/authorise", authorise)
	router.Run()
}

func homePage(c *gin.Context) {
	c.String(http.StatusOK, "Hello user")
}

func authorise(c *gin.Context) {
	c.String(http.StatusOK, "Getting authorisation")
	fmt.Println("Get_front_info")
	fmt.Println(c)
	fmt.Println(c.PostForm("text"))
	fmt.Println(c.PostForm("title"))

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
