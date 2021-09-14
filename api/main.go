package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/lhonda/clean-architecture-go-version/config"
	"github.com/lhonda/clean-architecture-go-version/entity"
	"github.com/lhonda/clean-architecture-go-version/infrastructure/repository"
	"github.com/lhonda/clean-architecture-go-version/usecase/order"
	"github.com/lhonda/clean-architecture-go-version/usecase/pizza"
	"log"
	"net/http"
)

func listPizzas(c *fiber.Ctx) error {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE)
	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	pizzaRepo := repository.NewPizzaMySQL(db)
	pizzaService := pizza.NewService(pizzaRepo)
	pizzas, err := pizzaService.ListPizzas()

	if err != nil {
		log.Fatal(err.Error())
	}

	return c.JSON(pizzas)
}

func createOrder(c *fiber.Ctx) error {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE)
	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	orderRepo := repository.NewOrderMySQL(db)

	orderService := order.NewService(orderRepo)

	o := new(entity.Order)
	if err := c.BodyParser(o); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	newOrder, err := orderService.CreateOrder(o.Owner, o.Pizzas)
	if err != nil {
		log.Fatal(err.Error())
	}

	o.ID = newOrder.ID

	c.Status(http.StatusCreated)
	c.JSON(o)

	return nil
}

func main() {
	app := fiber.New()
	app.Get("/pizzas", listPizzas)
	app.Post("/orders", createOrder)
	app.Listen(":3000")
}
