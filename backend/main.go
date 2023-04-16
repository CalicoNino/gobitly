package main

import (
	"gobitly/controllers"
	"gobitly/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// run db
	db.Connect()
	r.POST("/create", controllers.CreateGobitly)
	r.GET("/getAll", controllers.GetAllGobitlies)
	r.GET("/ping", controllers.Ping)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
