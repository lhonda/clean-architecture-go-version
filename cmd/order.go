package main

import (
    "github.com/spf13/cobra"
)

var createOrder = &cobra.Command{
    Use:   "create-order owner=<Guido> pizzas=peperoni,cheese",
    Short: "Create a new Pizza order",
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
    },
}

func init() {
    RootCmd.AddCommand(createOrder)
}
