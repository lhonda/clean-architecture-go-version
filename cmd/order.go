package cmd

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lhonda/clean-architecture-go-version/config"
	"github.com/lhonda/clean-architecture-go-version/entity"
	"github.com/lhonda/clean-architecture-go-version/infrastructure/repository"
	"github.com/lhonda/clean-architecture-go-version/usecase/order"
	"github.com/lhonda/clean-architecture-go-version/usecase/pizza"
	"github.com/spf13/cobra"

	"log"
	"strings"
)

var (
	owner  string
	pizzas string
)

var createOrder = &cobra.Command{
	Use:   "create-order owner=<Guido> pizzas=peperoni,cheese",
	Short: "Create a new Pizza order",
	Run: func(cmd *cobra.Command, args []string) {
		dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE)
		db, err := sql.Open("mysql", dataSourceName)

		if err != nil {
			log.Fatal(err.Error())
		}

		defer db.Close()

		orderRepo := repository.NewOrderMySQL(db)

		orderService := order.NewService(orderRepo)

		pizzaRepo := repository.NewPizzaMySQL(db)

		pizzaService := pizza.NewService(pizzaRepo)

		ps := strings.Split(pizzas, "=")
		var pizzas []entity.Pizza

		for _, p := range strings.Split(ps[0],",") {
			pizza, _ := pizzaService.GetPizzaByName(p)
			pizzas = append(pizzas, *pizza)
		}
		newOrder, err := orderService.CreateOrder(owner, pizzas)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("Created new order:", newOrder.ID)
	},
}

func init() {
	createOrder.Flags().StringVar(&owner, "owner", "Guido", "")
	createOrder.Flags().StringVar(&pizzas, "pizzas", "queijo", "pizzas")
	createOrder.MarkFlagRequired("owner")
	createOrder.MarkFlagRequired("pizzas")
	RootCmd.AddCommand(createOrder)
}
