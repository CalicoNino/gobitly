package controllers

import (
	"gobitly/db"
	"gobitly/models"
	"gobitly/utils"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

func CreateGobitly(c *gin.Context) {
	var input models.GobitlyInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.GobitlyResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
		return
	}

	if validationErr := validate.Struct(&input); validationErr != nil {
		c.JSON(http.StatusBadRequest, models.GobitlyResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": validationErr.Error()}})
		return
	}

	_, err := url.ParseRequestURI(input.Url)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.GobitlyResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": "Invalid URL"}})
		return
	}

	now := time.Now()
	expiry := now.AddDate(0, 1, 0)
	mongodbId := primitive.NewObjectID()

	newGobitly := models.Gobitly{
		ID:        mongodbId,
		Redirect:  input.Url,
		Gobitly:   utils.RandomURL(mongodbId.String(), input.Url, now.GoString()),
		CreatedAt: time.Unix(now.Unix(), 0).String(),
		ExpiredAt: time.Unix(expiry.Unix(), 0).String(),
		Clicked:   0,
	}

	result, err := db.InsertGobitly(newGobitly)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.GobitlyResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, models.GobitlyResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
}
