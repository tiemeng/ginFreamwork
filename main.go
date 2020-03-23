package main

import (
	. "ginFreamwork/todoList"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	server:=gin.Default()
	//server.GET("/", func(context *gin.Context) {
	//	context.JSON(http.StatusOK,gin.H{
	//		"code" : http.StatusOK,
	//		"message" : "请求成功",
	//		"data" : "hello",
	//	})
	//})
	server.GET("/index",Index)
	todo :=server.Group("/todo/")
	{
		todo.GET("list",TodoList)
		todo.POST("add",Add)
		todo.GET("detail",GetDetail)
		todo.GET("delete",Del)
		todo.GET("finish",Finish)
	}
	server.Run()
}
