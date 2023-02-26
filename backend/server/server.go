package server

import (
	"fmt"
	"gobitly/datastore"
	"gobitly/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func redirect(c *fiber.Ctx) error {
	gobitlyUrl := c.Params("redirect")
	gobitly, err := datastore.FindByGobitly(gobitlyUrl)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not find gobitly in DB " + err.Error(),
		})
	}
	// grab any stats you want...
	gobitly.Clicked += 1
	err = datastore.UpdateGobitly(gobitly)
	if err != nil {
		fmt.Printf("error updating: %v\n", err)
	}

	return c.Redirect(gobitly.Redirect, fiber.StatusTemporaryRedirect)
}

func getAllGobitlies(c *fiber.Ctx) error {
	gobitlies, err := datastore.GetAllGobitlies()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all gobitlies links " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(gobitlies)
}

func getGobitly(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not parse id " + err.Error(),
		})
	}

	gobitly, err := datastore.GetGobitly(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not retreive gobitly from db " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(gobitly)
}

func createGobitly(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var gobitly datastore.Gobitly
	err := c.BodyParser(&gobitly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing JSON " + err.Error(),
		})
	}

	if gobitly.Random {
		gobitly.Gobitly = utils.RandomURL(8)
	}

	err = datastore.CreateGobitly(gobitly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not create gobitly in db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(gobitly)

}

func updateGobitly(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var gobitly datastore.Gobitly

	err := c.BodyParser(&gobitly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not parse json " + err.Error(),
		})
	}

	err = datastore.UpdateGobitly(gobitly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not update gobitly link in DB " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(gobitly)
}

func deleteGobitly(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not parse id from url " + err.Error(),
		})
	}

	err = datastore.DeleteGobitly(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not delete from db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "gobitly deleted.",
	})
}

func SetupAndListen() {

	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/r/:redirect", redirect)

	router.Get("/gobitly", getAllGobitlies)
	router.Get("/gobitly/:id", getGobitly)
	router.Post("/gobitly", createGobitly)
	router.Patch("/gobitly", updateGobitly)
	router.Delete("/gobitly/:id", deleteGobitly)

	router.Listen(":3000")

}
