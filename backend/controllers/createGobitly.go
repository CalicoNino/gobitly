package controllers

import (
	"gobitly/db"
	"gobitly/models"
	"gobitly/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

func CreateGobitly(c *gin.Context) {
	var gobitly models.Gobitly

	//validate the request body
	if err := c.BindJSON(&gobitly); err != nil {
		c.JSON(http.StatusBadRequest, responses.GobitlyResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&gobitly); validationErr != nil {
		c.JSON(http.StatusBadRequest, responses.GobitlyResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
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
		c.JSON(http.StatusInternalServerError, responses.GobitlyResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, responses.GobitlyResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
}
