package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(databaseCmd)
}

var databaseCmd = &cobra.Command{
	Use:   "database [action]",
	Short: "interact with firestore",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("provide an action")
	},
}
