package controllers

import (
	"gobitly/db"
	"gobitly/models"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

func CreateGobitly(c *gin.Context) {
	var gobitly models.Gobitly

	if err := c.BindJSON(&gobitly); err != nil {
		c.JSON(http.StatusBadRequest, models.GobitlyResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}

	if validationErr := validate.Struct(&gobitly); validationErr != nil {
		c.JSON(http.StatusBadRequest, models.GobitlyResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": validationErr.Error()}})
		return
	}

	_, err := url.ParseRequestURI(gobitly.Redirect)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.GobitlyResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": "Invalid URL"}})
		return
	}

	newGobitly := models.Gobitly{
		ID:       primitive.NewObjectID(),
		Redirect: gobitly.Redirect,
		Gobitly:  gobitly.Gobitly,
		Random:   gobitly.Random,
		Clicked:  0,
	}

	result, err := db.InsertGobitly(newGobitly)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.GobitlyResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, models.GobitlyResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
}
