package controllers

import (
	"gobitly/db"
	"gobitly/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VisitGobitly(c *gin.Context) {
	gobitlyId := c.Param("gobitlyId")

	_, err2 := db.UpdateGobitlyClick(gobitlyId)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, models.GobitlyResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"error": err2.Error()}})
		return
	}

	result, err := db.GetGobitly(gobitlyId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.GobitlyResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, models.GobitlyResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
}
