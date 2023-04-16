package controllers

import (
	"gobitly/db"
	"gobitly/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGobitly(c *gin.Context) {
	var gobitlyId models.GobitlyIdInput

	//validate the request body
	if err := c.BindJSON(&gobitlyId); err != nil {
		c.JSON(http.StatusBadRequest, models.GobitlyResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}

	if validationErr := validate.Struct(&gobitlyId); validationErr != nil {
		c.JSON(http.StatusBadRequest, models.GobitlyResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": validationErr.Error()}})
		return
	}

	result, err := db.GetGobitly(gobitlyId.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.GobitlyResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, models.GobitlyResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
}
