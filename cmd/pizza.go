package cmd

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // load mysql driver
	"github.com/lhonda/clean-architecture-go-version/config"
	"github.com/lhonda/clean-architecture-go-version/infrastructure/repository"
	"github.com/lhonda/clean-architecture-go-version/usecase/pizza"
	"github.com/spf13/cobra"

	"log"
)

var listPizzas = &cobra.Command{
	Use:   "list-pizzas",
	Short: "list pizzas",
	Run: func(cmd *cobra.Command, args []string) {
		dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", config.DbUser, config.DbPassword, config.DbHost, config.DbDatabase)
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

		for i, p := range pizzas {
			fmt.Println(i, ")", p.Name, ",created at:", p.CreatedAt)
		}
	},
}

func init() {
	RootCmd.AddCommand(listPizzas)
}
