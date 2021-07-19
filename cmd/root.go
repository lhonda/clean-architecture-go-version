package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)


//var r ingredient.Repository := ingredient.
//var s ingredient.Service

var RootCmd = &cobra.Command{
	Use: "Pizza order.",
	Short: "cli version to set your order",
	Long: "cli version to set your order and list.",
	//Run: func(cmd *cobra.Command, args []string) {
	//	if err := ingredient.(); err != nil {
	//		return err
	//	}
	//	return nil
	// },
}


func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}