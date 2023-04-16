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

	r.GET("/:gobitlyId", controllers.VisitGobitly)
	r.POST("/create", controllers.CreateGobitly)
	r.GET("/get-all", controllers.GetAllGobitlies)
	r.DELETE("/delete", controllers.DeleteGobitly)
	r.PUT("/clicked", controllers.UpdateGobitlyClick)
	r.GET("/get", controllers.GetGobitly)
	r.GET("/ping", controllers.Ping)
	r.Run()
}
