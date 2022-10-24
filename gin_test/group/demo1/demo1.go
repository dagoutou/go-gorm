package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func test() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("狗子来了！")
		context.Next()
		fmt.Println("狗子跑了！")
	}
}
func main() {
	g := gin.Default()
	g.Use(test())
	gr := g.Group("/v1")
	{
		gr.GET("/test", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"gouzi": "11111",
			})
		})
	}
	g.Run(":8080")

}
