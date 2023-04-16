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
	r.DELETE("/:gobitlyId", controllers.DeleteGobitly)
	r.PUT("/clicked/:gobitlyId", controllers.UpdateGobitlyClick)
	r.GET("/id/:gobitlyId", controllers.GetGobitlyById)
	r.GET("/ping", controllers.Ping)
	r.Run()
}
