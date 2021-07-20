package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd -
var RootCmd = &cobra.Command{
	Use: "Pizza order.",
	Short: "cli version to set your order",
	Long: `cli version to set your order and list.`,

}


func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}