package main

import (
	"fmt"
	//不一致的情况
	method "p1/internal/pkg"
	// 自定义模块名 + 路径(注意包名和目录保持一致)
	"p1/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	//打印斐波那契数列
	f := utils.Fibonacii()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
	//打印hello world
	method.PrintHelloWord()

	//运行 http 服务
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
