package controllers

import (
	"gobitly/db"
	"gobitly/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGobitlyById(c *gin.Context) {
	gobitlyId := c.Param("gobitlyId")

	result, err := db.GetGobitlyById(gobitlyId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.GobitlyResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, models.GobitlyResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
}
