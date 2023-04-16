package controllers

import (
	"gobitly/db"
	"gobitly/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllGobitlies(c *gin.Context) {

	result, err := db.GetAllGobitlies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.GobitlyResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, responses.GobitlyResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
}
