package cmd

import (
	"database/sql"
	"fmt"
	"github.com/lhonda/clean-architecture-go-version/config"
	"github.com/lhonda/clean-architecture-go-version/entity"
	"github.com/lhonda/clean-architecture-go-version/infrastructure/repository"
	"github.com/lhonda/clean-architecture-go-version/usecase/order"
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
		dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE)
		db, err := sql.Open("mysql", dataSourceName)

		if err != nil {
			log.Fatal(err.Error())
		}

		defer db.Close()

		repo := repository.NewOrderMySQL(db)

		service := order.NewService(repo)

		owner := args[0]
		pizzaNames := args[1]
		ps := strings.Split(pizzaNames, "=")
		var pizzas []entity.Pizza
		for _, p := range ps {
			pizzas = append(pizzas, )
		}
		newOrder, err := service.CreateOrder(owner, pizzas)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("Created new order:", newOrder.ID)
	},
}

func init() {
	createOrder.PersistentFlags().StringVar(&owner, "owner", "Guido", "")
	createOrder.PersistentFlags().StringVar(&pizzas, "pizzas", "queijo", "pizzas")
	RootCmd.AddCommand(createOrder)
}
