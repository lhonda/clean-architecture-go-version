package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd -
var (
	RootCmd = &cobra.Command{
		Use:   "Pizza order.",
		Short: "cli version to set your order",
		Long:  `cli version to set your order and list.`,
	}

	DB_USER        string
	DB_PASSWORD    string
	DB_HOST        string
	dataSourceName string
)

func init() {
	dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", DB_USER, DB_PASSWORD, DB_HOST, "clean_arch")
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
