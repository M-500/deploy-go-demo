package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()
	engine.GET("/hello", SayHello)
	engine.Run(":18888")
}

func SayHello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello World",
	})
}
