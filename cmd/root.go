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

	DbUser         string
	DbPassword     string
	DbHost         string
	dataSourceName string
)

func init() {
	dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", DbUser, DbPassword, DbHost, "clean_arch")
}

// Execute function
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
